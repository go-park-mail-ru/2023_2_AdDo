package survey_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	postgres "main/internal/common/pgxiface"
	"main/internal/pkg/survey"
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

func (p *Postgres) SubmitSurveyAnswers(userId string, surveyId uint64, answers map[uint64]string) error {
	p.logger.Infoln("Survey Repo Get entered")

	query := `insert into answer (profile_id, survey_id, question_id, answer) values ( $1, $2, `
	for id, answer := range answers {
		query += `$3, $4 )`
		if _, err := p.Pool.Exec(context.Background(), query, userId, surveyId, id, answer); err != nil {
			p.logger.WithFields(logrus.Fields{
				"query": query,
				"err":   err,
			}).Errorln("Error while getting a survey")
			return err
		}
	}

	return nil
}

func (p *Postgres) IsUserSubmitSurvey(userId string, surveyId uint64) (bool, error) {
	p.logger.Infoln("Survey Repo Get entered")

	count := 0
	query := `select count(*) from answer where profile_id = $1 and survey_id = $2`
	if err := p.Pool.QueryRow(context.Background(), query, userId, surveyId).Scan(&count); err != nil {
		p.logger.WithFields(logrus.Fields{
			"query": query,
			"err":   err,
		}).Errorln("Error while getting a survey")
		return false, err
	}

	return count > 0, nil
}

func (p *Postgres) GetSurveyStats(surveyId uint64) (survey.StatResponse, error) {
	p.logger.Infoln("Survey Repo Get entered")

	result := survey.StatResponse{Id: surveyId}
	query := `select question.title, avg(answer) as average_rating from survey
				inner join answer  on survey.id = answer.survey_id
				inner join question  on answer.question_id = question.id
				where survey_id = $1
				group by question.title, question_id`
	rows, err := p.Pool.Query(context.Background(), query, surveyId)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"query": query,
			"err":   err,
		}).Errorln("Error while getting a survey stat")
		return result, err
	}

	for rows.Next() {
		question := ""
		avg := 0
		if err := rows.Scan(&question, &avg); err != nil {
			p.logger.WithFields(logrus.Fields{
				"query": query,
				"err":   err,
			}).Errorln("Error while getting a survey stat")
			return result, err
		}

		result.QuestionToAverage[question] = uint64(avg)
	}
	return result, nil
}

func (p *Postgres) Get(id uint64) (survey.Response, error) {
	p.logger.Infoln("Survey Repo Get entered")

	var result survey.Response
	query := `select question_id, question.title from survey
    join survey_question on survey_question.survey_id = survey.id
    join question on survey_question.question_id = question.id
                              where survey.id = $1`
	rows, err := p.Pool.Query(context.Background(), query, id)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"query": query,
			"err":   err,
		}).Errorln("Error while getting a survey")
		return survey.Response{}, err
	}

	result.Id = id

	for rows.Next() {
		questionId := 0
		questionText := ""
		if err := rows.Scan(&questionId, &questionText); err != nil {
			p.logger.WithFields(logrus.Fields{
				"query": query,
				"err":   err,
			}).Errorln("Error while getting a survey")
			return survey.Response{}, err
		}

		result.QuestionIdToText[uint64(questionId)] = questionText
	}

	return result, nil
}

func (p *Postgres) GetAllStats() ([]survey.StatResponse, error) {
	p.logger.Infoln("Survey Repo GetAllStats entered")

	return nil, nil
}
