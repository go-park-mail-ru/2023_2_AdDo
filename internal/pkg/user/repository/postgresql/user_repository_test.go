package user_repository

import (
	"github.com/pashagolub/pgxmock/v3"
	"github.com/sirupsen/logrus"
	user_domain "main/internal/pkg/user"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	mock, err := pgxmock.NewPool()
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	data := user_domain.User{Email: "John@email.com", Password: "John's password", Username: "John's username", BirthDate: "2003-12-01"}

	query := "insert into profile"
	mock.ExpectExec(query).WithArgs(data.Email, pgxmock.AnyArg(), data.Username, data.BirthDate).WillReturnResult(pgxmock.NewResult("insert", 1))

	err = repo.Create(data)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestUserRepository_GetById(t *testing.T) {
	mock, err := pgxmock.NewPool()
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	expectedUser := user_domain.User{
		Email:     "email@mail.com",
		Username:  "TestUser",
		BirthDate: "2000-01-01",
		Avatar:    "https://example.com/avatar.jpg",
	}

	profileTable := pgxmock.NewRows([]string{"email", "nickname", "birth_date", "avatar_url"}).
		AddRow(expectedUser.Email, expectedUser.Username, expectedUser.BirthDate, expectedUser.Avatar)

	mock.ExpectQuery("select email, nickname, birth_date, avatar_url from profile where id = ?").
		WithArgs(pgxmock.AnyArg()).WillReturnRows(profileTable)

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
	mock, err := pgxmock.NewPool()
	defer mock.Close()

	repo := Postgres{
		Pool:   mock,
		logger: logrus.New(),
	}

	expectedUserId := "rand-strs-uuid"

	rows := pgxmock.NewRows([]string{"id"}).
		AddRow(expectedUserId)

	mock.ExpectQuery("select id from profile").
		WithArgs("test@example.com", pgxmock.AnyArg()).
		WillReturnRows(rows)

	_, err = repo.CheckEmailAndPassword("test@example.com", "password")
	if err != nil {
		t.Errorf("Error checking email and password: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
