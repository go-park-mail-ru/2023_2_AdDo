package candidate_domain

import (
	"main/internal/pkg/track"
)

type UseCase interface {
	GetCandidateForUser(userId string) ([]track.Response, error)
}
