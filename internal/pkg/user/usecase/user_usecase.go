package user_usecase

import (
	"main/internal/pkg/session"
	"main/internal/pkg/user"
)

type WithStatefulSessions struct {
	UserRepo user_domain.Repository
	AuthRepo session.Repository
}

func NewWithStatefulSessions(userRepo user_domain.Repository, authRepo session.Repository) WithStatefulSessions {
	return WithStatefulSessions{
		UserRepo: userRepo,
		AuthRepo: authRepo,
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
