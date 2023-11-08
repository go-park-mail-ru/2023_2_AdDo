package user_usecase

import (
	"github.com/sirupsen/logrus"
	"io"
	avatar_domain "main/internal/pkg/image"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
)

type WithStatefulSessions struct {
	UserRepo      user_domain.Repository
	AuthRepo      session.Repository
	AvatarRepo    avatar_domain.Repository
	AvatarUseCase avatar_domain.UseCase
	logger        *logrus.Logger
}

func (useCase *WithStatefulSessions) UpdateUserInfo(userId string, user user_domain.User) error {
	//TODO implement me
	panic("implement me")
}

func NewWithStatefulSessions(
	userRepo user_domain.Repository,
	authRepo session.Repository,
	avatarRepo avatar_domain.Repository,
	avatarUseCase avatar_domain.UseCase,
	logger *logrus.Logger,
) WithStatefulSessions {
	return WithStatefulSessions{
		UserRepo:      userRepo,
		AuthRepo:      authRepo,
		AvatarRepo:    avatarRepo,
		AvatarUseCase: avatarUseCase,
		logger:        logger,
	}
}

func (useCase *WithStatefulSessions) Register(user user_domain.User) error {
	useCase.logger.Infoln("UserUseCase Register entered")

	err := useCase.UserRepo.Create(user)
	if err != nil {
		return err
	}
	useCase.logger.Infoln("New User Created")

	return nil
}

func (useCase *WithStatefulSessions) Login(email, password string) (string, error) {
	useCase.logger.Infoln("UserUseCase Login entered")

	id, err := useCase.UserRepo.CheckEmailAndPassword(email, password)
	if err != nil {
		return "", err
	}
	useCase.logger.Infoln("User credentials are successfully checked")

	sessionId, err := useCase.AuthRepo.Create(id)
	if err != nil {
		return "", err
	}
	useCase.logger.Infoln("New session id created for user ", email)

	return sessionId, nil
}

func (useCase *WithStatefulSessions) GetUserInfo(sessionId string) (user_domain.User, error) {
	useCase.logger.Infoln("UserUseCase GetUserInfo entered")

	id, err := useCase.AuthRepo.Get(sessionId)
	if err != nil {
		return user_domain.User{}, err
	}
	useCase.logger.Infoln("Got user id ", id)

	u, err := useCase.UserRepo.GetById(id)
	if err != nil {
		return u, err
	}
	useCase.logger.Infoln("Got user from db ", u.Email)

	return u, nil
}

func (useCase *WithStatefulSessions) Auth(sessionId string) (bool, error) {
	useCase.logger.Infoln("UserUseCase Auth entered")

	_, err := useCase.AuthRepo.Get(sessionId)
	if err != nil {
		return false, session.ErrSessionDoesNotExist
	}
	useCase.logger.Infoln("Got User Session From Database")

	return true, nil
}

func (useCase *WithStatefulSessions) Logout(sessionId string) error {
	useCase.logger.Infoln("UserUseCase LogOut entered")

	err := useCase.AuthRepo.Delete(sessionId)
	if err != nil {
		return session.ErrSessionDoesNotExist
	}
	useCase.logger.Infoln("session deleted from db")

	return nil
}

func (useCase *WithStatefulSessions) UploadAvatar(id string, src io.Reader, size int64) (string, error) {
	//oldPath, _ := useCase.UserRepo.GetAvatarPath(id)

	//avatar, err := useCase.AvatarUseCase.GetAvatar(id, src, size)
	//if err != nil {
	//	return "", err
	//}
	//
	//url, err := useCase.AvatarRepo.UploadAvatar(avatar)
	//if err != nil {
	//	return "", err
	//}

	//err = useCase.UserRepo.UpdateAvatarPath(id, url)
	//if err != nil {
	//	return "", err
	//}
	//
	//if oldPath != "" {
	//	useCase.AvatarRepo.Remove(oldPath)
	//}
	//
	//return url, nil
	return "", nil
}

//func (useCase *WithStatefulSessions) RemoveAvatar(id string) error {
//	oldPath, err := useCase.UserRepo.GetAvatarPath(id)
//	if err != nil {
//		return err
//	}
//
//	if oldPath == "" {
//		return avatar_domain.ErrAvatarDoesNotExist
//	}
//
//	useCase.AvatarRepo.Remove(oldPath)
//
//	err = useCase.UserRepo.RemoveAvatarPath(id)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
