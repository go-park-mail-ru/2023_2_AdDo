package grpc_server_user

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "main/internal/microservices/user/proto"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
)

type UserManager struct {
	UserRepo user_domain.Repository
	AuthRepo session.Repository
	Logger   *logrus.Logger
	pb.UnimplementedUserServiceServer
}

func NewUserManager(userRepo user_domain.Repository, authRepo session.Repository, logger *logrus.Logger) *UserManager {
	return &UserManager{
		UserRepo: userRepo,
		AuthRepo: authRepo,
		Logger:   logger,
	}
}

func DeserializeUserData(in *pb.UserData) user_domain.User {
	return user_domain.User{
		Id:        in.GetId(),
		Username:  in.GetUsername(),
		Email:     in.GetEmail(),
		Password:  in.GetPassword(),
		BirthDate: in.GetBirthDate(),
		Avatar:    in.GetAvatar(),
	}
}

func SerializeUserData(in user_domain.User) *pb.UserData {
	return &pb.UserData{
		Id:        in.Id,
		Username:  in.Username,
		Email:     in.Email,
		Password:  in.Password,
		BirthDate: in.BirthDate,
		Avatar:    in.Avatar,
	}
}

func (us *UserManager) Register(ctx context.Context, in *pb.UserData) (*pb.Status, error) {
	us.Logger.Infoln("User Micros Register entered")

	err := us.UserRepo.Create(DeserializeUserData(in))
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("New User Created")

	return &pb.Status{IsOk: true}, nil
}

func (us *UserManager) LogIn(ctx context.Context, in *pb.UserCredentials) (*pb.SessionId, error) {
	us.Logger.Infoln("User Micros Login entered")

	id, err := us.UserRepo.CheckEmailAndPassword(in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("User credentials are successfully checked")

	sessionId, err := us.AuthRepo.Create(id)
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("New session id created for user ", in.Email)

	return &pb.SessionId{SessionId: sessionId}, nil
}

func (us *UserManager) Auth(ctx context.Context, in *pb.SessionId) (*pb.Status, error) {
	us.Logger.Infoln("User Micros Auth entered")

	_, err := us.AuthRepo.Get(in.GetSessionId())
	if err != nil {
		return &pb.Status{IsOk: false}, session.ErrSessionDoesNotExist
	}
	us.Logger.Infoln("Got User Session From Database")

	return &pb.Status{IsOk: true}, nil
}

func (us *UserManager) GetUserInfo(ctx context.Context, in *pb.SessionId) (*pb.UserData, error) {
	us.Logger.Infoln("User Micros GetUserInfo entered")

	id, err := us.AuthRepo.Get(in.GetSessionId())
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("Got user id ", id)

	u, err := us.UserRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("Got user from database ", u.Email)

	return SerializeUserData(u), nil
}

func (us *UserManager) LogOut(ctx context.Context, in *pb.SessionId) (*pb.Status, error) {
	us.Logger.Infoln("User Micros LogOut entered")

	err := us.AuthRepo.Delete(in.GetSessionId())
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("session deleted from database")

	return &pb.Status{IsOk: true}, nil
}
