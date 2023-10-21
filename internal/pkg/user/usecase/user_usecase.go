package user_usecase

import (
	"io"
	avatar_domain "main/internal/pkg/avatar"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
)

type WithStatefulSessions struct {
	UserRepo      user_domain.Repository
	AuthRepo      session.Repository
	AvatarRepo    avatar_domain.Repository
	AvatarUseCase avatar_domain.UseCase
}

func NewWithStatefulSessions(
	userRepo user_domain.Repository,
	authRepo session.Repository,
	avatarRepo avatar_domain.Repository,
	avatarUseCase avatar_domain.UseCase,
) WithStatefulSessions {
	return WithStatefulSessions{
		UserRepo:      userRepo,
		AuthRepo:      authRepo,
		AvatarRepo:    avatarRepo,
		AvatarUseCase: avatarUseCase,
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

func (useCase *WithStatefulSessions) UploadAvatar(id uint64, src io.Reader, size int64) (string, error) {
	oldPath, _ := useCase.UserRepo.GetAvatarPath(id)

	avatar, err := useCase.AvatarUseCase.GetAvatar(id, src, size)
	if err != nil {
		return "", err
	}

	url, err := useCase.AvatarRepo.UploadAvatar(avatar)
	if err != nil {
		return "", err
	}

	err = useCase.UserRepo.UpdateAvatarPath(id, url)
	if err != nil {
		return "", err
	}

	if oldPath != "" {
		useCase.AvatarRepo.Remove(oldPath)
	}

	return url, nil
}

func (useCase *WithStatefulSessions) RemoveAvatar(id uint64) error {
	oldPath, err := useCase.UserRepo.GetAvatarPath(id)
	if err != nil {
		return err
	}

	if oldPath == "" {
		return avatar_domain.ErrAvatarDoesNotExist
	}

	useCase.AvatarRepo.Remove(oldPath)

	err = useCase.UserRepo.RemoveAvatarPath(id)
	if err != nil {
		return err
	}

	return nil
}
