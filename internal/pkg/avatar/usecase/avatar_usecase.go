package avatar_usecase

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"main/internal/pkg/avatar"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

type Default struct{}

func NewDefault() Default {
	return Default{}
}

func (useCase Default) generateAvatarName(id uint64, extansion string) string {
	data := []byte(strconv.FormatUint(id, 10) + time.Now().Format("20060102150405"))
	hash := md5.Sum(data)
	hashString := hex.EncodeToString(hash[:])
	return hashString + "." + extansion
}

func (useCase Default) GetAvatar(id uint64, src io.Reader, size int64) (avatar.Avatar, error) {
	if size > avatar.MaxAvatarSize {
		return avatar.Avatar{}, avatar.ErrAvatarIsTooLarge
	}

	data, err := io.ReadAll(src)
	if err != nil {
		return avatar.Avatar{}, avatar.ErrCannotRead
	}

	contentType := http.DetectContentType(data[:512])
	if !strings.HasPrefix(contentType, "image/") {
		return avatar.Avatar{}, avatar.ErrWrongAvatarType
	}

	src = bytes.NewReader(data)
	if err != nil {
		return avatar.Avatar{}, err
	}

	name := useCase.generateAvatarName(id, path.Base(contentType))
	return avatar.Avatar{
		Name:        name,
		Payload:     src,
		PayloadSize: size,
		ContentType: contentType,
	}, nil
}
