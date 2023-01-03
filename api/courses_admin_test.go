package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/joschahenningsen/TUM-Live/dao"
	"github.com/joschahenningsen/TUM-Live/mock_dao"
	"github.com/joschahenningsen/TUM-Live/model"
	"github.com/joschahenningsen/TUM-Live/tools"
	"github.com/joschahenningsen/TUM-Live/tools/testutils"
	"github.com/matthiasreumann/gomino"
	"net/http"
	"testing"
	"time"
)

func CourseAdminRouterWrapper(r *gin.Engine) {
	configGinCourseAdminRouter(r, dao.DaoWrapper{})
}

func TestCoursesCRUDAdmin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("DELETE/api/course/:courseID", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/", testutils.CourseFPV.ID)

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(), testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
		}.Method(http.MethodDelete).Url(url).Run(t, testutils.Equal)
	})
}

func TestCoursesLectureActions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("POST/api/course/:courseID/createLecture", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/createLecture", testutils.CourseFPV.ID)

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid body": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         nil,
				ExpectedCode: http.StatusBadRequest,
			},
			"lectureHallId set on 'premiere'": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares: testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body: createLectureRequest{
					LectureHallId: "1",
					Premiere:      true,
				},
				ExpectedCode: http.StatusBadRequest,
			},
			"invalid lectureHallId": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares: testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body: createLectureRequest{
					Title:         "Lecture 1",
					LectureHallId: "abc",
					Start:         time.Now(),
					Duration:      90,
					Premiere:      false,
					Vodup:         false,
					DateSeries:    []time.Time{},
				},
				ExpectedCode: http.StatusBadRequest,
			},
			"can not update course": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(), testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								UpdateCourse(gomock.Any(), gomock.Any()).
								Return(errors.New(""))
							return coursesMock
						}(),
						AuditDao: func() dao.AuditDao {
							auditMock := mock_dao.NewMockAuditDao(gomock.NewController(t))
							auditMock.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()
							return auditMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares: testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body: createLectureRequest{
					Title:         "Lecture 1",
					LectureHallId: "1",
					Start:         time.Now(),
					Duration:      90,
					Premiere:      false,
					Vodup:         false,
					DateSeries: []time.Time{
						time.Now(),
					},
				},
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(), testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								UpdateCourse(gomock.Any(), gomock.Any()).
								Return(nil)
							return coursesMock
						}(),
						AuditDao: func() dao.AuditDao {
							auditMock := mock_dao.NewMockAuditDao(gomock.NewController(t))
							auditMock.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()
							return auditMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares: testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body: createLectureRequest{
					Title:         "Lecture 1",
					LectureHallId: "1",
					Start:         time.Now(),
					Duration:      90,
					Premiere:      false,
					Vodup:         false,
					DateSeries: []time.Time{
						time.Now(),
					},
				},
				ExpectedCode: http.StatusOK,
			}}.
			Method(http.MethodPost).
			Url(url).
			Run(t, testutils.Equal)
	})
	t.Run("POST/api/course/:courseID/deleteLecture", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/deleteLectures", testutils.CourseFPV.ID)

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid body": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         nil,
				ExpectedCode: http.StatusBadRequest,
			},
			"invalid stream id in body": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamGBSLive.ID)).
								Return(testutils.StreamGBSLive, nil).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares: testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body: deleteLecturesRequest{StreamIDs: []string{
					fmt.Sprintf("%d", testutils.StreamGBSLive.ID)},
				},
				ExpectedCode: http.StatusForbidden,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								DeleteStream(fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return().
								AnyTimes()
							return streamsMock
						}(),
						AuditDao: testutils.GetAuditMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares: testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body: deleteLecturesRequest{StreamIDs: []string{
					fmt.Sprintf("%d", testutils.StreamFPVLive.ID)},
				},
				ExpectedCode: http.StatusOK,
			}}.Method(http.MethodPost).Url(url).Run(t, testutils.Equal)
	})
	t.Run("POST/api/course/:courseID/renameLecture/:streamID", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/renameLecture/%d", testutils.CourseFPV.ID, testutils.StreamFPVLive.ID)

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid streamID": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Url:          fmt.Sprintf("/api/course/%d/renameLecture/abc", testutils.CourseFPV.ID),
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"invalid body": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         nil,
				ExpectedCode: http.StatusBadRequest,
			},
			"stream not found": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(model.Stream{}, errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares: testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body: renameLectureRequest{
					Name: "Proofs #1",
				},
				ExpectedCode: http.StatusNotFound,
			},
			"can not update stream": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								UpdateStream(gomock.Any()).
								Return(errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares: testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body: renameLectureRequest{
					Name: "Proofs #1",
				},
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								UpdateStream(gomock.Any()).
								Return(nil).
								AnyTimes()
							return streamsMock
						}(),
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares: testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body: renameLectureRequest{
					Name: "Proofs #1",
				},
				ExpectedCode: http.StatusOK,
			}}.Method(http.MethodPost).Url(url).Run(t, testutils.Equal)
	})
	t.Run("POST/api/course/:courseID/updateLectureSeries/:streamID", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/updateLectureSeries/%d", testutils.CourseFPV.ID, testutils.StreamFPVLive.ID)

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"stream not found": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusNotFound,
			},
			"can not update lecture series": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								UpdateLectureSeries(testutils.StreamFPVLive).
								Return(errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								UpdateLectureSeries(testutils.StreamFPVLive).
								Return(nil).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusOK,
			}}.
			Method(http.MethodPost).
			Url(url).
			Run(t, testutils.Equal)
	})
	t.Run("DELETE/api/course/:courseID/deleteLectureSeries/:streamID", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/deleteLectureSeries/%d", testutils.CourseFPV.ID, testutils.StreamFPVLive.ID)

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"stream not found": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusNotFound,
			},
			"invalid series-identifier": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), gomock.Any()).
								Return(testutils.StreamGBSLive, nil). //StreamGBSLive.SeriesIdentifier == ""
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"can not delete lecture-series": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						AuditDao:   testutils.GetAuditMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), gomock.Any()).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								DeleteLectureSeries(testutils.StreamFPVLive.SeriesIdentifier).
								Return(errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						AuditDao:   testutils.GetAuditMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), gomock.Any()).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								DeleteLectureSeries(testutils.StreamFPVLive.SeriesIdentifier).
								Return(nil).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusOK,
			}}.
			Method(http.MethodDelete).
			Url(url).
			Run(t, testutils.Equal)
	})
	t.Run("PUT/api/course/:courseID/updateDescription/:streamID", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/updateDescription/%d", testutils.CourseFPV.ID, testutils.StreamFPVLive.ID)

		body := renameLectureRequest{
			Name: "New lecture name!",
		}
		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid streamID": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Url:          fmt.Sprintf("/api/course/%d/updateDescription/abc", testutils.CourseFPV.ID),
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"invalid body": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"can not find stream": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         body,
				ExpectedCode: http.StatusNotFound,
			},
			"can not update stream": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								UpdateStream(gomock.Any()).
								Return(errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         body,
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: testutils.GetStreamMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         body,
				ExpectedCode: http.StatusOK,
			}}.
			Method(http.MethodPut).
			Url(url).
			Run(t, testutils.Equal)
	})
}

func TestUnits(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("POST/api/course/:courseID/addUnit", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/addUnit", testutils.CourseFPV.ID)

		request := addUnitRequest{
			LectureID:   testutils.StreamFPVLive.ID,
			From:        0,
			To:          42,
			Title:       "New Unit",
			Description: "This is a new one!",
		}

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid body": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"can not find stream": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         request,
				ExpectedCode: http.StatusNotFound,
			},
			"can not update stream associations": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								UpdateStreamFullAssoc(gomock.Any()).
								Return(errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         request,
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: testutils.GetStreamMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         request,
				ExpectedCode: http.StatusOK,
			}}.Method(http.MethodPost).Url(url).Run(t, testutils.Equal)
	})
	t.Run("POST/api/course/:courseID/deleteUnit/:unitID", func(t *testing.T) {
		unit := testutils.StreamFPVLive.Units[0]
		url := fmt.Sprintf("/api/course/%d/deleteUnit/%d",
			testutils.CourseFPV.ID, unit.ID)

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"can not find unit": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetUnitByID(fmt.Sprintf("%d", unit.ID)).
								Return(unit, errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusNotFound,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: testutils.GetStreamMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusOK,
			}}.Method(http.MethodPost).Url(url).Run(t, testutils.Equal)
	})
}

func TestCuts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("POST/api/course/:courseID/submitCut", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/submitCut", testutils.CourseFPV.ID)

		request := submitCutRequest{
			LectureID: testutils.StreamFPVLive.ID,
			From:      0,
			To:        1000,
		}
		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid body": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         nil,
				ExpectedCode: http.StatusBadRequest,
			},
			"can not find stream": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         request,
				ExpectedCode: http.StatusNotFound,
			},
			"can not update stream": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(gomock.NewController(t))
							streamsMock.
								EXPECT().
								GetStreamByID(gomock.Any(), fmt.Sprintf("%d", testutils.StreamFPVLive.ID)).
								Return(testutils.StreamFPVLive, nil).
								AnyTimes()
							streamsMock.
								EXPECT().
								SaveStream(gomock.Any()).
								Return(errors.New("")).
								AnyTimes()
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         request,
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: testutils.GetStreamMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         request,
				ExpectedCode: http.StatusOK,
			}}.Method(http.MethodPost).Url(url).Run(t, testutils.Equal)
	})
}

func TestAdminFunctions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("GET/api/course/:courseID/admins", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/admins", testutils.CourseFPV.ID)

		response := []userForLecturerDto{
			{
				ID:    testutils.Admin.ID,
				Name:  testutils.Admin.Name,
				Login: testutils.Admin.GetLoginString(),
			},
			{
				ID:    testutils.Admin.ID,
				Name:  testutils.Admin.Name,
				Login: testutils.Admin.GetLoginString(),
			},
		}
		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"can not get course admins": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(),
									testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm,
									testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseAdmins(testutils.CourseFPV.ID).
								Return([]model.User{}, errors.New("")).
								AnyTimes()
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:      testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode:     http.StatusOK,
				ExpectedResponse: response,
			}}.Method(http.MethodGet).Url(url).Run(t, testutils.Equal)
	})
	t.Run("PUT/api/course/:courseID/admins/:userID", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/admins/%d", testutils.CourseFPV.ID, testutils.Admin.ID)
		urlStudent := fmt.Sprintf("/api/course/%d/admins/%d", testutils.CourseFPV.ID, testutils.Student.ID)

		resAdmin := userForLecturerDto{
			ID:    testutils.Admin.ID,
			Name:  testutils.Admin.Name,
			Login: testutils.Admin.GetLoginString(),
		}

		resStudent := userForLecturerDto{
			ID:    testutils.Student.ID,
			Name:  testutils.Student.Name,
			Login: testutils.Student.GetLoginString(),
		}

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Method:       http.MethodPut,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid userID": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Url:          fmt.Sprintf("/api/course/%d/admins/abc", testutils.CourseFPV.ID),
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"user not found": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						UsersDao: func() dao.UsersDao {
							usersMock := mock_dao.NewMockUsersDao(gomock.NewController(t))
							usersMock.
								EXPECT().
								GetUserByID(gomock.Any(), testutils.Admin.ID).
								Return(testutils.Admin, errors.New("")).AnyTimes()
							return usersMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusNotFound,
			},
			"can not add admin to course": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(),
									testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								AddAdminToCourse(testutils.Admin.ID, testutils.CourseFPV.ID).
								Return(errors.New("")).
								AnyTimes()
							return coursesMock
						}(),
						UsersDao: testutils.GetUsersMock(t),
						AuditDao: testutils.GetAuditMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
			"can not update user": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(),
									testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								AddAdminToCourse(testutils.Student.ID, testutils.CourseFPV.ID).
								Return(nil).
								AnyTimes()
							return coursesMock
						}(),
						UsersDao: func() dao.UsersDao {
							usersMock := mock_dao.NewMockUsersDao(gomock.NewController(t))
							usersMock.
								EXPECT().
								GetUserByID(gomock.Any(), testutils.Student.ID).
								Return(testutils.Student, nil).
								AnyTimes()
							usersMock.
								EXPECT().
								UpdateUser(gomock.Any()).
								Return(errors.New("")).
								AnyTimes()
							return usersMock
						}(),
						AuditDao: testutils.GetAuditMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Url:          urlStudent,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						UsersDao:   testutils.GetUsersMock(t),
						AuditDao:   testutils.GetAuditMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:      testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode:     http.StatusOK,
				ExpectedResponse: resAdmin,
			},
			"success, user not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(),
									testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								AddAdminToCourse(testutils.Student.ID, testutils.CourseFPV.ID).
								Return(nil).
								AnyTimes()
							return coursesMock
						}(),
						UsersDao: func() dao.UsersDao {
							usersMock := mock_dao.NewMockUsersDao(gomock.NewController(t))
							usersMock.
								EXPECT().
								GetUserByID(gomock.Any(), testutils.Student.ID).
								Return(testutils.Student, nil).
								AnyTimes()
							usersMock.
								EXPECT().
								UpdateUser(gomock.Any()).
								Return(nil).
								AnyTimes()
							return usersMock
						}(),
						AuditDao: testutils.GetAuditMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Url:              urlStudent,
				Middlewares:      testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode:     http.StatusOK,
				ExpectedResponse: resStudent,
			}}.
			Method(http.MethodPut).
			Url(url).
			Run(t, testutils.Equal)
	})
	t.Run("DELETE/api/course/:courseID/admins/:userID", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/admins/%d", testutils.CourseFPV.ID, testutils.Admin.ID)

		response := userForLecturerDto{
			ID:    testutils.Admin.ID,
			Name:  testutils.Admin.Name,
			Login: testutils.Admin.GetLoginString(),
		}

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid userID": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Url:          fmt.Sprintf("/api/course/%d/admins/abc", testutils.CourseFPV.ID),
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"can not get course admins": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(),
									testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseAdmins(testutils.CourseFPV.ID).
								Return([]model.User{}, errors.New("")).
								AnyTimes()
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
			"remove last admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(),
									testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseAdmins(testutils.CourseFPV.ID).
								Return([]model.User{testutils.Admin}, nil).
								AnyTimes()
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"invalid delete request": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(),
									testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseAdmins(testutils.CourseFPV.ID).
								Return([]model.User{testutils.Student}, nil). // student.id != admin.id from url
								AnyTimes()
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"can not remove admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						AuditDao: testutils.GetAuditMock(t),
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseBySlugYearAndTerm(gomock.Any(),
									testutils.CourseFPV.Slug, testutils.CourseFPV.TeachingTerm, testutils.CourseFPV.Year).
								Return(testutils.CourseFPV, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								GetCourseAdmins(testutils.CourseFPV.ID).
								Return([]model.User{testutils.Admin, testutils.Admin}, nil).
								AnyTimes()
							coursesMock.
								EXPECT().
								RemoveAdminFromCourse(testutils.Admin.ID, testutils.CourseFPV.ID).
								Return(errors.New("")).
								AnyTimes()
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						AuditDao:   testutils.GetAuditMock(t),
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:      testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode:     http.StatusOK,
				ExpectedResponse: response,
			}}.
			Method(http.MethodDelete).
			Url(url).
			Run(t, testutils.Equal)
	})
}

func TestPresets(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("POST/api/course/:courseID/presets", func(t *testing.T) {
		url := fmt.Sprintf("/api/course/%d/presets", testutils.CourseFPV.ID)

		var sourceMode model.SourceMode = 1
		selectedPreset := 1
		request := []lhResp{
			{
				LectureHallName: "HS-4",
				LectureHallID:   testutils.LectureHall.ID,
				Presets: []model.CameraPreset{
					{
						Name:          "Preset 1",
						PresetID:      1,
						Image:         "375ed239-c37d-450e-9d4f-1fbdb5a2dec5.jpg",
						LectureHallID: testutils.LectureHall.ID,
						IsDefault:     false,
					},
				},
				SourceMode:       sourceMode,
				SelectedPresetID: selectedPreset,
			},
		}

		presetSettings := []model.CameraPresetPreference{
			{
				LectureHallID: testutils.LectureHall.ID,
				PresetID:      selectedPreset,
			},
		}

		sourceSettings := []model.SourcePreference{
			{
				LectureHallID: testutils.LectureHall.ID,
				SourceMode:    sourceMode,
			},
		}

		afterChanges := testutils.CourseFPV
		afterChanges.SetCameraPresetPreference(presetSettings)
		afterChanges.SetSourcePreference(sourceSettings)

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid body": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"can not update course": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						AuditDao: testutils.GetAuditMock(t),
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()

							coursesMock.
								EXPECT().
								UpdateCourse(gomock.Any(), afterChanges).
								Return(errors.New(""))
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         request,
				ExpectedCode: http.StatusInternalServerError,
			},
			"success": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						AuditDao: testutils.GetAuditMock(t),
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(gomock.NewController(t))
							coursesMock.
								EXPECT().
								GetCourseById(gomock.Any(), testutils.CourseFPV.ID).
								Return(testutils.CourseFPV, nil).
								AnyTimes()

							coursesMock.
								EXPECT().
								UpdateCourse(gomock.Any(), afterChanges).
								Return(nil)
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				Body:         request,
				ExpectedCode: http.StatusOK,
			}}.
			Method(http.MethodPost).
			Url(url).
			Run(t, testutils.Equal)
	})
}

func TestUploadVOD(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("POST/api/course/:courseID/uploadVOD", func(t *testing.T) {
		baseUrl := fmt.Sprintf("/api/course/%d/uploadVOD", testutils.CourseFPV.ID)
		url := fmt.Sprintf("%s?start=2022-07-04T10:00:00.000Z&title=VOD1", baseUrl)

		ctrl := gomock.NewController(t)

		gomino.TestCases{
			"no context": {
				Router:       CourseAdminRouterWrapper,
				Url:          baseUrl,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler),
				ExpectedCode: http.StatusInternalServerError,
			},
			"not admin": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Url:          baseUrl,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			},
			"invalid query": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Url:          baseUrl,
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusBadRequest,
			},
			"can not create stream": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(ctrl)
							streamsMock.
								EXPECT().
								CreateStream(gomock.Any()).
								Return(errors.New(""))
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
			"can note create upload key": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: testutils.GetCoursesMock(t),
						StreamsDao: func() dao.StreamsDao {
							streamsMock := mock_dao.NewMockStreamsDao(ctrl)
							streamsMock.
								EXPECT().
								CreateStream(gomock.Any()).
								Return(nil)
							return streamsMock
						}(),
						UploadKeyDao: func() dao.UploadKeyDao {
							streamsMock := mock_dao.NewMockUploadKeyDao(ctrl)
							streamsMock.
								EXPECT().
								CreateUploadKey(gomock.Any(), gomock.Any()).
								Return(errors.New(""))
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
			"no workers available": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao:   testutils.GetCoursesMock(t),
						StreamsDao:   testutils.GetStreamMock(t),
						UploadKeyDao: testutils.GetUploadKeyMock(t),
						WorkerDao: func() dao.WorkerDao {
							streamsMock := mock_dao.NewMockWorkerDao(ctrl)
							streamsMock.
								EXPECT().
								GetAliveWorkers().
								Return([]model.Worker{})
							return streamsMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode: http.StatusInternalServerError,
			},
		}.
			Method(http.MethodPost).
			Url(url).
			Run(t, testutils.Equal)
	})
}

func TestGetTranscodingProgress(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Run("GET /api/course/:id/stream/:id/transcodingProgress", func(t *testing.T) {
		gomino.TestCases{
			"Admin, OK": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						StreamsDao: func() dao.StreamsDao {
							smock := mock_dao.NewMockStreamsDao(ctrl)
							smock.EXPECT().GetStreamByID(gomock.Any(), "1969").MinTimes(1).MaxTimes(1).Return(testutils.StreamFPVNotLive, nil)
							smock.EXPECT().GetTranscodingProgressByVersion(model.COMB, uint(1969)).MinTimes(1).MaxTimes(1).Return(model.TranscodingProgress{Progress: 69}, nil)
							return smock
						}(),
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(ctrl)
							coursesMock.EXPECT().GetCourseById(gomock.Any(), uint(40)).MinTimes(1).MaxTimes(1).Return(testutils.CourseFPV, nil)
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:      testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextAdmin)),
				ExpectedCode:     http.StatusOK,
				ExpectedResponse: "69",
			},
			"Student, Forbidden": {
				Router: func(r *gin.Engine) {
					wrapper := dao.DaoWrapper{
						CoursesDao: func() dao.CoursesDao {
							coursesMock := mock_dao.NewMockCoursesDao(ctrl)
							coursesMock.EXPECT().GetCourseById(gomock.Any(), uint(40)).MinTimes(1).MaxTimes(1).Return(testutils.CourseFPV, nil)
							return coursesMock
						}(),
					}
					configGinCourseAdminRouter(r, wrapper)
				},
				Middlewares:  testutils.GetMiddlewares(tools.ErrorHandler, testutils.TUMLiveContext(testutils.TUMLiveContextStudent)),
				ExpectedCode: http.StatusForbidden,
			}}.
			Method(http.MethodGet).
			Url("/api/course/40/stream/1969/transcodingProgress").
			Run(t, testutils.Equal)
	})
}