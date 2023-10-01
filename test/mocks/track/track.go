// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/track/domain.go

// Package track_mock is a generated GoMock package.
package track_mock

import (
	track "main/internal/pkg/track"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockUseCase) Add(track track.Track) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", track)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockUseCaseMockRecorder) Add(track interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUseCase)(nil).Add), track)
}

// GetAll mocks base method.
func (m *MockUseCase) GetAll() ([]track.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]track.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUseCaseMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUseCase)(nil).GetAll))
}

// GetFavourite mocks base method.
func (m *MockUseCase) GetFavourite(userId uint64) ([]track.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavourite", userId)
	ret0, _ := ret[0].([]track.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavourite indicates an expected call of GetFavourite.
func (mr *MockUseCaseMockRecorder) GetFavourite(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavourite", reflect.TypeOf((*MockUseCase)(nil).GetFavourite), userId)
}

// GetPopular mocks base method.
func (m *MockUseCase) GetPopular() ([]track.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPopular")
	ret0, _ := ret[0].([]track.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPopular indicates an expected call of GetPopular.
func (mr *MockUseCaseMockRecorder) GetPopular() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPopular", reflect.TypeOf((*MockUseCase)(nil).GetPopular))
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(track track.Track) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", track)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(track interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), track)
}

// GetAll mocks base method.
func (m *MockRepository) GetAll() ([]track.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]track.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepository)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockRepository) GetById(id uint64) (track.Track, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(track.Track)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRepository)(nil).GetById), id)
}

// GetByUserId mocks base method.
func (m *MockRepository) GetByUserId(userId uint64) ([]track.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", userId)
	ret0, _ := ret[0].([]track.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockRepositoryMockRecorder) GetByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockRepository)(nil).GetByUserId), userId)
}
