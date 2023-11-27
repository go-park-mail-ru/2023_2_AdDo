package grpc_image

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	image_proto "main/internal/microservices/image/proto"
)

type Client struct {
	imageManager image_proto.ImageServiceClient
	logger       *logrus.Logger
}

func NewClient(pm image_proto.ImageServiceClient, logger *logrus.Logger) Client {
	return Client{imageManager: pm, logger: logger}
}

func (c *Client) UploadAvatar(src io.Reader, size int64) (string, error) {
	c.logger.Infoln("Image client UploadAvatar entered")

	data, err := io.ReadAll(src)
	if err != nil {
		c.logger.Errorln(err.Error())
	}
	result, err := c.imageManager.UploadAvatar(context.Background(), &image_proto.Image{
		Data: data,
		Size: uint64(size),
	})
	if err != nil {
		return "", err
	}
	c.logger.Infoln("avatar uploaded", result.GetUrl())

	return result.GetUrl(), nil
}

func (c *Client) UploadPlaylistImage(src io.Reader, size int64) (string, error) {
	c.logger.Infoln("Image client UploadPlaylistImage entered")

	data, err := io.ReadAll(src)
	if err != nil {
		c.logger.Errorln(err.Error())
	}
	result, err := c.imageManager.UploadPlaylistImage(context.Background(), &image_proto.Image{
		Data: data,
		Size: uint64(size),
	})
	if err != nil {
		return "", err
	}
	c.logger.Infoln("playlist images uploaded", result.GetUrl())

	return result.GetUrl(), nil
}

func (c *Client) RemoveImage(url string) error {
	c.logger.Infoln("Image client RemoveImage entered")

	if _, err := c.imageManager.RemoveImage(context.Background(), &image_proto.ImageUrl{Url: url}); err != nil {
		return err
	}
	c.logger.Infoln("images removed")

	return nil
}
