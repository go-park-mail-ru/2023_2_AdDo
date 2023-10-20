package server_auth

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "main/internal/microservices/auth/proto"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
)

type AuthManager struct {
	pb.UnimplementedAuthServiceServer
	session session.UseCase
	user    user_domain.UseCase
	logger  *logrus.Logger
}

func NewAuthManager(sessionUseCase session.UseCase, userUseCase user_domain.UseCase, logger *logrus.Logger) *AuthManager {
	return &AuthManager{
		session: sessionUseCase,
		user:    userUseCase,
		logger:  logger,
	}
}

func DeserializeUser(in *pb.UserData) user_domain.User {
	return user_domain.User{
		Id:        in.GetId(),
		Username:  in.GetUsername(),
		Email:     in.GetEmail(),
		Password:  in.GetPassword(),
		BirthDate: in.GetBirthDate(),
		Avatar:    in.GetAvatar(),
	}
}

func DeserializeUserCredentials(in *pb.UserCredentials) user_domain.UserCredentials {
	return user_domain.UserCredentials{
		Email:    in.GetEmail(),
		Password: in.GetPassword(),
	}
}

func SerializeUser(in user_domain.User) *pb.UserData {
	return &pb.UserData{
		Id:        in.Id,
		Email:     in.Email,
		Password:  in.Password,
		Username:  in.Username,
		BirthDate: in.BirthDate,
		Avatar:    in.Avatar,
	}
}

func (am *AuthManager) SignUp(ctx context.Context, in *pb.UserData) (*pb.SessionId, error) {
	am.logger.Infoln("SignUp in Auth Service entered")

	user := DeserializeUser(in)
	am.logger.Infoln("UserData deserialized", user)

	err := user.Validate()
	if err != nil {
		return &pb.SessionId{}, err
	}
	am.logger.Infoln("User model is valid")

	err = am.user.Register(user)
	if err != nil {
		return &pb.SessionId{}, err
	}
	am.logger.Infoln("User signed up")

	sessionId, err := am.user.Login(user.Email, user.Password)
	if err != nil {
		return &pb.SessionId{}, err
	}
	am.logger.Infoln("Session for user created")

	return &pb.SessionId{SessionId: sessionId}, nil
}

func (am *AuthManager) LogIn(ctx context.Context, in *pb.UserCredentials) (*pb.SessionId, error) {
	am.logger.Infoln("LogIn in Auth Service entered")

	credentials := DeserializeUserCredentials(in)
	am.logger.Infoln("User credentials deserialized", credentials)

	err := credentials.Validate()
	if err != nil {
		return &pb.SessionId{}, err
	}
	am.logger.Infoln("User credentials are valid")

	sessionId, err := am.user.Login(credentials.Email, credentials.Password)
	if err != nil {
		return &pb.SessionId{}, err
	}
	am.logger.Infoln("User got a new session id")

	return &pb.SessionId{SessionId: sessionId}, nil
}

func (am *AuthManager) LogOut(ctx context.Context, in *pb.SessionId) (*pb.Status, error) {
	am.logger.Infoln("LogOut in Auth Service entered")

	sessionId := in.GetSessionId()
	am.logger.Infoln("Got sessionId ", sessionId)

	if err := am.user.Logout(sessionId); err != nil {
		return &pb.Status{Code: 403}, err
	}
	am.logger.Infoln("Session deleted from database")

	return &pb.Status{Code: 200}, nil
}

func (am *AuthManager) Auth(ctx context.Context, in *pb.SessionId) (*pb.Status, error) {
	am.logger.Infoln("Auth Procedure in Auth Service entered")

	sessionId := in.GetSessionId()
	am.logger.Infoln("Got sessionId ", sessionId)

	isAuth, err := am.user.Auth(sessionId)
	if err != nil || !isAuth {
		return &pb.Status{Code: 401}, nil
	}
	am.logger.Infoln("auth success")

	return &pb.Status{Code: 200}, nil
}

func (am *AuthManager) UserInfo(ctx context.Context, in *pb.SessionId) (*pb.UserData, error) {
	am.logger.Infoln("UserInfo in Auth Service entered")

	sessionId := in.GetSessionId()
	am.logger.Infoln("Got sessionId ", sessionId)

	user, err := am.user.GetUserInfo(sessionId)
	if err != nil {
		return &pb.UserData{}, err
	}
	am.logger.Infoln("Got user info from database successfully", user)

	return SerializeUser(user), nil
}
