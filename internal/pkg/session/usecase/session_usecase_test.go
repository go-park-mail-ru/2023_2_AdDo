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

	userId := 1
	const sessionId = "sessionId"

	mockSessionRepo.EXPECT().GetByUserId(uint64(userId)).Return(sessionId, nil)
	isSame, err := useCase.CheckSession(sessionId, uint64(userId))

	assert.Equal(t, nil, err)
	assert.True(t, isSame)
}
