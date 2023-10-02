package session_repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	user_domain "main/internal/pkg/user"
	"testing"
)

func TestSessionRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()
	repo := Postgres{
		database: db,
	}
	data := user_domain.User{Id: 1}

	query := "insert into session"
	mock.ExpectExec(query).WithArgs(data.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	query = "select session_id from session"
	mock.ExpectQuery(query).WithArgs(data.Id).WillReturnRows(sqlmock.NewRows([]string{"session_id"}).AddRow("session_id"))
	sessionId, err := repo.Create(data.Id)
	if err != nil || sessionId == "" {
		t.Errorf("Error creating session: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestUserRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		database: db,
	}

	expectedUser := user_domain.User{Id: 1}
	sessionId := "sessionId"

	profileTable := sqlmock.NewRows([]string{"session_id"}).
		AddRow(sessionId)

	mock.ExpectQuery("select session_id from session").
		WithArgs(expectedUser.Id).WillReturnRows(profileTable)

	received, err := repo.GetByUserId(expectedUser.Id)
	if err != nil {
		t.Errorf("Error getting user by id: %v", err)
	}

	assert.Equal(t, sessionId, received)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestUserRepository_DeleteByUserId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := Postgres{
		database: db,
	}

	expectedUser := user_domain.User{Id: 1}

	mock.ExpectExec("delete from session").
		WithArgs(expectedUser.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteByUserId(expectedUser.Id)
	if err != nil {
		t.Errorf("Error getting user by id: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
