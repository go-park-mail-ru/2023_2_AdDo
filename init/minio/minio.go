package minio

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	UserAvatarBucketName    = "user-image"
	PlaylistImageBucketName = "playlist-image"
)

func InitMinio() (*minio.Client, error) {
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"

	mn, err := minio.New("service-minio:9000", &minio.Options{
		Creds: credentials.NewStaticV4(
			accessKeyID,
			secretAccessKey,
			"",
		),
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

	fmt.Println("Minio successful connected!")

	return mn, nil
}
