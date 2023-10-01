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
		Id:       1,
		Email:    "user@example.com",
		Username: "user1",
	}

	mockUserRepo.EXPECT().Create(mockUser).Return(mockUser.Id, nil)

	id, err := useCase.Register(mockUser)

	assert.NoError(t, err)
	assert.Equal(t, mockUser.Id, id)
}

func TestRegister_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user_mock.NewMockRepository(ctrl)

	useCase := &WithStatefulSessions{
		UserRepo: mockUserRepo,
	}

	mockUser := user_domain.User{
		Id:       1,
		Email:    "user@example.com",
		Username: "user1",
	}

	// Ожидаем, что метод Create вызовется с аргументом mockUser и вернет ошибку
	mockUserRepo.EXPECT().Create(mockUser).Return(uint64(0), errors.New("user already exists"))

	id, err := useCase.Register(mockUser)

	assert.Error(t, err)
	assert.Equal(t, user_domain.ErrUserAlreadyExist, err)
	assert.Equal(t, uint64(0), id)
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
		Id:       1,
		Email:    "user@example.com",
		Password: "my_password",
	}
	mockSessionId := "mockSessId"

	mockUserRepo.EXPECT().CheckEmailAndPassword(mockUser.Email, mockUser.Password).Return(uint64(mockUser.Id), nil)
	mockAuthRepo.EXPECT().Create(mockUser.Id).Return(mockSessionId, nil)
	id, session, err := useCase.Login(mockUser.Email, mockUser.Password)

	assert.Equal(t, nil, err)
	assert.Equal(t, mockUser.Id, id)
	assert.Equal(t, mockSessionId, session)
}

func TestGetUserInfo_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user_mock.NewMockRepository(ctrl)

	useCase := &WithStatefulSessions{
		UserRepo: mockUserRepo,
	}

	mockUser := user_domain.User{
		Id:       1,
		Email:    "user@example.com",
		Password: "my_password",
	}

	mockUserRepo.EXPECT().GetById(mockUser.Id).Return(mockUser, nil)
	user, err := useCase.GetUserInfo(mockUser.Id)

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
		Id:       1,
		Email:    "user@example.com",
		Password: "my_password",
	}
	mockSessionId := "mockSessId"

	mockAuthRepo.EXPECT().GetByUserId(mockUser.Id).Return(mockSessionId, nil)
	isAuth, err := useCase.Auth(mockUser.Id, mockSessionId)

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

	mockUser := user_domain.User{
		Id:       1,
		Email:    "user@example.com",
		Password: "my_password",
	}

	mockAuthRepo.EXPECT().DeleteByUserId(mockUser.Id).Return(nil)
	err := useCase.Logout(mockUser.Id)

	assert.Equal(t, nil, err)
}
