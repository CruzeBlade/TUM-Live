// Code generated by MockGen. DO NOT EDIT.
// Source: subtitles.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/joschahenningsen/TUM-Live/model"
)

// MockSubtitlesDao is a mock of SubtitlesDao interface.
type MockSubtitlesDao struct {
	ctrl     *gomock.Controller
	recorder *MockSubtitlesDaoMockRecorder
}

// MockSubtitlesDaoMockRecorder is the mock recorder for MockSubtitlesDao.
type MockSubtitlesDaoMockRecorder struct {
	mock *MockSubtitlesDao
}

// NewMockSubtitlesDao creates a new mock instance.
func NewMockSubtitlesDao(ctrl *gomock.Controller) *MockSubtitlesDao {
	mock := &MockSubtitlesDao{ctrl: ctrl}
	mock.recorder = &MockSubtitlesDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubtitlesDao) EXPECT() *MockSubtitlesDaoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSubtitlesDao) Create(arg0 context.Context, arg1 *model.Subtitles) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSubtitlesDaoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSubtitlesDao)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockSubtitlesDao) Delete(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSubtitlesDaoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSubtitlesDao)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockSubtitlesDao) Get(arg0 context.Context, arg1 uint) (model.Subtitles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(model.Subtitles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSubtitlesDaoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubtitlesDao)(nil).Get), arg0, arg1)
}
