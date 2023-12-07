package wave

import "main/internal/pkg/track"

type UseCase interface {
	GetMyWaveMusic(userId string, count uint32) ([]track.Response, error)
}
