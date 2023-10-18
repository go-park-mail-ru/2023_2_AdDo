package user_usecase

import (
	"github.com/sirupsen/logrus"
	"main/internal/pkg/session"
	"main/internal/pkg/user"
)

type WithStatefulSessions struct {
	UserRepo user_domain.Repository
	AuthRepo session.Repository
	logger   *logrus.Logger
}

func NewWithStatefulSessions(userRepo user_domain.Repository, authRepo session.Repository, logger *logrus.Logger) WithStatefulSessions {
	return WithStatefulSessions{
		UserRepo: userRepo,
		AuthRepo: authRepo,
		logger:   logger,
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
		return user_domain.User{}, nil
	}
	useCase.logger.Infoln("Got user id ", id)

	u, err := useCase.UserRepo.GetById(id)
	if err != nil {
		return u, err
	}
	useCase.logger.Infoln("Got user from database ", u.Email)

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
	useCase.logger.Infoln("session deleted from database")

	return nil
}
