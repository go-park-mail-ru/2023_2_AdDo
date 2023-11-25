package survey

type StatResponse struct {
	Id                uint64
	QuestionToAverage map[string]uint64
}

type Response struct {
	Id               uint64
	QuestionIdToText map[uint64]string `json:"Questions" example:"{1: '1', 2: '2'}"`
}

type UseCase interface {
	SubmitSurvey(userId string, surveyId uint64, answer map[uint64]string) error
	IsSubmit(userId string, surveyId uint64) (bool, error)
	GetSurveyStats(surveyId uint64) (StatResponse, error)
	Get(surveyId uint64) (Response, error)
	GetAllStats() ([]StatResponse, error)
}

type Repository interface {
	SubmitSurveyAnswers(userId string, surveyId uint64, answers map[uint64]string) error
	IsUserSubmitSurvey(userId string, surveyId uint64) (bool, error)
	GetSurveyStats(surveyId uint64) (StatResponse, error)
	Get(surveyId uint64) (Response, error)
	GetAllStats() ([]StatResponse, error)
}
