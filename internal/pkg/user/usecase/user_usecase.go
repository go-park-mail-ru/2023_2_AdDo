package user_usecase

import (
	"fmt"
	avatar_domain "main/internal/pkg/avatar"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
	"mime/multipart"
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

func (useCase *WithStatefulSessions) UploadAvatar(sessionId string, src multipart.File, size int64) error {
	id, err := useCase.AuthRepo.Get(sessionId)
	if err != nil {
		return session.ErrSessionDoesNotExist
	}

	oldPath, _ := useCase.UserRepo.GetAvatarPath(id)
	fmt.Printf("oldPath is <%s>\n", oldPath)

	path, err := useCase.AvatarRepo.Create(avatar_domain.Avatar{
		Payload:     src,
		PayloadSize: size,
		UserId:      id,
	})
	if err != nil {
		return err
	}
	fmt.Println("Create avatar record")

	err = useCase.UserRepo.UpdateAvatarPath(id, path)
	if err != nil {
		return err
	}
	fmt.Println("Update db")

	if oldPath != "" {
		useCase.AvatarRepo.Remove(oldPath)
	}
	fmt.Println("Remove old file")

	return nil
}

// docs here
func (useCase *WithStatefulSessions) RemoveAvatar(sessionId string) error {
	id, err := useCase.AuthRepo.Get(sessionId)
	if err != nil {
		return session.ErrSessionDoesNotExist
	}

	oldPath, _ := useCase.UserRepo.GetAvatarPath(id)
	fmt.Printf("oldPath is <%s>\n", oldPath)

	if oldPath == "" {
		return user_domain.ErrAvatarDoesNotExist
	}

	useCase.AvatarRepo.Remove(oldPath)
	fmt.Println("Remove old file")

	err = useCase.UserRepo.RemoveAvatarPath(id)
	if err != nil {
		return err
	}
	fmt.Println("Remove record from db")

	return nil
}
