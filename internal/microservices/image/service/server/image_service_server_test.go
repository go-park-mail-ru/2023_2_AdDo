package grpc_image_server

import (
	"context"
	"github.com/golang/mock/gomock"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	image_proto "main/internal/microservices/image/proto"
	avatar_mock "main/test/mocks/avatar"
	"testing"
)

func Test_ImageServiceServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockImageRepo := avatar_mock.NewMockRepository(ctrl)

	imageManager := ImageManager{
		repoImage: mockImageRepo,
		logger:    logrus.New(),
	}

	t.Run("RemoveImage", func(t *testing.T) {
		const imageUrl = "/path/to/image.png"
		in := &image_proto.ImageUrl{Url: imageUrl}

		mockImageRepo.EXPECT().Remove(in.GetUrl()).Return(nil)
		result, err := imageManager.RemoveImage(context.Background(), in)
		assert.Nil(t, err)
		assert.Equal(t, &google_proto.Empty{}, result)
	})
}
