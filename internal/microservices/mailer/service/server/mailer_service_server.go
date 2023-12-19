package server

import (
	"bytes"
	"context"
	"html/template"
	"main/internal/microservices/mailer/proto"
	domain "main/internal/pkg/mailer"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

const (
	ResetPasswordSubject = "Восстановление пароля"
	SiteUrl              = "https://musicon.space/auth/reset_password/"
)

func NewMailerServer(smpt domain.Smtp, redisRepo domain.Repository, logger *logrus.Logger) MailerServer {
	dialer := gomail.NewDialer(smpt.Host, smpt.Port, smpt.Username, smpt.Password)

	templates := template.Must(template.ParseGlob("/templates/*"))

	return MailerServer{
		dialer:    dialer,
		sender:    smpt.Sender,
		redisRepo: redisRepo,
		logger:    logger,
		tmpl:      templates,
	}
}

type MailerServer struct {
	sender    string
	dialer    *gomail.Dialer
	redisRepo domain.Repository
	tmpl      *template.Template
	logger    *logrus.Logger
	proto.UnimplementedMailerServiceServer
}

func (ms MailerServer) SendToken(ctx context.Context, payload *proto.Payload) (*empty.Empty, error) {
	ms.logger.Infoln("Mailer Mircos SendToken entered")

	recipient := payload.GetPayload()

	resetToken, err := ms.redisRepo.CreateToken(recipient)
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)

	ms.tmpl.ExecuteTemplate(body, domain.ResetPasswordHtmlFile, domain.EmailData{
		URL:     SiteUrl + resetToken,
		Subject: ResetPasswordSubject,
	})

	msg := gomail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", ms.sender)
	msg.SetBody("text/html", body.String())
	msg.SetHeader("Subject", ResetPasswordSubject)

	if err := ms.dialer.DialAndSend(msg); err != nil {
		ms.logger.Errorf("Error send message to %s: %s", recipient, err.Error())
		return nil, err
	}
	ms.logger.Infoln("Message successfull sent")

	return &empty.Empty{}, nil
}

func (ms MailerServer) CheckToken(ctx context.Context, resetToken *proto.Payload) (*proto.Payload, error) {
	ms.logger.Infoln("Mailer Mircos CheckToken entered")

	rt := resetToken.GetPayload()

	email, err := ms.redisRepo.CheckToken(rt)
	if err != nil {
		return nil, err
	}
	ms.logger.Infoln("Reset token was successfully checked")

	if err := ms.redisRepo.Delete(rt); err != nil {
		return nil, err
	}
	ms.logger.Info("Reset token was successfully deleted")

	return &proto.Payload{Payload: email}, nil
}
