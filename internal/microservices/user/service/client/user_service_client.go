package grpc_user

import (
	"context"
	"github.com/sirupsen/logrus"
	session_proto "main/internal/microservices/session/proto"
	user_proto "main/internal/microservices/user/proto"
	grpc_server_user "main/internal/microservices/user/service/server"
	user_domain "main/internal/pkg/user"
)

type Client struct {
	userClient user_proto.UserServiceClient
	logger     *logrus.Logger
}

func NewClient(client user_proto.UserServiceClient, logger *logrus.Logger) Client {
	return Client{
		userClient: client,
		logger:     logger,
	}
}

func (c *Client) Register(user user_domain.User) error {
	c.logger.Infoln("Grpc client to UserService: Register Method")

	_, err := c.userClient.Register(context.Background(), grpc_server_user.SerializeUserData(user))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Login(email, password string) (string, error) {
	c.logger.Infoln("Grpc client to UserService: Login Method")

	pbSessionId, err := c.userClient.LogIn(context.Background(), &user_proto.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	return pbSessionId.GetSessionId(), nil
}

func (c *Client) Auth(sessionId string) (bool, error) {
	c.logger.Infoln("Grpc client to UserService: Auth Method")

	isAuth, err := c.userClient.Auth(context.Background(), &session_proto.SessionId{SessionId: sessionId})
	if err != nil || !isAuth.GetIsOk() {
		return false, err
	}

	return true, nil
}

func (c *Client) GetUserInfo(sessionId string) (user_domain.User, error) {
	c.logger.Infoln("Grpc client to UserService: GetUserInfo Method")

	u, err := c.userClient.GetUserInfo(context.Background(), &session_proto.SessionId{SessionId: sessionId})
	if err != nil {
		return user_domain.User{}, err
	}

	return grpc_server_user.DeserializeUserData(u), nil
}

func (c *Client) Logout(sessionId string) error {
	c.logger.Infoln("Grpc client to UserService: Logout Method")

	_, err := c.userClient.LogOut(context.Background(), &session_proto.SessionId{SessionId: sessionId})
	if err != nil {
		return err
	}

	return nil
}
