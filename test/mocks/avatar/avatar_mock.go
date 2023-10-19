// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/avatar/domain.go

// Package avatar_mock is a generated GoMock package.
package avatar_mock

import (
	io "io"
	avatar_domain "main/internal/pkg/avatar"
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

// RemoveAvatar mocks base method.
func (m *MockUseCase) RemoveAvatar(userId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAvatar", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAvatar indicates an expected call of RemoveAvatar.
func (mr *MockUseCaseMockRecorder) RemoveAvatar(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAvatar", reflect.TypeOf((*MockUseCase)(nil).RemoveAvatar), userId)
}

// UploadAvatar mocks base method.
func (m *MockUseCase) UploadAvatar(userId uint64, src io.Reader, size int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", userId, src, size)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAvatar indicates an expected call of UploadAvatar.
func (mr *MockUseCaseMockRecorder) UploadAvatar(userId, src, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockUseCase)(nil).UploadAvatar), userId, src, size)
}

// MockS3Repository is a mock of S3Repository interface.
type MockS3Repository struct {
	ctrl     *gomock.Controller
	recorder *MockS3RepositoryMockRecorder
}

// MockS3RepositoryMockRecorder is the mock recorder for MockS3Repository.
type MockS3RepositoryMockRecorder struct {
	mock *MockS3Repository
}

// NewMockS3Repository creates a new mock instance.
func NewMockS3Repository(ctrl *gomock.Controller) *MockS3Repository {
	mock := &MockS3Repository{ctrl: ctrl}
	mock.recorder = &MockS3RepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3Repository) EXPECT() *MockS3RepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockS3Repository) Create(avatar avatar_domain.Avatar) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", avatar)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockS3RepositoryMockRecorder) Create(avatar interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockS3Repository)(nil).Create), avatar)
}

// Remove mocks base method.
func (m *MockS3Repository) Remove(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockS3RepositoryMockRecorder) Remove(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockS3Repository)(nil).Remove), path)
}

// MockDbRepository is a mock of DbRepository interface.
type MockDbRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDbRepositoryMockRecorder
}

// MockDbRepositoryMockRecorder is the mock recorder for MockDbRepository.
type MockDbRepositoryMockRecorder struct {
	mock *MockDbRepository
}

// NewMockDbRepository creates a new mock instance.
func NewMockDbRepository(ctrl *gomock.Controller) *MockDbRepository {
	mock := &MockDbRepository{ctrl: ctrl}
	mock.recorder = &MockDbRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDbRepository) EXPECT() *MockDbRepositoryMockRecorder {
	return m.recorder
}

// GetAvatarPath mocks base method.
func (m *MockDbRepository) GetAvatarPath(userId uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatarPath", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvatarPath indicates an expected call of GetAvatarPath.
func (mr *MockDbRepositoryMockRecorder) GetAvatarPath(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatarPath", reflect.TypeOf((*MockDbRepository)(nil).GetAvatarPath), userId)
}

// RemoveAvatarPath mocks base method.
func (m *MockDbRepository) RemoveAvatarPath(userId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAvatarPath", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAvatarPath indicates an expected call of RemoveAvatarPath.
func (mr *MockDbRepositoryMockRecorder) RemoveAvatarPath(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAvatarPath", reflect.TypeOf((*MockDbRepository)(nil).RemoveAvatarPath), userId)
}

// UpdateAvatarPath mocks base method.
func (m *MockDbRepository) UpdateAvatarPath(userId uint64, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatarPath", userId, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatarPath indicates an expected call of UpdateAvatarPath.
func (mr *MockDbRepositoryMockRecorder) UpdateAvatarPath(userId, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatarPath", reflect.TypeOf((*MockDbRepository)(nil).UpdateAvatarPath), userId, path)
}
