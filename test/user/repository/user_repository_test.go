package user_repository_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	user_domain "main/internal/pkg/user"
	user_repository "main/internal/pkg/user/repository/postgresql"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := user_repository.Postgres{
		Database: db,
	}

	// Ожидаемый SQL-запрос
	mock.ExpectExec("INSERT INTO user").
		WithArgs("John", "Doe").
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(user_domain.User{Username: "John", Password: "Doe"})
	if err != nil {
		t.Errorf("Error creating user: %v", err)
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

	repo := user_repository.Postgres{
		Database: db,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "lastname"}).
		AddRow(1, "John", "Doe")

	// Ожидаемый SQL-запрос
	mock.ExpectQuery("SELECT * FROM user").
		WithArgs(1).
		WillReturnRows(rows)

	user, err := repo.GetById(1)
	if err != nil {
		t.Errorf("Error getting user by id: %v", err)
	}

	expectedUser := user_domain.User{Id: 1, Username: "John", Password: "Doe"}
	if user != expectedUser {
		t.Errorf("Expected user %+v but got %+v", expectedUser, user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestUserRepository_CheckEmailAndPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := user_repository.Postgres{
		Database: db,
	}

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	// Ожидаемый SQL-запрос
	mock.ExpectQuery("SELECT * FROM user").
		WithArgs("test@example.com", "password").
		WillReturnRows(rows)

	userId, err := repo.CheckEmailAndPassword("test@example.com", "password")
	if err != nil {
		t.Errorf("Error checking email and password: %v", err)
	}

	expectedUserId := uint64(1)
	if userId != expectedUserId {
		t.Errorf("Expected user id %d but got %d", expectedUserId, userId)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
