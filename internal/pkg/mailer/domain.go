package domain

import "time"

type Smtp struct {
	Port     int
	Host     string
	Username string
	Password string
	Sender   string
}

type EmailData struct {
	URL     string
	Subject string
}

const (
	ResetPasswordHtmlFile = "reset_password.html"
	ResetTokenTimeToLive  = 10 * time.Minute
)

type UseCase interface {
	SendToken(email string) error
	GetEmail(resetToken string) (string, error)
}

type Repository interface {
	CreateToken(email string) (string, error)
	CheckToken(resetToken string) (string, error)
	Delete(resetToken string) error
}
