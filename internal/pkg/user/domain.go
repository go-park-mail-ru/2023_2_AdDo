package user_domain

import (
	"errors"
	"io"

	"github.com/asaskevich/govalidator"
)

type UserCredentials struct {
	Email    string `valid:"length(1|30), email, required, printableascii" json:"Email" example:"example@gmail.com"`
	Password string `valid:"length(6|30), required, printableascii" json:"Password" example:"password"`
}

type User struct {
	Id        uint64 `valid:"-" json:"Id" example:"1"`
	Username  string `valid:"length(2|30), required, printableascii" json:"Username" example:"john"`
	Email     string `valid:"length(1|30), email, required, printableascii" json:"Email" example:"example@gmail.com"`
	Password  string `valid:"length(6|30), required, printableascii" json:"Password" example:"password"`
	BirthDate string `valid:"required" json:"BirthDate" example:"2000-01-01"`
	Avatar    string `valid:"url_optional" json:"Avatar" example:"http://test/image/1.jpg,http://test/image/2.jpg"`
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	return nil
}

func (uC *UserCredentials) Validate() error {
	_, err := govalidator.ValidateStruct(uC)
	if err != nil {
		return err
	}
	return nil
}

type ResponseId struct {
	Id uint64 `json:"Id" example:"1"`
}

type AvatarUseCase interface {
	UploadAvatar(userId uint64, src io.Reader, size int64) error
	RemoveAvatar(userId uint64) error
}

type AvatarDbRepository interface {
	UpdateAvatarPath(userId uint64, path string) error
	GetAvatarPath(userId uint64) (string, error)
	RemoveAvatarPath(userId uint64) error
}

type UseCase interface {
	Register(user User) error
	Login(email, password string) (string, error)
	Auth(sessionId string) (bool, error)
	GetUserInfo(sessionId string) (User, error)
	Logout(sessionId string) error
	AvatarUseCase
}

type Repository interface {
	Create(user User) error
	GetById(id uint64) (User, error)
	CheckEmailAndPassword(email string, password string) (uint64, error)
	AvatarDbRepository
}

var (
	ErrWrongCredentials = errors.New("wrong user credentials")
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrUserDoesNotExist = errors.New("user does not exist")
)
