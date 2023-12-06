package candidate

import (
	"main/internal/pkg/track"
)

type UseCase interface {
	GetCandidateForUser(userId string) ([]track.Response, error)
}

type Repository interface {
	GetSimilarTracksFromCluster() ([]track.Response, error)
}
