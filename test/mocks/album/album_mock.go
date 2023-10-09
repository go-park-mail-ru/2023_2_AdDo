// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/album/domain.go

// Package album_mock is a generated GoMock package.
package album_mock

import (
	album "main/internal/pkg/album"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

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

// GetByArtistId mocks base method.
func (m *MockRepository) GetByArtistId(artistId uint64) ([]album.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByArtistId", artistId)
	ret0, _ := ret[0].([]album.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByArtistId indicates an expected call of GetByArtistId.
func (mr *MockRepositoryMockRecorder) GetByArtistId(artistId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByArtistId", reflect.TypeOf((*MockRepository)(nil).GetByArtistId), artistId)
}

// GetByTrackId mocks base method.
func (m *MockRepository) GetByTrackId(trackId uint64) ([]album.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTrackId", trackId)
	ret0, _ := ret[0].([]album.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTrackId indicates an expected call of GetByTrackId.
func (mr *MockRepositoryMockRecorder) GetByTrackId(trackId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTrackId", reflect.TypeOf((*MockRepository)(nil).GetByTrackId), trackId)
}
