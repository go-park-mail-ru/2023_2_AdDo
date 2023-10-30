package image

import (
	"bytes"
	"errors"
	"io"
	"main/internal/common/utils"
	"net/http"
	"path"
	"strings"
)

type Base struct {
	Name        string
	Payload     io.Reader
	PayloadSize int64
	ContentType string
}

func CreateImageFromSource(src io.Reader, size int64) (Base, error) {
	if size > MaxAvatarSize {
		return Base{}, ErrAvatarIsTooLarge
	}

	data, err := io.ReadAll(src)
	if err != nil {
		return Base{}, ErrCannotRead
	}

	contentType := http.DetectContentType(data[:512])
	if !strings.HasPrefix(contentType, "image/") {
		return Base{}, ErrWrongAvatarType
	}

	src = bytes.NewReader(data)
	if err != nil {
		return Base{}, err
	}

	name := utils.GetImageName(path.Base(contentType))
	return Base{
		Name:        name,
		Payload:     src,
		PayloadSize: size,
		ContentType: contentType,
	}, nil
}

type UseCase interface {
	UploadAvatar(src io.Reader, size int64) (string, error)
	UploadPlaylistImage(src io.Reader, size int64) (string, error)
	RemoveImage(url string) error
}

type Repository interface {
	UploadAvatar(image Base) (string, error)
	UploadPlaylistImage(image Base) (string, error)
	Remove(path string) error
}

const MaxAvatarSize = 16 * 1024 * 1024

var (
	ErrAvatarDoesNotExist = errors.New("image does not exist")
	ErrAvatarIsTooLarge   = errors.New("file is too large")
	ErrCannotRead         = errors.New("cannot read file")
	ErrWrongAvatarType    = errors.New("wrong image type")
)