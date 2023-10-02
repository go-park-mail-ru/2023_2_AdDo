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

func (useCase *WithStatefulSessions) Register(user user_domain.User) (uint64, error) {
	id, err := useCase.UserRepo.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (useCase *WithStatefulSessions) Login(email, password string) (uint64, string, error) {
	id, err := useCase.UserRepo.CheckEmailAndPassword(email, password)
	if err != nil || id == 0 {
		return 0, "", user_domain.ErrUserDoesNotExist
	}

	sessionId, err := useCase.AuthRepo.Create(id)
	if err != nil || id == 0 {
		return 0, "", user_domain.ErrSessionCreatingFailed
	}

	return id, sessionId, nil
}

func (useCase *WithStatefulSessions) GetUserInfo(id uint64) (user_domain.User, error) {
	u, err := useCase.UserRepo.GetById(id)
	if err != nil {
		return u, user_domain.ErrUserDoesNotExist
	}
	return u, nil
}

func (useCase *WithStatefulSessions) Auth(id uint64, sessionId string) (bool, error) {
	sessionIdInDb, err := useCase.AuthRepo.GetByUserId(id)
	if err != nil {
		return false, user_domain.ErrSessionDoesNotExist
	}
	return sessionId == sessionIdInDb, nil
}

func (useCase *WithStatefulSessions) Logout(id uint64) error {
	err := useCase.AuthRepo.DeleteByUserId(id)
	if err != nil {
		return user_domain.ErrSessionDoesNotExist
	}
	return nil
}
