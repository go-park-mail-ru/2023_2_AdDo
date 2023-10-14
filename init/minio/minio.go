package minio

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio() (*minio.Client, error) {
	accessKeyID := "musicon"
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

	// mn.Ping()
	
	fmt.Println("Minio successful connected!")

	return mn, nil
}
