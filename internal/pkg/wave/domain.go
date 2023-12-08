package wave

import "main/internal/pkg/track"

type UseCase interface {
	GetMyWaveMusic(userId string, count uint32) ([]track.Response, error)
}

const UserTrackPoolSize = 50

type PoolRepository interface {
	SaveTracksToUserPool(userId string, tracks []track.Response) error
	GetTracksFromUserPool(userId string, count uint32) []track.Response
}
