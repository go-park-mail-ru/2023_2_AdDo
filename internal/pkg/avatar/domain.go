package avatar_domain

import (
	"errors"
	"io"
)

type Avatar struct {
	Payload     io.Reader
	PayloadSize int64
	UserId      uint64
	ContentType string
}

type UseCase interface {
	UploadAvatar(userId uint64, src io.Reader, size int64) error
	RemoveAvatar(userId uint64) error
}

type S3Repository interface {
	Create(avatar Avatar) (string, error)
	Remove(path string) error
}

type DbRepository interface {
	UpdateAvatarPath(userId uint64, path string) error
	GetAvatarPath(userId uint64) (string, error)
	RemoveAvatarPath(userId uint64) error
}

const MaxAvatarSize = 16 * 1024 * 1024

var (
	ErrAvatarDoesNotExist = errors.New("avatar does not exist")
	ErrAvatarIsTooLarge   = errors.New("file is too large")
	ErrCannotRead         = errors.New("cannot read file")
	ErrWrongAvatarType    = errors.New("wrong avatar type")
)
