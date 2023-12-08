package recommendation

import "main/internal/pkg/track"

// Юзкейс Нейронки, которую мы кормим кандидатами, она в свою очередь будет классифицировать и ранжировать их
type ServiceUseCase interface {
	ClassifyCandidates(userId string, candidates []track.Response) ([]track.Response, error)
}
