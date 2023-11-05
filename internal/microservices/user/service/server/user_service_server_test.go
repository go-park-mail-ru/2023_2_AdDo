package grpc_server_user

import (
	"context"
	"github.com/golang/mock/gomock"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	session_proto "main/internal/microservices/session/proto"
	user_proto "main/internal/microservices/user/proto"
	user_domain "main/internal/pkg/user"
	session_mock "main/test/mocks/session"
	user_mock "main/test/mocks/user"
	"testing"
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

	t.Run("LogIn", func(t *testing.T) {
		in := &user_proto.UserCredentials{
			Email:    "ivan@company.com",
			Password: "qwerty",
		}
		const userId = "userId"
		const sessionId = "sessId"
		mockUserRepo.EXPECT().CheckEmailAndPassword(in.GetEmail(), in.GetPassword()).Return(userId, nil)
		mockAuthRepo.EXPECT().Create(userId).Return(sessionId, nil)

		result, err := userManager.LogIn(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &session_proto.SessionId{SessionId: sessionId}, result)
	})
}
