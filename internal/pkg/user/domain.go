package user_domain

import (
	"errors"
)

type User struct {
	Id        uint64
	Username  string
	Email     string
	Password  string
	BirthDate string
	Avatar    string
}

type ResponseId struct {
	Id uint64
}

type Usecase interface {
	Register(user User) (uint64, error)
	Login(email, password string) (uint64, string, error)
	Auth(userId uint64, sessionId string) (bool, error)
	//GetById(id uint64) (*User, error)
	Logout(id uint64) error
}

type Repository interface {
	Create(user User) (uint64, error)
	GetById(id uint64) (User, error)
	CheckEmailAndPassword(email string, password string) (uint64, error)
}

var (
	ErrUserAlreadyExist      = errors.New("user already exist")
	ErrUserDoesNotExist      = errors.New("user does not exist")
	ErrSessionDoesNotExist   = errors.New("session does not exist")
	ErrSessionCreatingFailed = errors.New("session does not exist")
)
