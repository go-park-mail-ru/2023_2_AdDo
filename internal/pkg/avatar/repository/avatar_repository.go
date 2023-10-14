package avatar_repository

import (
	"context"
	"fmt"
	avatar_domain "main/internal/pkg/avatar"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
)

const (
	BucketName  = "user-avatar"
	AvatarExtension = ".png"
	ContentType = "image/png"
)

type Minio struct {
	client *minio.Client
}

func NewMinio(client *minio.Client) Minio {
	return Minio{client: client}
}

func generateAvatarName(id uint64) string {
	return strings.Join([]string{
		"avatar", 
		strconv.FormatUint(id, 10), time.Now().Format("20060102150405"),
	}, "-") + AvatarExtension
}

func (mn Minio) Create(avatar avatar_domain.Avatar) (string, error) {
	objName :=  generateAvatarName(avatar.UserId)
	fmt.Println("Before put avatar record")
	_, err := mn.client.PutObject(
		context.Background(),
		BucketName,
		objName,
		avatar.Payload,
		avatar.PayloadSize,
		minio.PutObjectOptions{ContentType: ContentType},
	)
	if err != nil {
		return "", err
	}
	
	return filepath.Join("/", BucketName, objName), nil
}

func (mn Minio) Remove(path string) error {
	return mn.client.RemoveObject(
		context.Background(),
		BucketName,
		filepath.Base(path),
		minio.RemoveObjectOptions{},
	)
}
