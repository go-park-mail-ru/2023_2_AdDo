package recommendation

import "main/internal/pkg/track"

const UserTrackPoolSize = 50

type PoolRepository interface {
	SaveTracksToUserPool(userId string, tracks []track.Response) error
	GetTracksFromUserPool(userId string, count uint32) []track.Response
}

// Юзкейс Нейронки, которую мы кормим кандидатами, она в свою очередь будет классифицировать и ранжировать их
type ServiceUseCase interface {
	ClassifyCandidates(userId string, candidates []track.Response) ([]track.Response, error)
}
