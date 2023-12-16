package domain

type Smtp struct {
	Port     int
	Host     string
	Username string
	Password string
	Sender   string
}

type EmailData struct {
	URL     string
	Name    string
	Subject string
}

const (
	ResetPasswordHtml = "reset_password.html"
)
