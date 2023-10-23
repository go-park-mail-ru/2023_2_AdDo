package grpc_session

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "main/internal/microservices/session/proto"
)

type Client struct {
	manager pb.SessionServiceClient
	logger  *logrus.Logger
}

func NewClient(sm pb.SessionServiceClient, logger *logrus.Logger) Client {
	return Client{
		manager: sm,
		logger:  logger,
	}
}

func (c *Client) CheckSession(sessionId string) (bool, error) {
	c.logger.Infoln("Session Client CheckSessionId entered")

	isOk, err := c.manager.CheckSession(context.Background(), &pb.SessionId{SessionId: sessionId})
	if err != nil {
		return false, err
	}
	c.logger.Infoln("grpc request no error")

	if !isOk.GetIsOk() {
		return false, nil
	}
	c.logger.Infoln("Session check ok!")

	return true, nil
}

func (c *Client) GetUserId(sessionId string) (string, error) {
	c.logger.Infoln("Session Client GetUserId entered")

	pbUserId, err := c.manager.GetUserId(context.Background(), &pb.SessionId{SessionId: sessionId})
	if err != nil {
		return "", err
	}

	return pbUserId.GetUserId(), nil
}
