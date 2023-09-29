package user_usecase

import (
	"main/internal/pkg/session"
	"main/internal/pkg/user"
)

type WithStatefulSessions struct {
	userRepo user_domain.Repository
	authRepo session.Repository
}

func NewWithStatefulSessions(userRepo user_domain.Repository, authRepo session.Repository) WithStatefulSessions {
	return WithStatefulSessions{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}

func (useCase *WithStatefulSessions) Register(user user_domain.User) (uint64, error) {
	id, err := useCase.userRepo.Create(user)
	if err != nil {
		return 0, user_domain.ErrUserAlreadyExist
	}
	return id, nil
}

func (useCase *WithStatefulSessions) Login(email, password string) (uint64, string, error) {
	id, err := useCase.userRepo.CheckEmailAndPassword(email, password)
	if err != nil || id == 0 {
		return 0, "", user_domain.ErrUserDoesNotExist
	}

	sessionId, err := useCase.authRepo.Create(id)
	if err != nil || id == 0 {
		return 0, "", user_domain.ErrSessionCreatingFailed
	}

	return id, sessionId, nil
}

func (useCase *WithStatefulSessions) GetUserById(id uint64) (user_domain.User, error) {
	u, err := useCase.userRepo.GetById(id)
	if err != nil {
		return u, user_domain.ErrUserDoesNotExist
	}
	return u, nil
}

func (useCase *WithStatefulSessions) CreateUserSession(id uint64) (string, error) {
	sessionId, err := useCase.authRepo.Create(id)
	if err != nil {
		return "", user_domain.ErrUserDoesNotExist
	}
	return sessionId, nil
}

func (useCase *WithStatefulSessions) Auth(id uint64, sessionId string) (bool, error) {
	sessionIdInDb, err := useCase.authRepo.GetByUserId(id)
	if err != nil {
		return false, user_domain.ErrSessionDoesNotExist
	}
	return sessionId == sessionIdInDb, nil
}

func (useCase *WithStatefulSessions) Logout(id uint64) error {
	err := useCase.authRepo.DeleteByUserId(id)
	if err != nil {
		return user_domain.ErrSessionDoesNotExist
	}
	return nil
}
