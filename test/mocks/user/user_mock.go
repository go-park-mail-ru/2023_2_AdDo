// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/user/domain.go

// Package user_mock is a generated GoMock package.
package user_mock

import (
	io "io"
	user_domain "main/internal/pkg/user"
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

// Auth mocks base method.
func (m *MockUseCase) Auth(sessionId string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", sessionId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Auth indicates an expected call of Auth.
func (mr *MockUseCaseMockRecorder) Auth(sessionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockUseCase)(nil).Auth), sessionId)
}

// GetUserInfo mocks base method.
func (m *MockUseCase) GetUserInfo(sessionId string) (user_domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", sessionId)
	ret0, _ := ret[0].(user_domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockUseCaseMockRecorder) GetUserInfo(sessionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockUseCase)(nil).GetUserInfo), sessionId)
}

// GetUserName mocks base method.
func (m *MockUseCase) GetUserName(userId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserName", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserName indicates an expected call of GetUserName.
func (mr *MockUseCaseMockRecorder) GetUserName(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserName", reflect.TypeOf((*MockUseCase)(nil).GetUserName), userId)
}

// Login mocks base method.
func (m *MockUseCase) Login(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUseCaseMockRecorder) Login(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUseCase)(nil).Login), email, password)
}

// Logout mocks base method.
func (m *MockUseCase) Logout(sessionId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", sessionId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockUseCaseMockRecorder) Logout(sessionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUseCase)(nil).Logout), sessionId)
}

// Register mocks base method.
func (m *MockUseCase) Register(user user_domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockUseCaseMockRecorder) Register(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUseCase)(nil).Register), user)
}

// RemoveAvatar mocks base method.
func (m *MockUseCase) RemoveAvatar(userId string) error {
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

// UpdateUserInfo mocks base method.
func (m *MockUseCase) UpdateUserInfo(userId string, user user_domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", userId, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo.
func (mr *MockUseCaseMockRecorder) UpdateUserInfo(userId, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockUseCase)(nil).UpdateUserInfo), userId, user)
}

// UploadAvatar mocks base method.
func (m *MockUseCase) UploadAvatar(userId string, src io.Reader, size int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", userId, src, size)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadAvatar indicates an expected call of UploadAvatar.
func (mr *MockUseCaseMockRecorder) UploadAvatar(userId, src, size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockUseCase)(nil).UploadAvatar), userId, src, size)
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

// CheckEmailAndPassword mocks base method.
func (m *MockRepository) CheckEmailAndPassword(email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckEmailAndPassword", email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckEmailAndPassword indicates an expected call of CheckEmailAndPassword.
func (mr *MockRepositoryMockRecorder) CheckEmailAndPassword(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckEmailAndPassword", reflect.TypeOf((*MockRepository)(nil).CheckEmailAndPassword), email, password)
}

// Create mocks base method.
func (m *MockRepository) Create(user user_domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), user)
}

// GetAvatarPath mocks base method.
func (m *MockRepository) GetAvatarPath(userId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatarPath", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvatarPath indicates an expected call of GetAvatarPath.
func (mr *MockRepositoryMockRecorder) GetAvatarPath(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatarPath", reflect.TypeOf((*MockRepository)(nil).GetAvatarPath), userId)
}

// GetById mocks base method.
func (m *MockRepository) GetById(id string) (user_domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(user_domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRepository)(nil).GetById), id)
}

// GetUserNameById mocks base method.
func (m *MockRepository) GetUserNameById(userId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserNameById", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserNameById indicates an expected call of GetUserNameById.
func (mr *MockRepositoryMockRecorder) GetUserNameById(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserNameById", reflect.TypeOf((*MockRepository)(nil).GetUserNameById), userId)
}

// RemoveAvatarPath mocks base method.
func (m *MockRepository) RemoveAvatarPath(userId string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAvatarPath", userId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAvatarPath indicates an expected call of RemoveAvatarPath.
func (mr *MockRepositoryMockRecorder) RemoveAvatarPath(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAvatarPath", reflect.TypeOf((*MockRepository)(nil).RemoveAvatarPath), userId)
}

// UpdateAvatarPath mocks base method.
func (m *MockRepository) UpdateAvatarPath(userId, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatarPath", userId, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatarPath indicates an expected call of UpdateAvatarPath.
func (mr *MockRepositoryMockRecorder) UpdateAvatarPath(userId, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatarPath", reflect.TypeOf((*MockRepository)(nil).UpdateAvatarPath), userId, path)
}

// UpdateUserInfo mocks base method.
func (m *MockRepository) UpdateUserInfo(user user_domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo.
func (mr *MockRepositoryMockRecorder) UpdateUserInfo(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockRepository)(nil).UpdateUserInfo), user)
}
