package survey_repository

import (
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
)

type Postgres struct {
	Pool   postgres.PgxIFace
	logger *logrus.Logger
}

func NewPostgres(pool postgres.PgxIFace, logger *logrus.Logger) Postgres {
	return Postgres{
		Pool:   pool,
		logger: logger,
	}
}

func (p *Postgres) SubmitSurveyAnswers(userId string, surveyId uint64, answers map[int]string) error {}

func (p *Postgres) IsUserSubmitSurvey(userId string, surveyId uint64) (bool, error) {}
func (p *Postgres) GetSurveyStats(surveyId uint64) (StatResponse, error) {}
func (p *Postgres) IsUserSubmitSurvey(userId string, surveyId uint64) (bool, error) {}
func (p *Postgres) IsUserSubmitSurvey(userId string, surveyId uint64) (bool, error) {}


	Get(title string) (Response, error)
	GetAllStats() ([]StatResponse, error)
