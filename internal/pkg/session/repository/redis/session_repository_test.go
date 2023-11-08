package session_repository_redis

import (
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionRepository(t *testing.T) {
	db, mock := redismock.NewClientMock()

	repo := Redis{
		database: db,
		logger:   logrus.New(),
	}

	const sessionId = "session"
	const userId = "1"

	t.Run("Get success", func(t *testing.T) {
		mock.ExpectGet(sessionId).SetVal(userId)

		receivedUserId, err := repo.Get(sessionId)
		assert.Nil(t, err)
		assert.Equal(t, userId, receivedUserId)
	})

	t.Run("Get with error", func(t *testing.T) {
		mock.ExpectGet(sessionId).SetErr(errors.New("fail"))

		receivedUserId, err := repo.Get(sessionId)
		assert.Equal(t, errors.New("fail"), err)
		assert.Empty(t, receivedUserId)
	})

	t.Run("Delete with error", func(t *testing.T) {
		mock.ExpectDel(sessionId).SetErr(errors.New("fail"))

		err := repo.Delete(sessionId)
		assert.Equal(t, errors.New("fail"), err)
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
