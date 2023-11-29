package minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	UserAvatarBucketName    = "user-avatar"
	PlaylistImageBucketName = "playlist-preview"
)

func InitMinio() (*minio.Client, error) {
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	mn, err := minio.New("api.s3.musicon.space", &minio.Options{
		Creds: credentials.NewStaticV4(
			accessKeyID,
			secretAccessKey,
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
