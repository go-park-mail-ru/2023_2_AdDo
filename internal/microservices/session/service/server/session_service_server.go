package grpc_session_server

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "main/internal/microservices/session/proto"
	"main/internal/pkg/session"
)

type SessionManager struct {
	repoSession session.Repository
	logger      *logrus.Logger
	pb.UnimplementedSessionServiceServer
}

func NewSessionManager(repoSession session.Repository, logger *logrus.Logger) SessionManager {
	return SessionManager{logger: logger, repoSession: repoSession}
}

func (sm *SessionManager) CheckSession(ctx context.Context, in *pb.SessionId) (*pb.Status, error) {
	sm.logger.Infoln("Session Micros CheckSession entered")

	if _, err := sm.repoSession.Get(in.GetSessionId()); err != nil {
		return &pb.Status{IsOk: false}, err
	}
	sm.logger.Infoln("Session id matched with an database one")

	return &pb.Status{IsOk: true}, nil
}

func (sm *SessionManager) GetUserId(ctx context.Context, in *pb.SessionId) (*pb.UserId, error) {
	sm.logger.Infoln("Session Micros GetUserId entered")

	userId, err := sm.repoSession.Get(in.GetSessionId())
	if err != nil {
		return nil, err
	}
	sm.logger.Infoln("Got user id: ", userId)

	return &pb.UserId{UserId: userId}, nil
}
