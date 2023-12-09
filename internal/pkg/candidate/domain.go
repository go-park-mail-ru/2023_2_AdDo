package candidate

import (
	"main/internal/pkg/track"
)

type UseCase interface {
	GetCandidateForDaily(userId string) ([]track.Response, error)
	GetCandidateForWave(userId string) ([]track.Response, error)
}
