package grpc_image_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	image_proto "main/internal/microservices/image/proto"
	"main/internal/pkg/image"
)

type ImageManager struct {
	repoImage image.Repository
	logger    *logrus.Logger
	image_proto.UnimplementedImageServiceServer
}

func NewImageManager(rp image.Repository, logger *logrus.Logger) ImageManager {
	return ImageManager{
		repoImage: rp,
		logger:    logger,
	}
}

func (im *ImageManager) UploadAvatar(ctx context.Context, in *image_proto.Image) (*image_proto.ImageUrl, error) {
	im.logger.Infoln("Image micros UploadAvatar entered")

	base, err := image.CreateImageFromSource(in.GetData(), int64(in.GetSize()))
	if err != nil {
		im.logger.Errorln(err)
		return nil, err
	}
	im.logger.Infoln("Image got from message")

	url, err := im.repoImage.UploadAvatar(base)
	if err != nil {
		im.logger.Errorln(err)
		return nil, err
	}
	im.logger.Infoln("Image uploaded")

	return &image_proto.ImageUrl{Url: url}, nil
}

func (im *ImageManager) UploadPlaylistImage(ctx context.Context, in *image_proto.Image) (*image_proto.ImageUrl, error) {
	im.logger.Infoln("Image micros UploadPlaylistImage entered")

	base, err := image.CreateImageFromSource(in.GetData(), int64(in.GetSize()))
	if err != nil {
		return nil, err
	}
	im.logger.Infoln("Image got from message")

	url, err := im.repoImage.UploadPlaylistImage(base)
	if err != nil {
		return nil, err
	}
	im.logger.Infoln("Image uploaded")

	return &image_proto.ImageUrl{Url: url}, nil
}

func (im *ImageManager) RemoveImage(ctx context.Context, url *image_proto.ImageUrl) (*google_proto.Empty, error) {
	im.logger.Infoln("Image micros RemoveImage entered")

	err := im.repoImage.Remove(url.GetUrl())
	if err != nil {
		return nil, err
	}
	im.logger.Infoln("Image deleted")

	return &google_proto.Empty{}, nil
}
