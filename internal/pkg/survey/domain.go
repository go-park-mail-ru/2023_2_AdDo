package survey

type StatResponse struct {
	id                uint64
	questionToAverage map[string]int
}

type Response struct {
	id               uint64
	questionIdToText map[uint64]string
}

type UseCase interface {
	SubmitSurvey(userId string, surveyId uint64, answer map[int]string) error
	IsSubmit(userId string, surveyId uint64) (bool, error)
	GetSurveyStats(surveyId uint64) (StatResponse, error)
	Get(title string) (Response, error)
	GetAllStats() ([]StatResponse, error)
}

type Repository interface {
	SubmitSurveyAnswers(userId string, surveyId uint64, answers map[int]string) error
	IsUserSubmitSurvey(userId string, surveyId uint64) (bool, error)
	GetSurveyStats(surveyId uint64) (StatResponse, error)
	Get(title string) (Response, error)
	GetAllStats() ([]StatResponse, error)
}
