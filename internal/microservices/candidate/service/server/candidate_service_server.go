package candidate_service_server

import (
	"context"
	"github.com/sirupsen/logrus"
	candidate_proto "main/internal/microservices/candidate/proto"
	session_proto "main/internal/microservices/session/proto"
	grpc_track_server "main/internal/microservices/track/service/server"
	"main/internal/pkg/activity"
	"main/internal/pkg/cluster"
	"main/internal/pkg/track"
)

type CandidateManager struct {
	logger      *logrus.Logger
	trackRepo   track.Repository
	clusterRepo cluster_domain.Repository
	candidate_proto.UnimplementedCandidateServiceServer
	recentActivityRepo activity.KeyValueRepository
}

func NewCandidateManager(ar activity.KeyValueRepository, tr track.Repository, cr cluster_domain.Repository, logger *logrus.Logger) CandidateManager {
	return CandidateManager{
		logger:             logger,
		trackRepo:          tr,
		clusterRepo:        cr,
		recentActivityRepo: ar,
	}
}

//Daily
// Includes
//  The Hottest User Tracks
//  The Activity for Last day
// Excludes
//  UserBlackListMusic
//  UninterestedTracks(fast skip, less then 10% all dur)

const HotTrackDailyNum = 5              // 10
const DailyTrackCandidatePoolSize = 250 // 500
func (cm *CandidateManager) GetCandidatesForDailyPlaylist(ctx context.Context, id *session_proto.UserId) (*candidate_proto.Candidates, error) {
	cm.logger.Infoln("Candidate Micros Get Candidates For Daily")

	hotTracks, err := cm.trackRepo.GetHotTracks(id.GetUserId(), HotTrackDailyNum)
	if err != nil {
		return nil, err
	}

	lastDayTracks, err := cm.trackRepo.GetLastDayTracks(id.GetUserId())
	if err != nil {
		return nil, err
	}

	countPerTrack := (len(hotTracks) + len(lastDayTracks)) / DailyTrackCandidatePoolSize

	tracksByHot, err := cm.clusterRepo.GetNearestTracks(hotTracks, countPerTrack)
	if err != nil {
		return nil, err
	}

	trackForLastDay, err := cm.clusterRepo.GetNearestTracks(lastDayTracks, countPerTrack)
	if err != nil {
		return nil, err
	}

	ids := append(tracksByHot, trackForLastDay...)
	// excluding

	result, err := cm.trackRepo.GetTracksByIds(ids)
	if err != nil {
		return nil, err
	}

	return &candidate_proto.Candidates{
		Tracks: grpc_track_server.SerializeTracks(result),
	}, nil
}

//Wave
// Includes
//  The Hottest User Tracks
//  Recent Activity
//  Rotation Tracks
// Excludes
//  UninterestedTracks(fast skip, less then 10% all dur)
//  UserBlackListMusic
//  This Wave music(delivery)

const HotTrackWaveNum = 3
const WaveTrackCandidatePoolSize = 100

func (cm *CandidateManager) GetCandidatesForWave(ctx context.Context, id *session_proto.UserId) (*candidate_proto.Candidates, error) {
	cm.logger.Infoln("Candidate Micros Get Candidates For Wave")

	hotTracks, err := cm.trackRepo.GetHotTracks(id.GetUserId(), HotTrackWaveNum)
	if err != nil {
		return nil, err
	}

	recentActivity, err := cm.recentActivityRepo.GetAllRecentActivity(id.GetUserId())
	if err != nil {
		return nil, err
	}
	recentActivityIds := make([]track.Id, 0)
	for _, act := range recentActivity {
		recentActivityIds = append(recentActivityIds, track.Id{Id: act.TrackId})
	}

	countPerTrack := (len(hotTracks) + len(recentActivity)) / WaveTrackCandidatePoolSize

	tracksByHot, err := cm.clusterRepo.GetNearestTracks(hotTracks, countPerTrack)
	if err != nil {
		return nil, err
	}

	trackByRecentActivity, err := cm.clusterRepo.GetNearestTracks(recentActivityIds, countPerTrack)
	if err != nil {
		return nil, err
	}

	ids := append(tracksByHot, trackByRecentActivity...)
	// excluding

	result, err := cm.trackRepo.GetTracksByIds(ids)
	if err != nil {
		return nil, err
	}

	return &candidate_proto.Candidates{
		Tracks: grpc_track_server.SerializeTracks(result),
	}, nil
}
