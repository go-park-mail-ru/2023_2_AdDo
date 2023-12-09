package cluster_domain

import "main/internal/pkg/track"

type Repository interface {
	GetNearestTracks(ids []track.Id, countPerTrack int) ([]track.Id, error)
}
