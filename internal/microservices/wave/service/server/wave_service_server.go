package grpc_wave_server

import (
	"context"
	"github.com/sirupsen/logrus"
	proto2 "main/internal/microservices/track/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/microservices/wave/proto"
	"main/internal/pkg/wave"
)

type WaveManager struct {
	logger *logrus.Logger
	proto.UnimplementedWaveServiceServer
	wavePool wave.PoolRepository
}

func NewWaveManager(wp wave.PoolRepository, logger *logrus.Logger) WaveManager {
	return WaveManager{wavePool: wp, logger: logger}
}

func (wm *WaveManager) GetMyWaveMusic(ctx context.Context, in *proto.UserIdCount) (*proto2.TracksResponse, error) {
	wm.logger.Infoln("Wave Micros get my wave music entered")

	tracks, err := wm.wavePool.GetTracksFromUserPool(in.GetUserId(), in.GetTrackCount())
	if err != nil {
		return nil, nil
	}

	return grpc_track_server.SerializeTracks(tracks), nil
}
