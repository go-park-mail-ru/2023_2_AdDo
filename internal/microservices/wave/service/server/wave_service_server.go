package grpc_wave_server

import (
	"context"
	"github.com/sirupsen/logrus"
	proto2 "main/internal/microservices/track/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/microservices/wave/proto"
	"main/internal/pkg/track"
	"main/internal/pkg/wave"
)

type WaveManager struct {
	logger *logrus.Logger
	proto.UnimplementedWaveServiceServer
	wavePool  wave.PoolRepository
	trackRepo track.Repository
}

func NewWaveManager(tr track.Repository, wp wave.PoolRepository, logger *logrus.Logger) WaveManager {
	return WaveManager{trackRepo: tr, wavePool: wp, logger: logger}
}

func (wm *WaveManager) GetMyWaveMusic(ctx context.Context, in *proto.UserIdCount) (*proto2.TracksResponse, error) {
	wm.logger.Infoln("Wave Micros get my wave music entered")

	tracks, err := wm.trackRepo.GetWaveTracks(in.GetUserId(), in.GetTrackCount())
	if err != nil || len(tracks) == 0 {
		wm.logger.Errorln("error while getting WAVE tracks", err)
		randomTracks, err := wm.trackRepo.GetRandomTracksForWave(in.GetUserId(), in.GetTrackCount())
		if err != nil {
			wm.logger.Errorln("error while getting RANDOM tracks for wave", err)
			return nil, err
		}

		return grpc_track_server.SerializeTracks(randomTracks), nil
	}

	err = wm.trackRepo.DeleteLastTakenFromWave(in.GetUserId(), tracks)
	if err != nil {
		return nil, err
	}

	return grpc_track_server.SerializeTracks(tracks), nil
}
