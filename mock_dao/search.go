// Code generated by MockGen. DO NOT EDIT.
// Source: search.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/joschahenningsen/TUM-Live/model"
)

// MockSearchDao is a mock of SearchDao interface.
type MockSearchDao struct {
	ctrl     *gomock.Controller
	recorder *MockSearchDaoMockRecorder
}

// MockSearchDaoMockRecorder is the mock recorder for MockSearchDao.
type MockSearchDaoMockRecorder struct {
	mock *MockSearchDao
}

// NewMockSearchDao creates a new mock instance.
func NewMockSearchDao(ctrl *gomock.Controller) *MockSearchDao {
	mock := &MockSearchDao{ctrl: ctrl}
	mock.recorder = &MockSearchDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchDao) EXPECT() *MockSearchDaoMockRecorder {
	return m.recorder
}

// Search mocks base method.
func (m *MockSearchDao) Search(q string, courseId uint) ([]model.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", q, courseId)
	ret0, _ := ret[0].([]model.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockSearchDaoMockRecorder) Search(q, courseId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockSearchDao)(nil).Search), q, courseId)
}