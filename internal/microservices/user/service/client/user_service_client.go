package grpc_user

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	grpc_image "main/internal/microservices/image/service/client"
	session_proto "main/internal/microservices/session/proto"
	user_proto "main/internal/microservices/user/proto"
	grpc_server_user "main/internal/microservices/user/service/server"
	user_domain "main/internal/pkg/user"
)

type Client struct {
	userClient  user_proto.UserServiceClient
	imageClient grpc_image.Client
	logger      *logrus.Logger
}

func NewClient(client user_proto.UserServiceClient, imageClient grpc_image.Client, logger *logrus.Logger) Client {
	return Client{
		userClient:  client,
		logger:      logger,
		imageClient: imageClient,
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

func (c *Client) UploadAvatar(userId string, src io.Reader, size int64) (string, error) {
	c.logger.Infoln("User grpc client UploadAvatar entered")

	avatarUrl, err := c.imageClient.UploadAvatar(src, size)
	if err != nil {
		return "", err
	}
	c.logger.Infoln("Image uploaded")

	_, err = c.userClient.UploadAvatar(context.Background(), &user_proto.ImageToUser{
		Url: &user_proto.ImageUrl{Url: avatarUrl},
		Id:  &session_proto.UserId{UserId: userId},
	})
	if err != nil {
		return "", err
	}
	c.logger.Infoln("Avatar path add to db")

	return avatarUrl, nil
}

func (c *Client) RemoveAvatar(userId string) error {
	c.logger.Infoln("User grpc client RemoveAvatar entered")

	avatarUrl, err := c.userClient.RemoveAvatar(context.Background(), &session_proto.UserId{UserId: userId})
	if err != nil {
		return err
	}
	c.logger.Infoln("avatar removed", avatarUrl.GetUrl())

	err = c.imageClient.RemoveImage(avatarUrl.GetUrl())
	if err != nil {
		return err
	}
	c.logger.Infoln("image removed")

	if err != nil {
		return err
	}
	c.logger.Infoln("Avatar Removed")

	return nil
}

func (c *Client) UpdateUserInfo(userId string, u user_domain.User) error {
	c.logger.Infoln("User grpc client UpdateUserInfo entered")

	u.Id = userId
	_, err := c.userClient.UpdateUserInfo(context.Background(), grpc_server_user.SerializeUserData(u))
	if err != nil {
		return err
	}
	c.logger.Infoln("Info updated successfully")

	return nil
}
