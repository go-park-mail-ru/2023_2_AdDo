package recommendation

import "main/internal/pkg/track"

const MyWaveTrackBatch = 5
const UserTrackPoolSize = 50

type PoolRepository interface {
	SaveTracksToUserPool(userId string, tracks []track.Response) error
	GetTracksFromUserPool(userId string, count uint32) []track.Response
}

type ServiceUseCase interface {
	GetMyWaveMusic(userId string, count uint32) ([]track.Response, error)
}

type WorkerUseCase interface {
	Run() error
}
