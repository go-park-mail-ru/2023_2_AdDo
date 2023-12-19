package user_domain

import (
	"errors"
	"io"

	"github.com/asaskevich/govalidator"
	xssvalidator "github.com/infiniteloopcloud/xss-validator"
)

type UserCredentials struct {
	Email    string `valid:"length(1|30), email, required, printableascii" json:"Email" example:"example@gmail.com"`
	Password string `valid:"length(6|30), required, printableascii" json:"Password" example:"password"`
}

type User struct {
	Id        string `valid:"-" json:"Id" example:"qwer-werw-we4w"`
	Username  string `valid:"length(2|30), required, printableascii" json:"Username" example:"john"`
	Email     string `valid:"length(1|30), email, required, printableascii" json:"Email" example:"example@gmail.com"`
	Password  string `valid:"length(6|30), required, printableascii" json:"Password" example:"password"`
	BirthDate string `valid:"required" json:"BirthDate" example:"2000-01-01"`
	Avatar    string `valid:"url_optional" json:"Avatar" example:"http://test/images/1.jpg,http://test/images/2.jpg"`
}

type ForgotPasswordInput struct {
	Email string `valid:"length(1|30), email, required, printableascii" json:"Email" example:"example@gmail.com"`
}

type ResetPasswordInput struct {
	Password        string `valid:"length(6|30), required, printableascii" json:"Password" example:"password"`
	ConfirmPassword string `valid:"length(6|30), required, printableascii" json:"ConfirmPassword" example:"password"`
}

func (u *User) ValidateForUpdate() error {
	u.Password = "password"
	return u.Validate()
}

func (u *User) Validate() error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	err = xssvalidator.Validate(u.Email, xssvalidator.DefaultRules...)
	if err != nil {
		return err
	}

	err = xssvalidator.Validate(u.Password, xssvalidator.DefaultRules...)
	if err != nil {
		return err
	}

	err = xssvalidator.Validate(u.Username, xssvalidator.DefaultRules...)
	if err != nil {
		return err
	}

	err = xssvalidator.Validate(u.BirthDate, xssvalidator.DefaultRules...)
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

	err = xssvalidator.Validate(uC.Email, xssvalidator.DefaultRules...)
	if err != nil {
		return err
	}

	err = xssvalidator.Validate(uC.Password, xssvalidator.DefaultRules...)
	if err != nil {
		return err
	}

	return nil
}

func (fpi ForgotPasswordInput) Validate() error {
	_, err := govalidator.ValidateStruct(fpi)
	if err != nil {
		return err
	}

	return nil
}

func (rpi ResetPasswordInput) Validate() error {
	_, err := govalidator.ValidateStruct(rpi)
	if err != nil {
		return err
	}

	if rpi.Password != rpi.ConfirmPassword {
		return ErrPasswordsDoNotMatch
	}

	return nil
}

type UploadAvatarResponse struct {
	Url string `json:"AvatarUrl" example:"/user-images/images.png"`
}

type UseCase interface {
	Register(user User) error
	Login(email, password string) (string, error)
	Auth(sessionId string) (bool, error)
	GetUserInfo(sessionId string) (User, error)
	Logout(sessionId string) error
	UpdateUserInfo(userId string, user User) error
	UploadAvatar(userId string, src io.Reader, size int64) (string, error)
	RemoveAvatar(userId string) error
	GetUserName(userId string) (string, error)
	ForgotPassword(email string) error
	UpdatePassword(resetToken, password string) error
}

type Repository interface {
	Create(user User) error
	GetById(id string) (User, error)
	CheckEmailAndPassword(email string, password string) (string, error)
	UpdateUserInfo(user User) error
	UpdateAvatarPath(userId string, path string) error
	GetAvatarPath(userId string) (string, error)
	RemoveAvatarPath(userId string) (string, error)
	GetUserNameById(userId string) (string, error)
	CheckEmail(email string) error
	UpdatePassword(email, password string) error
}

var (
	ErrWrongCredentials    = errors.New("wrong user credentials")
	ErrUserAlreadyExist    = errors.New("user already exist")
	ErrUserDoesNotExist    = errors.New("user does not exist")
	ErrPasswordsDoNotMatch = errors.New("passwords do not match")
)
