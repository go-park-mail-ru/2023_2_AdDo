package grpc_session_server

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	pb "main/internal/microservices/session/proto"
	session_mock "main/test/mocks/session"
	"testing"
)

func Test_SessionServiceServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessionRepo := session_mock.NewMockRepository(ctrl)

	sessionManager := &SessionManager{
		repoSession: mockSessionRepo,
		logger:      logrus.New(),
	}

	const sessionId = "sessId"
	const userId = "user"
	in := &pb.SessionId{SessionId: sessionId}

	t.Run("CheckSession Success", func(t *testing.T) {
		mockSessionRepo.EXPECT().Get(in.GetSessionId()).Return(userId, nil)

		result, err := sessionManager.CheckSession(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &pb.Status{IsOk: true}, result)
	})

	t.Run("CheckSession Error", func(t *testing.T) {
		mockSessionRepo.EXPECT().Get(in.GetSessionId()).Return("", errors.New("error"))

		result, err := sessionManager.CheckSession(context.Background(), in)
		assert.NotNil(t, err)
		assert.Equal(t, &pb.Status{IsOk: false}, result)
	})

	t.Run("GetUserId", func(t *testing.T) {
		mockSessionRepo.EXPECT().Get(in.GetSessionId()).Return(userId, nil)

		result, err := sessionManager.GetUserId(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &pb.UserId{UserId: userId}, result)
	})
}
