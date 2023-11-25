package survey

type StatResponse struct {
	Id                uint64
	QuestionToAverage map[string]int
}

type Response struct {
	Id               uint64
	QuestionIdToText map[uint64]string
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
