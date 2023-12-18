package server

import (
	"context"
	"main/internal/microservices/mailer/proto"
	domain "main/internal/pkg/mailer"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func NewMailerServer(smpt domain.Smtp, redisRepo domain.Repository, logger *logrus.Logger) MailerServer {
	dialer := gomail.NewDialer(smpt.Host, smpt.Port, smpt.Username, smpt.Password)

	return MailerServer{
		dialer:    dialer,
		sender:    smpt.Sender,
		redisRepo: redisRepo,
		logger:    logger,
	}
}

type MailerServer struct {
	sender    string
	dialer    *gomail.Dialer
	redisRepo domain.Repository
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

	msg := gomail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", ms.sender)
	msg.SetBody("text/html", "Reset token is "+resetToken)
	msg.SetHeader("Subject", "Subject")
	// msg.SetBody("text/plain", plainBody.String())

	if err := ms.dialer.DialAndSend(msg); err != nil {
		ms.logger.Errorf("Error send message: %s", err.Error())
		return nil, err
	}
	ms.logger.Infoln("Message successfull sent")

	return &empty.Empty{}, nil
}

// TODO: include style.html to teplateFile
// 	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
// 	if err != nil {
// 		return err
// 	}
//
// 	body := new(bytes.Buffer)
// 	err = tmpl.ExecuteTemplate(body, templateFile, emailData)
// 	if err != nil {
// 		return err
// 	}
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

func (ms MailerServer) CheckToken(ctx context.Context, resetToken *proto.Payload) (*proto.Payload, error) {
	ms.logger.Infoln("Mailer Mircos GetEmail entered")

	email, err := ms.redisRepo.CheckToken(resetToken.Payload)
	if err != nil {
		return nil, err
	}

	ms.logger.Infoln("Mailer Mircos GetEmail entered")

	return &proto.Payload{Payload: email}, nil
}
