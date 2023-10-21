// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/avatar/domain.go

// Package avatar_mock is a generated GoMock package.
package avatar_mock

import (
	io "io"
	avatar "main/internal/pkg/avatar"
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

// GetAvatar mocks base method.
func (m *MockUseCase) GetAvatar(id uint64, src io.Reader, size int64) (avatar.Avatar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatar", id, src, size)
	ret0, _ := ret[0].(avatar.Avatar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvatar indicates an expected call of GetAvatar.
func (mr *MockUseCaseMockRecorder) GetAvatar(id, src, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatar", reflect.TypeOf((*MockUseCase)(nil).GetAvatar), id, src, size)
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

// Remove mocks base method.
func (m *MockRepository) Remove(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockRepositoryMockRecorder) Remove(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockRepository)(nil).Remove), path)
}

// UploadAvatar mocks base method.
func (m *MockRepository) UploadAvatar(avatar avatar.Avatar) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", avatar)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadAvatar indicates an expected call of UploadAvatar.
func (mr *MockRepositoryMockRecorder) UploadAvatar(avatar interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockRepository)(nil).UploadAvatar), avatar)
}

// UploadPlaylistImage mocks base method.
func (m *MockRepository) UploadPlaylistImage(avatar avatar.Avatar) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadPlaylistImage", avatar)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadPlaylistImage indicates an expected call of UploadPlaylistImage.
func (mr *MockRepositoryMockRecorder) UploadPlaylistImage(avatar interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPlaylistImage", reflect.TypeOf((*MockRepository)(nil).UploadPlaylistImage), avatar)
}
