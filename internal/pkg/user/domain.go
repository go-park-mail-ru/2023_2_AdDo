package user_domain

import (
	"errors"
)

type User struct {
	Id        uint64 `json:"id" example:"1"`
	Username  string `json:"username" example:"john"`
	Email     string `json:"email" example:"example@gmail.com"`
	Password  string `json:"password" example:"password"`
	BirthDate string `json:"date" example:"2000-01-01"`
	Avatar    string `json:"avatar" example:"http://test/image/1.jpg,http://test/image/2.jpg"`
}

type ResponseId struct {
	Id uint64 `json:"id" example:"1"`
}

type UseCase interface {
	Register(user User) (uint64, error)
	Login(email, password string) (uint64, string, error)
	Auth(userId uint64, sessionId string) (bool, error)
	GetUserInfo(id uint64) (User, error)
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
