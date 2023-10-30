package grpc_server_user

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	session_proto "main/internal/microservices/session/proto"
	user_proto "main/internal/microservices/user/proto"
	"main/internal/pkg/session"
	user_domain "main/internal/pkg/user"
)

type UserManager struct {
	UserRepo user_domain.Repository
	AuthRepo session.Repository
	Logger   *logrus.Logger
	user_proto.UnimplementedUserServiceServer
}

func NewUserManager(userRepo user_domain.Repository, authRepo session.Repository, logger *logrus.Logger) *UserManager {
	return &UserManager{
		UserRepo: userRepo,
		AuthRepo: authRepo,
		Logger:   logger,
	}
}

func DeserializeUserData(in *user_proto.UserData) user_domain.User {
	return user_domain.User{
		Id:        in.GetId(),
		Username:  in.GetUsername(),
		Email:     in.GetEmail(),
		Password:  in.GetPassword(),
		BirthDate: in.GetBirthDate(),
		Avatar:    in.GetAvatar(),
	}
}

func SerializeUserData(in user_domain.User) *user_proto.UserData {
	return &user_proto.UserData{
		Id:        in.Id,
		Username:  in.Username,
		Email:     in.Email,
		Password:  in.Password,
		BirthDate: in.BirthDate,
		Avatar:    in.Avatar,
	}
}

func (us *UserManager) Register(ctx context.Context, in *user_proto.UserData) (*google_proto.Empty, error) {
	us.Logger.Infoln("User Micros Register entered")

	err := us.UserRepo.Create(DeserializeUserData(in))
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("New User Created")

	return &google_proto.Empty{}, nil
}

func (us *UserManager) LogIn(ctx context.Context, in *user_proto.UserCredentials) (*session_proto.SessionId, error) {
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

	return &session_proto.SessionId{SessionId: sessionId}, nil
}

func (us *UserManager) Auth(ctx context.Context, in *session_proto.SessionId) (*session_proto.Status, error) {
	us.Logger.Infoln("User Micros Auth entered")

	_, err := us.AuthRepo.Get(in.GetSessionId())
	if err != nil {
		return &session_proto.Status{IsOk: false}, session.ErrSessionDoesNotExist
	}
	us.Logger.Infoln("Got User Session From Database")

	return &session_proto.Status{IsOk: true}, nil
}

func (us *UserManager) GetUserInfo(ctx context.Context, in *session_proto.SessionId) (*user_proto.UserData, error) {
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

func (us *UserManager) LogOut(ctx context.Context, in *session_proto.SessionId) (*google_proto.Empty, error) {
	us.Logger.Infoln("User Micros LogOut entered")

	err := us.AuthRepo.Delete(in.GetSessionId())
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("session deleted from database")

	return &google_proto.Empty{}, nil
}

func (us *UserManager) UploadAvatar(ctx context.Context, in *user_proto.ImageToUser) (*google_proto.Empty, error) {
	us.Logger.Infoln("User Micros UploadAvatar entered")

	err := us.UserRepo.UpdateAvatarPath(in.GetId().GetUserId(), in.GetUrl().GetUrl())
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("Avatar Path updated")

	return &google_proto.Empty{}, nil
}

func (us *UserManager) RemoveAvatar(ctx context.Context, in *session_proto.UserId) (*user_proto.ImageUrl, error) {
	us.Logger.Infoln("User Micros RemoveAvatar entered")

	imageUrl, err := us.UserRepo.RemoveAvatarPath(in.GetUserId())
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("Avatar removed")

	return &user_proto.ImageUrl{Url: imageUrl}, nil
}

func (us *UserManager) UpdateUserInfo(ctx context.Context, in *user_proto.UserData) (*google_proto.Empty, error) {
	us.Logger.Infoln("User Micros UpdateUserInfo entered")

	err := us.UserRepo.UpdateUserInfo(DeserializeUserData(in))
	if err != nil {
		return nil, err
	}
	us.Logger.Infoln("Info updated")

	return &google_proto.Empty{}, nil
}
