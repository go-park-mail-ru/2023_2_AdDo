package user_usecase

import (
	"bytes"
	"io"
	avatar_domain "main/internal/pkg/avatar"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
	"net/http"
	"strings"
)

type WithStatefulSessions struct {
	UserRepo   user_domain.Repository
	AuthRepo   session.Repository
	AvatarRepo avatar_domain.S3Repository
}

func NewWithStatefulSessions(userRepo user_domain.Repository, authRepo session.Repository, avatarRepo avatar_domain.S3Repository) WithStatefulSessions {
	return WithStatefulSessions{
		UserRepo:   userRepo,
		AuthRepo:   authRepo,
		AvatarRepo: avatarRepo,
	}
}

func (useCase *WithStatefulSessions) Register(user user_domain.User) error {
	err := useCase.UserRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *WithStatefulSessions) Login(email, password string) (string, error) {
	id, err := useCase.UserRepo.CheckEmailAndPassword(email, password)
	if err != nil {
		return "", err
	}

	sessionId, err := useCase.AuthRepo.Create(id)
	if err != nil {
		return "", err
	}

	return sessionId, nil
}

func (useCase *WithStatefulSessions) GetUserInfo(sessionId string) (user_domain.User, error) {
	id, err := useCase.AuthRepo.Get(sessionId)
	if err != nil {
		return user_domain.User{}, nil
	}

	u, err := useCase.UserRepo.GetById(id)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (useCase *WithStatefulSessions) Auth(sessionId string) (bool, error) {
	_, err := useCase.AuthRepo.Get(sessionId)
	if err != nil {
		return false, session.ErrSessionDoesNotExist
	}

	return true, nil
}

func (useCase *WithStatefulSessions) Logout(sessionId string) error {
	err := useCase.AuthRepo.Delete(sessionId)
	if err != nil {
		return session.ErrSessionDoesNotExist
	}
	return nil
}

func (useCase *WithStatefulSessions) UploadAvatar(id uint64, src io.Reader, size int64) error {
	oldPath, _ := useCase.UserRepo.GetAvatarPath(id)

	if size > avatar_domain.MaxAvatarSize {
		return avatar_domain.ErrAvatarIsTooLarge
	}

	data, err := io.ReadAll(src)
	if err != nil {
		return avatar_domain.ErrCannotRead
	}
	metadata := data[:512]
	contentType := http.DetectContentType(metadata)

	if !strings.HasPrefix(contentType, "image/") {
		return avatar_domain.ErrWrongAvatarType
	}

	src = bytes.NewReader(data)
	url, err := useCase.AvatarRepo.Create(
		avatar_domain.Avatar{
			Payload:     src,
			PayloadSize: size,
			UserId:      id,
			ContentType: contentType,
		},
	)
	if err != nil {
		return err
	}

	err = useCase.UserRepo.UpdateAvatarPath(id, url)
	if err != nil {
		return err
	}

	if oldPath != "" {
		useCase.AvatarRepo.Remove(oldPath)
	}

	return nil
}

func (useCase *WithStatefulSessions) RemoveAvatar(id uint64) error {
	oldPath, _ := useCase.UserRepo.GetAvatarPath(id)

	if oldPath == "" {
		return avatar_domain.ErrAvatarDoesNotExist
	}

	useCase.AvatarRepo.Remove(oldPath)

	err := useCase.UserRepo.RemoveAvatarPath(id)
	if err != nil {
		return err
	}

	return nil
}
