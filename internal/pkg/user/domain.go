package user_domain

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

type User struct {
	Id        uint64 `valid:"-"`
	Username  string `valid:"length(2|30), alphanum, required"`
	Email     string `valid:"length(1|30), email, required"`
	Password  string `valid:"length(6|30), required"`
	BirthDate string `valid:"required"`
	Avatar    string `valid:"url_optional"`
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	return nil
}

type ResponseId struct {
	Id uint64
}

type UseCase interface {
	Register(user User) error
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
