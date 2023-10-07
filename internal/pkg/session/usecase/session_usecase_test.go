package session_usecase

import (
	"github.com/golang/mock/gomock"
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
	}

	const isAuthExpected = true
	const sessionId = "sessionId"
	const anyUserId = 1

	mockSessionRepo.EXPECT().Get(sessionId).Return(uint64(anyUserId), nil)
	isAuth, err := useCase.CheckSession(sessionId)

	assert.Equal(t, nil, err)
	assert.Equal(t, isAuthExpected, isAuth)
}
