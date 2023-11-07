package image_repository

import (
	"context"
	"fmt"
	minio_init "main/init/minio"
	avatar_domain "main/internal/pkg/image"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

type Minio struct {
	client *minio.Client
}

func NewMinio(client *minio.Client) Minio {
	return Minio{client: client}
}

func (mn Minio) UploadAvatar(avatar avatar_domain.Base) (string, error) {
	return mn.create(avatar, minio_init.UserAvatarBucketName)
}

func (mn Minio) UploadPlaylistImage(avatar avatar_domain.Base) (string, error) {
	return mn.create(avatar, minio_init.PlaylistImageBucketName)
}

func (mn Minio) create(avatar avatar_domain.Base, bucketName string) (string, error) {
	if _, err := mn.client.PutObject(
		context.Background(),
		minio_init.UserAvatarBucketName,
		avatar.Name,
		avatar.Payload,
		avatar.PayloadSize,
		minio.PutObjectOptions{ContentType: avatar.ContentType},
	); err != nil {
		return "", err
	}

	return filepath.Join("/", bucketName, avatar.Name), nil
}

func (mn Minio) Remove(path string) error {
	fmt.Println(filepath.Base(filepath.Dir(path)))
	return mn.client.RemoveObject(
		context.Background(),
		filepath.Base(filepath.Dir(path)),
		filepath.Base(path),
		minio.RemoveObjectOptions{},
	)
}
