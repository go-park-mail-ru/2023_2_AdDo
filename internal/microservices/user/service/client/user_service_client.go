package grpc_user

import (
	"context"
	pb "main/internal/microservices/user/proto"
	grpc_server_user "main/internal/microservices/user/service/server"
	user_domain "main/internal/pkg/user"
)

type Client struct {
	userManager grpc_server_user.UserManager
}

func NewClient(um grpc_server_user.UserManager) Client {
	return Client{
		userManager: um,
	}
}

func (c *Client) Register(user user_domain.User) error {
	c.userManager.Logger.Infoln("Grpc client to UserService: Register Method")

	_, err := c.userManager.Register(context.Background(), grpc_server_user.SerializeUserData(user))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Login(email, password string) (string, error) {
	c.userManager.Logger.Infoln("Grpc client to UserService: Login Method")

	pbSessionId, err := c.userManager.LogIn(context.Background(), &pb.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	return pbSessionId.GetSessionId(), nil
}

func (c *Client) Auth(sessionId string) (bool, error) {
	c.userManager.Logger.Infoln("Grpc client to UserService: Auth Method")

	isAuth, err := c.userManager.Auth(context.Background(), &pb.SessionId{SessionId: sessionId})
	if err != nil || !isAuth.GetIsOk() {
		return false, err
	}

	return true, nil
}

func (c *Client) GetUserInfo(sessionId string) (user_domain.User, error) {
	c.userManager.Logger.Infoln("Grpc client to UserService: GetUserInfo Method")

	u, err := c.userManager.GetUserInfo(context.Background(), &pb.SessionId{SessionId: sessionId})
	if err != nil {
		return user_domain.User{}, err
	}

	return grpc_server_user.DeserializeUserData(u), nil
}

func (c *Client) Logout(sessionId string) error {
	c.userManager.Logger.Infoln("Grpc client to UserService: Logout Method")

	_, err := c.userManager.LogOut(context.Background(), &pb.SessionId{SessionId: sessionId})
	if err != nil {
		return err
	}

	return nil
}
