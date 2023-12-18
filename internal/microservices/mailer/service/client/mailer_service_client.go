package client

import (
	"context"
	"main/internal/microservices/mailer/proto"

	"github.com/sirupsen/logrus"
)

type Client struct {
	mailer proto.MailerServiceClient
	logger *logrus.Logger
}

func NewClient(mailer proto.MailerServiceClient, logger *logrus.Logger) Client {
	return Client{mailer: mailer, logger: logger}
}

func (c Client) SendToken(email string) error {
	c.logger.Infoln("Grpc client to MailerService: SendToken method")

	if _, err := c.mailer.SendToken(context.Background(), &proto.Payload{Payload: email}); err != nil {
		return err
	}
	return nil
}

func (c Client) GetEmail(resetToken string) (string, error) {
	c.logger.Infoln("Grpc client to MailerService: GetEmail method")

	payload, err := c.mailer.CheckToken(context.Background(), &proto.Payload{Payload: resetToken})
	if err != nil {
		return "", err
	}
	return payload.GetPayload(), nil
}
