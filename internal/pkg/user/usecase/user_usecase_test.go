package user_usecase

import (
	"errors"
	user_domain "main/internal/pkg/user"
	session_mock "main/test/mocks/session"
	user_mock "main/test/mocks/user"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user_mock.NewMockRepository(ctrl)

	useCase := &WithStatefulSessions{
		UserRepo: mockUserRepo,
		logger:   logrus.New(),
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
		logger:   logrus.New(),
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
		logger:   logrus.New(),
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
		logger:   logrus.New(),
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
		logger:   logrus.New(),
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
		logger:   logrus.New(),
	}

	const sessionId = "sessionId"

	mockAuthRepo.EXPECT().Delete(sessionId).Return(nil)
	err := useCase.Logout(sessionId)

	assert.Equal(t, nil, err)
}

//func TestUploadAvatar_Success(t *testing.T) {
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockUserRepo := user_mock.NewMockRepository(ctrl)
//	mockAvatarRepo := avatar_mock.NewMockRepository(ctrl)
//	mockAvatarUseCase := avatar_mock.NewMockUseCase(ctrl)
//
//	useCase := &WithStatefulSessions{
//		UserRepo:      mockUserRepo,
//		AvatarRepo:    mockAvatarRepo,
//		AvatarUseCase: mockAvatarUseCase,
//	}
//
//	const (
//		mockUserId = "1"
//		mockPath   = "/user-images/images.png"
//	)
//	mockPayload := new(bytes.Buffer)
//	mockPayloadSize := int64(mockPayload.Len())
//
//	avatar := avatar_domain.Base{
//		Name:        "images-name",
//		Payload:     mockPayload,
//		PayloadSize: mockPayloadSize,
//		ContentType: "images/png",
//	}
//	t.Run("Create images", func(t *testing.T) {
//		mockOldPath := ""
//		mockUserRepo.EXPECT().GetAvatarPath(mockUserId).Return(mockOldPath, nil)
//		mockAvatarUseCase.EXPECT().GetAvatar(mockUserId, mockPayload, mockPayloadSize).Return(avatar, nil)
//		mockAvatarRepo.EXPECT().UploadAvatar(avatar).Return(mockPath, nil)
//		mockUserRepo.EXPECT().UpdateAvatarPath(mockUserId, mockPath)
//
//		_, err := useCase.UploadAvatar(mockUserId, mockPayload, mockPayloadSize)
//
//		assert.Equal(t, nil, err)
//	})
//
//	t.Run("Update images", func(t *testing.T) {
//		mockOldPath := "/user-images/images.png"
//		mockUserRepo.EXPECT().GetAvatarPath(mockUserId).Return(mockOldPath, nil)
//		mockAvatarUseCase.EXPECT().GetAvatar(mockUserId, mockPayload, mockPayloadSize).Return(avatar, nil)
//		mockAvatarRepo.EXPECT().UploadAvatar(avatar).Return(mockPath, nil)
//		mockUserRepo.EXPECT().UpdateAvatarPath(mockUserId, mockPath)
//		mockAvatarRepo.EXPECT().Remove(mockOldPath)
//
//		_, err := useCase.UploadAvatar(mockUserId, mockPayload, mockPayloadSize)
//
//		assert.Equal(t, nil, err)
//	})
//
//}
//
//func TestRemoveAvatar_Success(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockUserRepo := user_mock.NewMockRepository(ctrl)
//	mockAvatarRepo := avatar_mock.NewMockRepository(ctrl)
//
//	useCase := &WithStatefulSessions{
//		UserRepo:   mockUserRepo,
//		AvatarRepo: mockAvatarRepo,
//	}
//
//	const (
//		mockUserID  = "1"
//		mockOldPath = "/user_avatar/images.png"
//	)
//
//	mockUserRepo.EXPECT().GetAvatarPath(mockUserID).Return(mockOldPath, nil)
//	mockAvatarRepo.EXPECT().Remove(mockOldPath).Return(nil)
//	mockUserRepo.EXPECT().RemoveAvatarPath(mockUserID).Return(nil)
//
//	err := useCase.RemoveAvatar(mockUserID)
//
//	assert.Equal(t, nil, err)
//}
//
//func TestRemoveAvatar_AvatarDoesNotExist(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockUserRepo := user_mock.NewMockRepository(ctrl)
//
//	useCase := &WithStatefulSessions{
//		UserRepo: mockUserRepo,
//	}
//
//	const (
//		mockUserID  = "1"
//		mockOldPath = ""
//	)
//
//	mockUserRepo.EXPECT().GetAvatarPath(mockUserID).Return(mockOldPath, nil)
//
//	err := useCase.RemoveAvatar(mockUserID)
//
//	assert.Equal(t, avatar_domain.ErrAvatarDoesNotExist, err)
//}
