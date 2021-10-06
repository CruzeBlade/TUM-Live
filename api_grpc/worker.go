// Package api_grpc handles communication with workers
package api_grpc

import (
	"TUM-Live/api"
	"TUM-Live/dao"
	"TUM-Live/model"
	"TUM-Live/tools"
	"context"
	"errors"
	"fmt"
	go_anel_pwrctrl "github.com/RBG-TUM/go-anel-pwrctrl"
	"github.com/getsentry/sentry-go"
	"github.com/joschahenningsen/TUM-Live-Worker-v2/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"strings"
	"sync"
	"time"
)

var mutex = sync.RWMutex{}

type server struct {
	pb.UnimplementedFromWorkerServer
}

//NotifySilenceResults handles the results of silence detection sent by a worker
func (s server) NotifySilenceResults(ctx context.Context, request *pb.SilenceResults) (*pb.Status, error) {
	if _, err := dao.GetWorkerByID(ctx, request.WorkerID); err != nil {
		return nil, err
	}
	if _, err := dao.GetStreamByID(ctx, fmt.Sprintf("%d", request.GetStreamID())); err != nil {
		return nil, err
	}
	var silences []model.Silence
	for i := range request.Starts {
		silences = append(silences, model.Silence{
			Start:    uint(request.Starts[i]),
			End:      uint(request.Ends[i]),
			StreamID: uint(request.StreamID),
		})
	}
	if len(silences) == 0 {
		return &pb.Status{Ok: true}, nil
	}
	if err := dao.UpdateSilences(silences, fmt.Sprintf("%d", request.StreamID)); err != nil {
		return nil, err
	}
	return &pb.Status{Ok: true}, nil

}

// SendSelfStreamRequest handles the request from a worker when a stream starts publishing via obs, etc.
// returns an error if anything goes wrong OR the stream may not be published.
func (s server) SendSelfStreamRequest(ctx context.Context, request *pb.SelfStreamRequest) (*pb.SelfStreamResponse, error) {
	if _, err := dao.GetWorkerByID(ctx, request.GetWorkerID()); err != nil {
		return nil, err
	}
	stream, err := dao.GetStreamByKey(ctx, request.StreamKey)
	if err != nil {
		return nil, err
	}
	// reject streams that are more than 30 minutes in the future or more than 30 minutes past
	if !(time.Now().After(stream.Start.Add(time.Minute*-30)) && time.Now().Before(stream.End.Add(time.Minute*30))) {
		log.WithFields(log.Fields{"streamId": stream.ID}).Warn("Stream rejected, time out of bounds")
		return nil, errors.New("stream rejected")
	}
	course, err := dao.GetCourseById(ctx, stream.CourseID)
	if err != nil {
		log.WithError(err).Warn("Can't get stream for worker")
		return nil, err
	}
	return &pb.SelfStreamResponse{
		StreamID:    uint32(stream.ID),
		CourseSlug:  course.Slug,
		CourseYear:  uint32(course.Year),
		StreamStart: timestamppb.New(stream.Start),
		CourseTerm:  course.TeachingTerm,
		UploadVoD:   course.VODEnabled,
	}, nil
}

var lightLock = sync.Mutex{}

// NotifyStreamStart handles workers notification about streams being started
// Deprecated: this is now "NotifyStreamStarted"
func (s server) NotifyStreamStart(ctx context.Context, request *pb.StreamStarted) (*pb.Status, error) {
	mutex.Lock()
	defer mutex.Unlock()
	_, err := dao.GetWorkerByID(ctx, request.GetWorkerID())
	if err != nil {
		log.WithField("request", request).Warn("Got stream start with invalid WorkerID")
		return nil, err
	}
	stream, err := dao.GetStreamByID(ctx, fmt.Sprintf("%d", request.StreamID))
	if err != nil {
		log.WithError(err).Warn("Can't get stream by ID to set live")
		return nil, err
	}
	stream.LiveNow = true
	switch request.GetSourceType() {
	case "CAM":
		stream.PlaylistUrlCAM = request.HlsUrl
	case "PRES":
		stream.PlaylistUrlPRES = request.HlsUrl
	default:
		stream.PlaylistUrl = request.HlsUrl
	}
	err = dao.SaveStream(&stream)
	if err != nil {
		log.WithError(err).Error("Can't save stream when setting live")
		return nil, err
	}
	return nil, nil
}

//NotifyStreamFinished handles workers notification about streams being finished
func (s server) NotifyStreamFinished(ctx context.Context, request *pb.StreamFinished) (*pb.Status, error) {
	if _, err := dao.GetWorkerByID(ctx, request.GetWorkerID()); err != nil {
		return nil, errors.New("authentication failed: invalid worker id")
	} else {
		stream, err := dao.GetStreamByID(ctx, fmt.Sprintf("%d", request.StreamID))
		if err != nil {
			log.WithError(err).Error("Can't find stream to set not live")
		} else {
			if dao.HasStreamLectureHall(stream.ID) {
				if lectureHall, err := dao.GetLectureHallByID(stream.LectureHallID); err != nil {
					return nil, err
				} else {
					client := go_anel_pwrctrl.New(lectureHall.PwrCtrlIp, tools.Cfg.PWRCTRLAuth)
					lightLock.Lock()
					err := client.TurnOff(lectureHall.LiveLightIndex)
					if err != nil {
						log.WithError(err).Error("Can't turn off live light.")
					}
					lightLock.Unlock()
				}
			}
		}
		err = dao.SetStreamNotLiveById(uint(request.StreamID))
		if err != nil {
			log.WithError(err).Error("Can't set stream not live")
		}
		api.NotifyViewersLiveState(uint(request.StreamID), false)
	}
	return &pb.Status{Ok: true}, nil
}

//SendHeartBeat receives heartbeat messages sent by workers
func (s server) SendHeartBeat(ctx context.Context, request *pb.HeartBeat) (*pb.Status, error) {
	if worker, err := dao.GetWorkerByID(ctx, request.GetWorkerID()); err != nil {
		return nil, errors.New("authentication failed: invalid worker id")
	} else {
		worker.Workload = int(request.Workload)
		worker.LastSeen = time.Now()
		worker.Status = strings.Join(request.Jobs, ", ")
		err := dao.SaveWorker(worker)
		if err != nil {
			return nil, err
		}
		return &pb.Status{Ok: true}, nil
	}
}

//NotifyTranscodingFinished receives and handles messages from workers about finished transcoding
func (s server) NotifyTranscodingFinished(ctx context.Context, request *pb.TranscodingFinished) (*pb.Status, error) {
	if _, err := dao.GetWorkerByID(ctx, request.WorkerID); err != nil {
		return nil, err
	}
	stream, err := dao.GetStreamByID(ctx, fmt.Sprintf("%d", request.StreamID))
	if err != nil {
		return nil, err
	}
	stream.Files = append(stream.Files, model.File{StreamID: stream.ID, Path: request.FilePath})
	err = dao.SaveStream(&stream)
	if err != nil {
		log.WithError(err).Error("Can't save stream")
		return nil, err
	}
	return &pb.Status{Ok: true}, nil
}

//NotifyUploadFinished receives and handles messages from workers about finished uploads
func (s server) NotifyUploadFinished(ctx context.Context, req *pb.UploadFinished) (*pb.Status, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, err := dao.GetWorkerByID(ctx, req.WorkerID); err != nil {
		return nil, err
	}
	stream, err := dao.GetStreamByID(ctx, fmt.Sprintf("%d", req.StreamID))
	if err != nil {
		return nil, err
	}
	stream.Recording = true
	switch req.SourceType {
	case "CAM":
		stream.PlaylistUrlCAM = req.HLSUrl
	case "PRES":
		stream.PlaylistUrlPRES = req.HLSUrl
	default:
		stream.PlaylistUrl = req.HLSUrl
	}
	if err = dao.SaveStream(&stream); err != nil {
		return nil, err
	}
	return &pb.Status{Ok: true}, nil
}

//NotifyStreamStarted receives stream started events from workers
func (s server) NotifyStreamStarted(ctx context.Context, request *pb.StreamStarted) (*pb.Status, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, err := dao.GetWorkerByID(ctx, request.WorkerID); err != nil {
		return nil, err
	}
	stream, err := dao.GetStreamByID(ctx, fmt.Sprintf("%d", request.GetStreamID()))
	if err != nil {
		log.WithError(err).Println("Can't find stream")
		return nil, err
	}
	if dao.HasStreamLectureHall(stream.ID) {
		if lectureHall, err := dao.GetLectureHallByID(stream.LectureHallID); err != nil {
			return nil, err
		} else {
			lightLock.Lock()
			client := go_anel_pwrctrl.New(lectureHall.PwrCtrlIp, tools.Cfg.PWRCTRLAuth)
			go func() {
				err := client.TurnOn(lectureHall.LiveLightIndex)
				if err != nil {
					log.WithError(err).Error("Can't turn on live light.")
				}
			}()
			defer lightLock.Unlock()
		}
	}
	stream.LiveNow = true
	switch request.GetSourceType() {
	case "CAM":
		stream.PlaylistUrlCAM = request.HlsUrl
	case "PRES":
		stream.PlaylistUrlPRES = request.HlsUrl
	default:
		stream.PlaylistUrl = request.HlsUrl
	}
	err = dao.SaveStream(&stream)
	if err != nil {
		log.WithError(err).Println("Can't set stream live")
		return nil, err
	}
	api.NotifyViewersLiveState(stream.Model.ID, true)
	return &pb.Status{Ok: true}, nil
}

// init initializes a gRPC server on port 50052
func init() {
	log.Info("Serving heartbeat")
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.WithError(err).Error("Failed to init grpc server")
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFromWorkerServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.WithError(err).Errorf("Can't serve grpc")
		}
	}()
}

// NotifyWorkers collects all streams that are due to stream
// (starts in the next 10 minutes from a lecture hall)
// and invokes the corresponding calls at the workers with the least workload via gRPC
func NotifyWorkers() {
	notifyWorkersPremieres()
	streams := dao.GetDueStreamsForWorkers()
	workers := dao.GetAliveWorkers()
	if len(workers) == 0 && len(streams) != 0 {
		log.Error("not enough workers to handle streams")
		return
	}
	for i := range streams {
		courseForStream, err := dao.GetCourseById(context.Background(), streams[i].CourseID)
		if err != nil {
			log.WithError(err).Warn("Can't get course for stream, skipping")
			sentry.CaptureException(err)
			continue
		}
		lectureHallForStream, err := dao.GetLectureHallByID(streams[i].LectureHallID)
		if err != nil {
			log.WithError(err).Error("Can't get lecture hall for stream, skipping")
			sentry.CaptureException(err)
			continue
		}
		sources := []string{lectureHallForStream.CombIP, lectureHallForStream.PresIP, lectureHallForStream.CamIP}
		for sourceNum, source := range sources {
			if source == "" {
				continue
			}
			var sourceType string
			if sourceNum == 1 {
				sourceType = "PRES"
			} else if sourceNum == 2 {
				sourceType = "CAM"
			} else {
				sourceType = "COMB"
			}
			req := pb.StreamRequest{
				SourceType:    sourceType,
				SourceUrl:     source,
				CourseSlug:    courseForStream.Slug,
				Start:         timestamppb.New(streams[i].Start),
				End:           timestamppb.New(streams[i].End),
				PublishStream: courseForStream.LiveEnabled,
				PublishVoD:    courseForStream.VODEnabled,
				StreamID:      uint32(streams[i].ID),
				CourseTerm:    courseForStream.TeachingTerm,
				CourseYear:    uint32(courseForStream.Year),
			}
			workerIndex := getWorkerWithLeastWorkload(workers)
			workers[workerIndex].Workload += 3
			conn, err := grpc.Dial(fmt.Sprintf("%s:50051", workers[workerIndex].Host), grpc.WithInsecure())
			if err != nil {
				log.WithError(err).Error("Unable to dial server")
				workers[workerIndex].Workload -= 1 // decrease workers load only by one (backoff)
				continue
			}
			client := pb.NewToWorkerClient(conn)
			req.WorkerId = workers[workerIndex].WorkerID
			resp, err := client.RequestStream(context.Background(), &req)
			if err != nil || !resp.Ok {
				log.WithError(err).Error("could not assign stream!")
				workers[workerIndex].Workload -= 1 // decrease workers load only by one (backoff)
				continue
			}
			_ = conn.Close()
		}
	}
}

//notifyWorkersPremieres looks for premieres that should be streamed and assigns them to workers.
func notifyWorkersPremieres() {
	streams := dao.GetDuePremieresForWorkers()
	workers := dao.GetAliveWorkers()

	if len(workers) == 0 && len(streams) != 0 {
		log.Error("Not enough alive workers for premiere")
		return
	}
	for i := range streams {
		if len(streams[i].Files) == 0 {
			log.WithField("streamID", streams[i].ID).Warn("Request to self stream without file")
			continue
		}
		workerIndex := getWorkerWithLeastWorkload(workers)
		workers[workerIndex].Workload += 3
		req := pb.PremiereRequest{
			StreamID: uint32(streams[i].ID),
			FilePath: streams[i].Files[0].Path,
		}
		conn, err := grpc.Dial(fmt.Sprintf("%s:50051", workers[workerIndex].Host), grpc.WithInsecure())
		if err != nil {
			log.WithError(err).Error("Unable to dial server")
			_ = conn.Close()
			workers[workerIndex].Workload -= 1
			continue
		}
		client := pb.NewToWorkerClient(conn)
		req.WorkerID = workers[workerIndex].WorkerID
		resp, err := client.RequestPremiere(context.Background(), &req)
		if err != nil || !resp.Ok {
			log.WithError(err).Error("could not assign premiere!")
			workers[workerIndex].Workload -= 1
			_ = conn.Close()
			continue
		}
		_ = conn.Close()
	}
}

//getWorkerWithLeastWorkload Gets the index of the worker from workers with the least workload.
//workers must not be empty!
func getWorkerWithLeastWorkload(workers []model.Worker) int {
	foundWorker := 0
	for i := range workers {
		if workers[i].Workload < workers[foundWorker].Workload {
			foundWorker = i
		}
	}
	return foundWorker
}