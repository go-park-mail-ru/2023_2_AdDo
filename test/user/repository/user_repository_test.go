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
	data := user_domain.User{Email: "John@email.com", Password: "John's password", Username: "John's username", BirthDate: "2003-12-01"}

	query := "insert into profile"
	mock.ExpectExec(query).WithArgs(data.Email, sqlmock.AnyArg(), data.Username, data.BirthDate).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Create(data)
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

	expectedUser := user_domain.User{
		Id:        1,
		Email:     "email@mail.com",
		Username:  "TestUser",
		BirthDate: "2000-01-01",
		Avatar:    "https://example.com/avatar.jpg",
	}

	profileTable := sqlmock.NewRows([]string{"email", "nickname", "birth_date", "avatar_url"}).
		AddRow(expectedUser.Email, expectedUser.Username, expectedUser.BirthDate, expectedUser.Avatar)

	mock.ExpectQuery("select email, nickname, birth_date, avatar_url from profile where id = ?").
		WithArgs(expectedUser.Id).WillReturnRows(profileTable)

	user, err := repo.GetById(expectedUser.Id)
	if err != nil {
		t.Errorf("Error getting user by id: %v", err)
	}

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
	expectedUserId := uint64(1)

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	mock.ExpectQuery("select id from profile").
		WithArgs("test@example.com", sqlmock.AnyArg()).
		WillReturnRows(rows)

	userId, err := repo.CheckEmailAndPassword("test@example.com", "password")
	if err != nil {
		t.Errorf("Error checking email and password: %v", err)
	}

	if userId != expectedUserId {
		t.Errorf("Expected user id %d but got %d", expectedUserId, userId)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
