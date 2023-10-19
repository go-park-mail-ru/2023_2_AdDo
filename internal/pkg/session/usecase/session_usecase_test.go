package session_usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	session_mock "main/test/mocks/session"
	"testing"
)

func TestDefault_CheckSession_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessionRepo := session_mock.NewMockRepository(ctrl)

	useCase := &Default{
		repoSession: mockSessionRepo,
		logger:      logrus.New(),
	}

	const isAuthExpected = true
	const sessionId = "sessionId"
	const anyUserId = "qwer-qwer-qwer"

	mockSessionRepo.EXPECT().Get(sessionId).Return(anyUserId, nil)
	isAuth, err := useCase.CheckSession(sessionId)

	assert.Equal(t, nil, err)
	assert.Equal(t, isAuthExpected, isAuth)

	mockSessionRepo.EXPECT().Get(sessionId).Return(anyUserId, nil)
	userId, err := useCase.GetUserId(sessionId)

	assert.Equal(t, nil, err)
	assert.Equal(t, anyUserId, userId)
}
