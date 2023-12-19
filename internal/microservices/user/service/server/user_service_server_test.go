package grpc_server_user

import (
	"context"
	"errors"
	image_proto "main/internal/microservices/image/proto"
	session_proto "main/internal/microservices/session/proto"
	user_proto "main/internal/microservices/user/proto"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
	session_mock "main/test/mocks/session"
	user_mock "main/test/mocks/user"
	"testing"

	"github.com/golang/mock/gomock"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_UserServiceServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := user_mock.NewMockRepository(ctrl)
	mockAuthRepo := session_mock.NewMockRepository(ctrl)

	userManager := &UserManager{
		UserRepo: mockUserRepo,
		AuthRepo: mockAuthRepo,
		Logger:   logrus.New(),
	}

	t.Run("Register", func(t *testing.T) {
		in := &user_proto.UserData{
			Id:       "1",
			Email:    "ivan@company.com",
			Username: "ivan",
		}
		deserialized := user_domain.User{
			Id:       in.GetId(),
			Email:    in.GetEmail(),
			Username: in.GetUsername(),
		}

		mockUserRepo.EXPECT().Create(deserialized).Return(nil)

		result, err := userManager.Register(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	const (
		userId    = "userId"
		sessionId = "sessId"
		avatarUrl = "/path/to/avatar.png"
	)

	t.Run("LogIn", func(t *testing.T) {
		in := &user_proto.UserCredentials{
			Email:    "ivan@company.com",
			Password: "qwerty",
		}
		mockUserRepo.EXPECT().CheckEmailAndPassword(in.GetEmail(), in.GetPassword()).Return(userId, nil)
		mockAuthRepo.EXPECT().Create(userId).Return(sessionId, nil)

		result, err := userManager.LogIn(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &session_proto.SessionId{SessionId: sessionId}, result)
	})

	inSessionId := &session_proto.SessionId{SessionId: sessionId}

	t.Run("Auth Success", func(t *testing.T) {
		mockAuthRepo.EXPECT().Get(inSessionId.GetSessionId()).Return(userId, nil)
		result, err := userManager.Auth(context.Background(), inSessionId)

		assert.Nil(t, err)
		assert.Equal(t, &session_proto.Status{IsOk: true}, result)
	})

	t.Run("Auth Error", func(t *testing.T) {
		mockAuthRepo.EXPECT().Get(inSessionId.GetSessionId()).Return("", errors.New("error"))
		result, err := userManager.Auth(context.Background(), inSessionId)

		assert.Equal(t, &session_proto.Status{IsOk: false}, result)
		assert.Equal(t, session.ErrSessionDoesNotExist, err)
	})

	t.Run("LogOut", func(t *testing.T) {
		mockAuthRepo.EXPECT().Delete(inSessionId.GetSessionId()).Return(nil)

		result, err := userManager.LogOut(context.Background(), inSessionId)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	user := user_domain.User{
		Id:        "qwer-qwer-qwer",
		Email:     "user@example.com",
		Password:  "my_password",
		BirthDate: "2003-01-12",
		Avatar:    "Avatar url",
	}

	serializedUser := &user_proto.UserData{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		BirthDate: user.BirthDate,
		Avatar:    user.Avatar,
	}

	t.Run("GetUserInfo", func(t *testing.T) {
		mockAuthRepo.EXPECT().Get(inSessionId.GetSessionId()).Return(userId, nil)
		mockUserRepo.EXPECT().GetById(userId).Return(user, nil)

		result, err := userManager.GetUserInfo(context.Background(), inSessionId)
		assert.Nil(t, err)
		assert.Equal(t, serializedUser, result)
	})

	t.Run("UpdateUserInfo", func(t *testing.T) {
		mockUserRepo.EXPECT().UpdateUserInfo(user).Return(nil)

		result, err := userManager.UpdateUserInfo(context.Background(), serializedUser)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("UploadAvatar", func(t *testing.T) {
		in := &user_proto.ImageToUser{
			Url: &image_proto.ImageUrl{Url: avatarUrl},
			Id:  &session_proto.UserId{UserId: userId},
		}

		mockUserRepo.EXPECT().UpdateAvatarPath(in.GetId().GetUserId(), in.GetUrl().GetUrl()).Return(nil)
		result, err := userManager.UploadAvatar(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})

	t.Run("RemoveAvatar", func(t *testing.T) {
		in := &session_proto.UserId{UserId: userId}

		mockUserRepo.EXPECT().RemoveAvatarPath(in.GetUserId()).Return(avatarUrl, nil)

		result, err := userManager.RemoveAvatar(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &image_proto.ImageUrl{Url: avatarUrl}, result)
	})

	t.Run("CheckEmailExist", func(t *testing.T) {
		in := &user_proto.Email{Email: "user@mail.ru"}

		mockUserRepo.EXPECT().CheckEmailExist(in.GetEmail()).Times(1).Return(nil)

		_, err := userManager.CheckEmail(context.Background(), in)
		assert.Nil(t, err)
	})

	t.Run("UpdatePassword", func(t *testing.T) {
		in := &user_proto.UserCredentials{Email: "user@mail.ru", Password: "password"}

		mockUserRepo.EXPECT().UpdatePassword(in.GetEmail(), in.GetPassword()).Times(1).Return(nil)

		_, err := userManager.UpdatePassword(context.Background(), in)
		assert.Nil(t, err)
	})

}
