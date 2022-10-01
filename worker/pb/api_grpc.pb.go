// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: api.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ToWorkerClient is the client API for ToWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToWorkerClient interface {
	// Requests a stream from a lecture hall
	RequestStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*Status, error)
	RequestPremiere(ctx context.Context, in *PremiereRequest, opts ...grpc.CallOption) (*Status, error)
	RequestStreamEnd(ctx context.Context, in *EndStreamRequest, opts ...grpc.CallOption) (*Status, error)
	RequestWaveform(ctx context.Context, in *WaveformRequest, opts ...grpc.CallOption) (*WaveFormResponse, error)
	RequestCut(ctx context.Context, in *CutRequest, opts ...grpc.CallOption) (*CutResponse, error)
	GenerateThumbnails(ctx context.Context, in *GenerateThumbnailRequest, opts ...grpc.CallOption) (*Status, error)
	GenerateSectionImages(ctx context.Context, in *GenerateSectionImageRequest, opts ...grpc.CallOption) (*GenerateSectionImageResponse, error)
	DeleteSectionImage(ctx context.Context, in *DeleteSectionImageRequest, opts ...grpc.CallOption) (*Status, error)
}

type toWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewToWorkerClient(cc grpc.ClientConnInterface) ToWorkerClient {
	return &toWorkerClient{cc}
}

func (c *toWorkerClient) RequestStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.ToWorker/RequestStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toWorkerClient) RequestPremiere(ctx context.Context, in *PremiereRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.ToWorker/RequestPremiere", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toWorkerClient) RequestStreamEnd(ctx context.Context, in *EndStreamRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.ToWorker/RequestStreamEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toWorkerClient) RequestWaveform(ctx context.Context, in *WaveformRequest, opts ...grpc.CallOption) (*WaveFormResponse, error) {
	out := new(WaveFormResponse)
	err := c.cc.Invoke(ctx, "/api.ToWorker/RequestWaveform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toWorkerClient) RequestCut(ctx context.Context, in *CutRequest, opts ...grpc.CallOption) (*CutResponse, error) {
	out := new(CutResponse)
	err := c.cc.Invoke(ctx, "/api.ToWorker/RequestCut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toWorkerClient) GenerateThumbnails(ctx context.Context, in *GenerateThumbnailRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.ToWorker/GenerateThumbnails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toWorkerClient) GenerateSectionImages(ctx context.Context, in *GenerateSectionImageRequest, opts ...grpc.CallOption) (*GenerateSectionImageResponse, error) {
	out := new(GenerateSectionImageResponse)
	err := c.cc.Invoke(ctx, "/api.ToWorker/GenerateSectionImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toWorkerClient) DeleteSectionImage(ctx context.Context, in *DeleteSectionImageRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.ToWorker/DeleteSectionImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToWorkerServer is the server API for ToWorker service.
// All implementations must embed UnimplementedToWorkerServer
// for forward compatibility
type ToWorkerServer interface {
	// Requests a stream from a lecture hall
	RequestStream(context.Context, *StreamRequest) (*Status, error)
	RequestPremiere(context.Context, *PremiereRequest) (*Status, error)
	RequestStreamEnd(context.Context, *EndStreamRequest) (*Status, error)
	RequestWaveform(context.Context, *WaveformRequest) (*WaveFormResponse, error)
	RequestCut(context.Context, *CutRequest) (*CutResponse, error)
	GenerateThumbnails(context.Context, *GenerateThumbnailRequest) (*Status, error)
	GenerateSectionImages(context.Context, *GenerateSectionImageRequest) (*GenerateSectionImageResponse, error)
	DeleteSectionImage(context.Context, *DeleteSectionImageRequest) (*Status, error)
	mustEmbedUnimplementedToWorkerServer()
}

// UnimplementedToWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedToWorkerServer struct {
}

func (UnimplementedToWorkerServer) RequestStream(context.Context, *StreamRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestStream not implemented")
}
func (UnimplementedToWorkerServer) RequestPremiere(context.Context, *PremiereRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPremiere not implemented")
}
func (UnimplementedToWorkerServer) RequestStreamEnd(context.Context, *EndStreamRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestStreamEnd not implemented")
}
func (UnimplementedToWorkerServer) RequestWaveform(context.Context, *WaveformRequest) (*WaveFormResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestWaveform not implemented")
}
func (UnimplementedToWorkerServer) RequestCut(context.Context, *CutRequest) (*CutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestCut not implemented")
}
func (UnimplementedToWorkerServer) GenerateThumbnails(context.Context, *GenerateThumbnailRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateThumbnails not implemented")
}
func (UnimplementedToWorkerServer) GenerateSectionImages(context.Context, *GenerateSectionImageRequest) (*GenerateSectionImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSectionImages not implemented")
}
func (UnimplementedToWorkerServer) DeleteSectionImage(context.Context, *DeleteSectionImageRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSectionImage not implemented")
}
func (UnimplementedToWorkerServer) mustEmbedUnimplementedToWorkerServer() {}

// UnsafeToWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToWorkerServer will
// result in compilation errors.
type UnsafeToWorkerServer interface {
	mustEmbedUnimplementedToWorkerServer()
}

func RegisterToWorkerServer(s grpc.ServiceRegistrar, srv ToWorkerServer) {
	s.RegisterService(&ToWorker_ServiceDesc, srv)
}

func _ToWorker_RequestStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToWorkerServer).RequestStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToWorker/RequestStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToWorkerServer).RequestStream(ctx, req.(*StreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToWorker_RequestPremiere_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PremiereRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToWorkerServer).RequestPremiere(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToWorker/RequestPremiere",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToWorkerServer).RequestPremiere(ctx, req.(*PremiereRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToWorker_RequestStreamEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToWorkerServer).RequestStreamEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToWorker/RequestStreamEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToWorkerServer).RequestStreamEnd(ctx, req.(*EndStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToWorker_RequestWaveform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaveformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToWorkerServer).RequestWaveform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToWorker/RequestWaveform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToWorkerServer).RequestWaveform(ctx, req.(*WaveformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToWorker_RequestCut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToWorkerServer).RequestCut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToWorker/RequestCut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToWorkerServer).RequestCut(ctx, req.(*CutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToWorker_GenerateThumbnails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateThumbnailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToWorkerServer).GenerateThumbnails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToWorker/GenerateThumbnails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToWorkerServer).GenerateThumbnails(ctx, req.(*GenerateThumbnailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToWorker_GenerateSectionImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateSectionImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToWorkerServer).GenerateSectionImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToWorker/GenerateSectionImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToWorkerServer).GenerateSectionImages(ctx, req.(*GenerateSectionImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToWorker_DeleteSectionImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSectionImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToWorkerServer).DeleteSectionImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ToWorker/DeleteSectionImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToWorkerServer).DeleteSectionImage(ctx, req.(*DeleteSectionImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ToWorker_ServiceDesc is the grpc.ServiceDesc for ToWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ToWorker",
	HandlerType: (*ToWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestStream",
			Handler:    _ToWorker_RequestStream_Handler,
		},
		{
			MethodName: "RequestPremiere",
			Handler:    _ToWorker_RequestPremiere_Handler,
		},
		{
			MethodName: "RequestStreamEnd",
			Handler:    _ToWorker_RequestStreamEnd_Handler,
		},
		{
			MethodName: "RequestWaveform",
			Handler:    _ToWorker_RequestWaveform_Handler,
		},
		{
			MethodName: "RequestCut",
			Handler:    _ToWorker_RequestCut_Handler,
		},
		{
			MethodName: "GenerateThumbnails",
			Handler:    _ToWorker_GenerateThumbnails_Handler,
		},
		{
			MethodName: "GenerateSectionImages",
			Handler:    _ToWorker_GenerateSectionImages_Handler,
		},
		{
			MethodName: "DeleteSectionImage",
			Handler:    _ToWorker_DeleteSectionImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// FromWorkerClient is the client API for FromWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FromWorkerClient interface {
	// JoinWorkers is a request to the server to join the worker pool.
	JoinWorkers(ctx context.Context, in *JoinWorkersRequest, opts ...grpc.CallOption) (*JoinWorkersResponse, error)
	SendHeartBeat(ctx context.Context, in *HeartBeat, opts ...grpc.CallOption) (*Status, error)
	NotifyTranscodingProgress(ctx context.Context, opts ...grpc.CallOption) (FromWorker_NotifyTranscodingProgressClient, error)
	NotifyTranscodingFinished(ctx context.Context, in *TranscodingFinished, opts ...grpc.CallOption) (*Status, error)
	NotifySilenceResults(ctx context.Context, in *SilenceResults, opts ...grpc.CallOption) (*Status, error)
	NotifyStreamStarted(ctx context.Context, in *StreamStarted, opts ...grpc.CallOption) (*Status, error)
	NotifyStreamFinished(ctx context.Context, in *StreamFinished, opts ...grpc.CallOption) (*Status, error)
	NotifyUploadFinished(ctx context.Context, in *UploadFinished, opts ...grpc.CallOption) (*Status, error)
	NotifyThumbnailsFinished(ctx context.Context, in *ThumbnailsFinished, opts ...grpc.CallOption) (*Status, error)
	SendSelfStreamRequest(ctx context.Context, in *SelfStreamRequest, opts ...grpc.CallOption) (*SelfStreamResponse, error)
	GetStreamInfoForUpload(ctx context.Context, in *GetStreamInfoForUploadRequest, opts ...grpc.CallOption) (*GetStreamInfoForUploadResponse, error)
	NewKeywords(ctx context.Context, in *NewKeywordsRequest, opts ...grpc.CallOption) (*Status, error)
}

type fromWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewFromWorkerClient(cc grpc.ClientConnInterface) FromWorkerClient {
	return &fromWorkerClient{cc}
}

func (c *fromWorkerClient) JoinWorkers(ctx context.Context, in *JoinWorkersRequest, opts ...grpc.CallOption) (*JoinWorkersResponse, error) {
	out := new(JoinWorkersResponse)
	err := c.cc.Invoke(ctx, "/api.FromWorker/JoinWorkers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) SendHeartBeat(ctx context.Context, in *HeartBeat, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.FromWorker/SendHeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) NotifyTranscodingProgress(ctx context.Context, opts ...grpc.CallOption) (FromWorker_NotifyTranscodingProgressClient, error) {
	stream, err := c.cc.NewStream(ctx, &FromWorker_ServiceDesc.Streams[0], "/api.FromWorker/NotifyTranscodingProgress", opts...)
	if err != nil {
		return nil, err
	}
	x := &fromWorkerNotifyTranscodingProgressClient{stream}
	return x, nil
}

type FromWorker_NotifyTranscodingProgressClient interface {
	Send(*NotifyTranscodingProgressRequest) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type fromWorkerNotifyTranscodingProgressClient struct {
	grpc.ClientStream
}

func (x *fromWorkerNotifyTranscodingProgressClient) Send(m *NotifyTranscodingProgressRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fromWorkerNotifyTranscodingProgressClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fromWorkerClient) NotifyTranscodingFinished(ctx context.Context, in *TranscodingFinished, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.FromWorker/NotifyTranscodingFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) NotifySilenceResults(ctx context.Context, in *SilenceResults, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.FromWorker/NotifySilenceResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) NotifyStreamStarted(ctx context.Context, in *StreamStarted, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.FromWorker/NotifyStreamStarted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) NotifyStreamFinished(ctx context.Context, in *StreamFinished, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.FromWorker/NotifyStreamFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) NotifyUploadFinished(ctx context.Context, in *UploadFinished, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.FromWorker/NotifyUploadFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) NotifyThumbnailsFinished(ctx context.Context, in *ThumbnailsFinished, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.FromWorker/NotifyThumbnailsFinished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) SendSelfStreamRequest(ctx context.Context, in *SelfStreamRequest, opts ...grpc.CallOption) (*SelfStreamResponse, error) {
	out := new(SelfStreamResponse)
	err := c.cc.Invoke(ctx, "/api.FromWorker/SendSelfStreamRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) GetStreamInfoForUpload(ctx context.Context, in *GetStreamInfoForUploadRequest, opts ...grpc.CallOption) (*GetStreamInfoForUploadResponse, error) {
	out := new(GetStreamInfoForUploadResponse)
	err := c.cc.Invoke(ctx, "/api.FromWorker/GetStreamInfoForUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fromWorkerClient) NewKeywords(ctx context.Context, in *NewKeywordsRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/api.FromWorker/NewKeywords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FromWorkerServer is the server API for FromWorker service.
// All implementations must embed UnimplementedFromWorkerServer
// for forward compatibility
type FromWorkerServer interface {
	// JoinWorkers is a request to the server to join the worker pool.
	JoinWorkers(context.Context, *JoinWorkersRequest) (*JoinWorkersResponse, error)
	SendHeartBeat(context.Context, *HeartBeat) (*Status, error)
	NotifyTranscodingProgress(FromWorker_NotifyTranscodingProgressServer) error
	NotifyTranscodingFinished(context.Context, *TranscodingFinished) (*Status, error)
	NotifySilenceResults(context.Context, *SilenceResults) (*Status, error)
	NotifyStreamStarted(context.Context, *StreamStarted) (*Status, error)
	NotifyStreamFinished(context.Context, *StreamFinished) (*Status, error)
	NotifyUploadFinished(context.Context, *UploadFinished) (*Status, error)
	NotifyThumbnailsFinished(context.Context, *ThumbnailsFinished) (*Status, error)
	SendSelfStreamRequest(context.Context, *SelfStreamRequest) (*SelfStreamResponse, error)
	GetStreamInfoForUpload(context.Context, *GetStreamInfoForUploadRequest) (*GetStreamInfoForUploadResponse, error)
	NewKeywords(context.Context, *NewKeywordsRequest) (*Status, error)
	mustEmbedUnimplementedFromWorkerServer()
}

// UnimplementedFromWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedFromWorkerServer struct {
}

func (UnimplementedFromWorkerServer) JoinWorkers(context.Context, *JoinWorkersRequest) (*JoinWorkersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinWorkers not implemented")
}
func (UnimplementedFromWorkerServer) SendHeartBeat(context.Context, *HeartBeat) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartBeat not implemented")
}
func (UnimplementedFromWorkerServer) NotifyTranscodingProgress(FromWorker_NotifyTranscodingProgressServer) error {
	return status.Errorf(codes.Unimplemented, "method NotifyTranscodingProgress not implemented")
}
func (UnimplementedFromWorkerServer) NotifyTranscodingFinished(context.Context, *TranscodingFinished) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyTranscodingFinished not implemented")
}
func (UnimplementedFromWorkerServer) NotifySilenceResults(context.Context, *SilenceResults) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifySilenceResults not implemented")
}
func (UnimplementedFromWorkerServer) NotifyStreamStarted(context.Context, *StreamStarted) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyStreamStarted not implemented")
}
func (UnimplementedFromWorkerServer) NotifyStreamFinished(context.Context, *StreamFinished) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyStreamFinished not implemented")
}
func (UnimplementedFromWorkerServer) NotifyUploadFinished(context.Context, *UploadFinished) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyUploadFinished not implemented")
}
func (UnimplementedFromWorkerServer) NotifyThumbnailsFinished(context.Context, *ThumbnailsFinished) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyThumbnailsFinished not implemented")
}
func (UnimplementedFromWorkerServer) SendSelfStreamRequest(context.Context, *SelfStreamRequest) (*SelfStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSelfStreamRequest not implemented")
}
func (UnimplementedFromWorkerServer) GetStreamInfoForUpload(context.Context, *GetStreamInfoForUploadRequest) (*GetStreamInfoForUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStreamInfoForUpload not implemented")
}
func (UnimplementedFromWorkerServer) NewKeywords(context.Context, *NewKeywordsRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewKeywords not implemented")
}
func (UnimplementedFromWorkerServer) mustEmbedUnimplementedFromWorkerServer() {}

// UnsafeFromWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FromWorkerServer will
// result in compilation errors.
type UnsafeFromWorkerServer interface {
	mustEmbedUnimplementedFromWorkerServer()
}

func RegisterFromWorkerServer(s grpc.ServiceRegistrar, srv FromWorkerServer) {
	s.RegisterService(&FromWorker_ServiceDesc, srv)
}

func _FromWorker_JoinWorkers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinWorkersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).JoinWorkers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/JoinWorkers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).JoinWorkers(ctx, req.(*JoinWorkersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_SendHeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).SendHeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/SendHeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).SendHeartBeat(ctx, req.(*HeartBeat))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_NotifyTranscodingProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FromWorkerServer).NotifyTranscodingProgress(&fromWorkerNotifyTranscodingProgressServer{stream})
}

type FromWorker_NotifyTranscodingProgressServer interface {
	SendAndClose(*Status) error
	Recv() (*NotifyTranscodingProgressRequest, error)
	grpc.ServerStream
}

type fromWorkerNotifyTranscodingProgressServer struct {
	grpc.ServerStream
}

func (x *fromWorkerNotifyTranscodingProgressServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fromWorkerNotifyTranscodingProgressServer) Recv() (*NotifyTranscodingProgressRequest, error) {
	m := new(NotifyTranscodingProgressRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FromWorker_NotifyTranscodingFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscodingFinished)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).NotifyTranscodingFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/NotifyTranscodingFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).NotifyTranscodingFinished(ctx, req.(*TranscodingFinished))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_NotifySilenceResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SilenceResults)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).NotifySilenceResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/NotifySilenceResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).NotifySilenceResults(ctx, req.(*SilenceResults))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_NotifyStreamStarted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamStarted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).NotifyStreamStarted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/NotifyStreamStarted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).NotifyStreamStarted(ctx, req.(*StreamStarted))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_NotifyStreamFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamFinished)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).NotifyStreamFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/NotifyStreamFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).NotifyStreamFinished(ctx, req.(*StreamFinished))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_NotifyUploadFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFinished)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).NotifyUploadFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/NotifyUploadFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).NotifyUploadFinished(ctx, req.(*UploadFinished))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_NotifyThumbnailsFinished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThumbnailsFinished)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).NotifyThumbnailsFinished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/NotifyThumbnailsFinished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).NotifyThumbnailsFinished(ctx, req.(*ThumbnailsFinished))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_SendSelfStreamRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelfStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).SendSelfStreamRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/SendSelfStreamRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).SendSelfStreamRequest(ctx, req.(*SelfStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_GetStreamInfoForUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamInfoForUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).GetStreamInfoForUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/GetStreamInfoForUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).GetStreamInfoForUpload(ctx, req.(*GetStreamInfoForUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FromWorker_NewKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FromWorkerServer).NewKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FromWorker/NewKeywords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FromWorkerServer).NewKeywords(ctx, req.(*NewKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FromWorker_ServiceDesc is the grpc.ServiceDesc for FromWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FromWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.FromWorker",
	HandlerType: (*FromWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinWorkers",
			Handler:    _FromWorker_JoinWorkers_Handler,
		},
		{
			MethodName: "SendHeartBeat",
			Handler:    _FromWorker_SendHeartBeat_Handler,
		},
		{
			MethodName: "NotifyTranscodingFinished",
			Handler:    _FromWorker_NotifyTranscodingFinished_Handler,
		},
		{
			MethodName: "NotifySilenceResults",
			Handler:    _FromWorker_NotifySilenceResults_Handler,
		},
		{
			MethodName: "NotifyStreamStarted",
			Handler:    _FromWorker_NotifyStreamStarted_Handler,
		},
		{
			MethodName: "NotifyStreamFinished",
			Handler:    _FromWorker_NotifyStreamFinished_Handler,
		},
		{
			MethodName: "NotifyUploadFinished",
			Handler:    _FromWorker_NotifyUploadFinished_Handler,
		},
		{
			MethodName: "NotifyThumbnailsFinished",
			Handler:    _FromWorker_NotifyThumbnailsFinished_Handler,
		},
		{
			MethodName: "SendSelfStreamRequest",
			Handler:    _FromWorker_SendSelfStreamRequest_Handler,
		},
		{
			MethodName: "GetStreamInfoForUpload",
			Handler:    _FromWorker_GetStreamInfoForUpload_Handler,
		},
		{
			MethodName: "NewKeywords",
			Handler:    _FromWorker_NewKeywords_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NotifyTranscodingProgress",
			Handler:       _FromWorker_NotifyTranscodingProgress_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
