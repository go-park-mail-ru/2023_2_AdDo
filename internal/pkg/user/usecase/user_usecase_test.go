package user_usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	user_domain "main/internal/pkg/user"
	session_mock "main/test/mocks/session"
	user_mock "main/test/mocks/user"
	"testing"
)

func TestRegister_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user_mock.NewMockRepository(ctrl)

	useCase := &WithStatefulSessions{
		UserRepo: mockUserRepo,
	}

	mockUser := user_domain.User{
		Id:       "qwer-qwer-qwer",
		Email:    "user@example.com",
		Username: "user1",
	}

	mockUserRepo.EXPECT().Create(mockUser).Return(nil)

	err := useCase.Register(mockUser)

	assert.NoError(t, err)
}

func TestRegister_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user_mock.NewMockRepository(ctrl)

	useCase := &WithStatefulSessions{
		UserRepo: mockUserRepo,
	}

	mockUser := user_domain.User{
		Id:       "qwer-qwer-qwer",
		Email:    "user@example.com",
		Username: "user1",
	}

	mockUserRepo.EXPECT().Create(mockUser).Return(errors.New("user already exists"))

	err := useCase.Register(mockUser)

	assert.Error(t, err)
}

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user_mock.NewMockRepository(ctrl)
	mockAuthRepo := session_mock.NewMockRepository(ctrl)

	useCase := &WithStatefulSessions{
		UserRepo: mockUserRepo,
		AuthRepo: mockAuthRepo,
	}

	mockUser := user_domain.User{
		Id:       "qwer-qwer-qwer",
		Email:    "user@example.com",
		Password: "my_password",
	}
	mockSessionId := "mockSessId"

	mockUserRepo.EXPECT().CheckEmailAndPassword(mockUser.Email, mockUser.Password).Return(mockUser.Id, nil)
	mockAuthRepo.EXPECT().Create(mockUser.Id).Return(mockSessionId, nil)
	session, err := useCase.Login(mockUser.Email, mockUser.Password)

	assert.Equal(t, nil, err)
	assert.Equal(t, mockSessionId, session)
}

func TestGetUserInfo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user_mock.NewMockRepository(ctrl)
	mockAuthRepo := session_mock.NewMockRepository(ctrl)

	useCase := &WithStatefulSessions{
		UserRepo: mockUserRepo,
		AuthRepo: mockAuthRepo,
	}

	mockUser := user_domain.User{
		Id:       "qwer-qwer-qwer",
		Email:    "user@example.com",
		Password: "my_password",
	}
	const sessionId = "sessionId"

	mockAuthRepo.EXPECT().Get(sessionId).Return(mockUser.Id, nil)
	mockUserRepo.EXPECT().GetById(mockUser.Id).Return(mockUser, nil)
	user, err := useCase.GetUserInfo(sessionId)

	assert.Equal(t, nil, err)
	assert.Equal(t, mockUser, user)
}

func TestAuth_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := session_mock.NewMockRepository(ctrl)

	useCase := &WithStatefulSessions{
		AuthRepo: mockAuthRepo,
	}

	mockUser := user_domain.User{
		Id:       "qwer-qwer-qwer",
		Email:    "user@example.com",
		Password: "my_password",
	}
	const sessionId = "sessionId"

	mockAuthRepo.EXPECT().Get(sessionId).Return(mockUser.Id, nil)
	isAuth, err := useCase.Auth(sessionId)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, isAuth)
}

func TestLogOut_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthRepo := session_mock.NewMockRepository(ctrl)

	useCase := &WithStatefulSessions{
		AuthRepo: mockAuthRepo,
	}

	const sessionId = "sessionId"

	mockAuthRepo.EXPECT().Delete(sessionId).Return(nil)
	err := useCase.Logout(sessionId)

	assert.Equal(t, nil, err)
}
