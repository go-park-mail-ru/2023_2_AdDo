package avatar

import (
	"errors"
	"io"
)

type Avatar struct {
	Name        string
	Payload     io.Reader
	PayloadSize int64
	ContentType string
}

type UseCase interface {
	GetAvatar(id string, src io.Reader, size int64) (Avatar, error)
}

type Repository interface {
	UploadAvatar(avatar Avatar) (string, error)
	UploadPlaylistImage(avatar Avatar) (string, error)
	Remove(path string) error
}

const MaxAvatarSize = 16 * 1024 * 1024

var (
	ErrAvatarDoesNotExist = errors.New("avatar does not exist")
	ErrAvatarIsTooLarge   = errors.New("file is too large")
	ErrCannotRead         = errors.New("cannot read file")
	ErrWrongAvatarType    = errors.New("wrong avatar type")
)
