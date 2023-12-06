package cluster

import "main/internal/pkg/track"

type Repository interface {
	GetNearestTracks(trackId uint64) ([]track.Response, error)
}
