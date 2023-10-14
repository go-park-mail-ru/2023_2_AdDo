package avatar_domain

import (
	"mime/multipart"
)

type Avatar struct {
	Payload     multipart.File
	PayloadSize int64
	UserId      uint64
}

type UseCase interface {
	Upload(src multipart.File, size int64) error
	Remove(src string) error
}

type S3Repository interface {
	Create(avatar Avatar) (string, error)
	Remove(path string) error
}

type DbRepository interface {
	UpdateAvatarPath(userId uint64, path string) error
	GetAvatarPath(userId uint64) error
	RemoveAvatarPath(userId uint64) error
}
