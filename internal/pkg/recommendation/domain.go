package recommendation

import "main/internal/pkg/track"

const MyWaveTrackBatch = 5
const UserTrackPoolSize = 50

type PoolRepository interface {
	SaveTracksToUserPool(userId string, tracks []track.Response) error
	GetTracksFromUserPool(userId string, count uint32) []track.Response
}

// UseCase нейронки которая классифицирует треки
type ServiceUseCase interface {
	//GetMyWaveMusic(userId string, count uint32) ([]track.Response, error)
	ClassifyCandidates(userId string, candidates []track.Response) ([]track.Response, error)
}

// забыл, что этот делает, видимо activity worker
type WorkerUseCase interface {
	Run() error
}
