package avatar_repository

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	avatar_domain "main/internal/pkg/avatar"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
)

const BucketName = "user-avatar"

type Minio struct {
	client *minio.Client
}

func NewMinio(client *minio.Client) Minio {
	return Minio{client: client}
}

func generateAvatarName(id uint64, extansion string) string {
	data := []byte(strconv.FormatUint(id, 10) + time.Now().Format("20060102150405"))
	hash := md5.Sum(data)
	hashString := hex.EncodeToString(hash[:])
	return strings.Join([]string{"avatar", hashString}, "-") + "." + extansion
}

func (mn Minio) Create(avatar avatar_domain.Avatar) (string, error) {
	objName := generateAvatarName(avatar.UserId, path.Base(avatar.ContentType))
	_, err := mn.client.PutObject(
		context.Background(),
		BucketName,
		objName,
		avatar.Payload,
		avatar.PayloadSize,
		minio.PutObjectOptions{ContentType: avatar.ContentType},
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
