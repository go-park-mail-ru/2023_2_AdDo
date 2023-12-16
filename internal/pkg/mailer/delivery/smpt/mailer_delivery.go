package mailer_delivery

import (
	"bytes"
	"embed"
	"html/template"
	domain "main/internal/pkg/mailer"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

//go:embed "templates"
var templateFS embed.FS

type Mailer struct {
	sender string
	dialer *gomail.Dialer
	logger *logrus.Logger
}

func New(smpt domain.Smtp, logger *logrus.Logger) Mailer {
	dialer := gomail.NewDialer(smpt.Host, smpt.Port, smpt.Username, smpt.Password)

	return Mailer{
		dialer: dialer,
		sender: smpt.Sender,
		logger: logger,
	}
}

func (m Mailer) Send(recipient, templateFile string, emailData domain.EmailData) error {
	// TODO: include style.html to teplateFile
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, templateFile, emailData)
	if err != nil {
		return err
	}
	//
	// 	plainBody := new(bytes.Buffer)
	// 	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	// 	if err != nil {
	// 		return err
	// 	}
	//
	// 	htmlBody := new(bytes.Buffer)
	// 	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	// 	if err != nil {
	// 		return err
	// 	}

	msg := gomail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetBody("text/html", body.String())
	// msg.SetHeader("Subject", subject.String())
	// msg.SetBody("text/plain", plainBody.String())

	if err := m.dialer.DialAndSend(msg); err != nil {
		m.logger.Errorf("Error send message: %s", err.Error())
		return err
	}
	m.logger.Infoln("Message successfull sent")

	return nil
}
