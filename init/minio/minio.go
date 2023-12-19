package minio

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	UserAvatarBucketName    = "user-avatar"
	PlaylistImageBucketName = "playlist-image"
	EnvMinioApiUrl          = "MINIO_API_URL"
	EnvAccessKeyID          = "MINIO_ACCESS_KEY"
	EnvSecretAccessKey      = "MINIO_SECRET_KEY"
	EnvUseSSL               = "MINIO_USE_SSL"
)

func InitMinio() (*minio.Client, error) {
	useSSL, err := strconv.ParseBool(os.Getenv(EnvUseSSL))
	if err != nil {
		return nil, err
	}

	mn, err := minio.New(os.Getenv(EnvMinioApiUrl), &minio.Options{
		Creds: credentials.NewStaticV4(
			os.Getenv(EnvAccessKeyID),
			os.Getenv(EnvSecretAccessKey),
			"",
		),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	requiredBuckets := []string{UserAvatarBucketName, PlaylistImageBucketName}
	for _, bucket := range requiredBuckets {
		isExist, err := mn.BucketExists(context.Background(), bucket)
		if err != nil {
			return nil, err
		}

		if !isExist {
			return nil, fmt.Errorf("bucket %s does not exist", bucket)
		}
	}

	return mn, nil
}
