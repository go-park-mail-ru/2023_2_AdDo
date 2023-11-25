package grpc_survey

import (
	"context"
	survey_proto "main/internal/microservices/survey/proto"
	survey_domain "main/internal/pkg/survey"
)

type Client struct {
	surveyClient survey_proto.SurveyServiceClient
}

func NewClient(client survey_proto.SurveyServiceClient) Client {
	return Client{
		surveyClient: client,
	}
}

func SerializeUint64ToString(in map[uint64]string) []*survey_proto.Uint64ToString {
	answers := make([]*survey_proto.Uint64ToString, len(in))
	for questionId, answerText := range in {
		answers = append(answers, &survey_proto.Uint64ToString{
			Key:   uint64(questionId),
			Value: answerText,
		})
	}
	return answers
}

func DeserializeAnswers(in []*survey_proto.Uint64ToString) map[uint64]string {
	result := make(map[uint64]string, len(in))
	for _, pair := range in {
		result[pair.GetKey()] = pair.GetValue()
	}
	return result
}

func DeserializeStringToUint64(in []*survey_proto.StringToUint64) map[string]uint64 {
	result := make(map[string]uint64, len(in))
	for _, pair := range in {
		result[pair.Key] = pair.Value
	}
	return result
}

func (c *Client) SubmitSurvey(userId string, surveyId uint64, answer map[uint64]string) error {
	_, err := c.surveyClient.SubmitSurvey(context.Background(), &survey_proto.Survey{
		UserSurvey: &survey_proto.UserSurvey{
			UserId:   userId,
			SurveyId: surveyId,
		},
		Answers: SerializeUint64ToString(answer),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) IsSubmit(userId string, surveyId uint64) (bool, error) {
	isSubmit, err := c.surveyClient.IsSubmit(context.Background(), &survey_proto.UserSurvey{
		UserId:   userId,
		SurveyId: surveyId,
	})
	if err != nil {
		return false, err
	}
	return isSubmit.Ok, nil
}

func (c *Client) GetSurveyStats(surveyId uint64) (survey_domain.StatResponse, error) {
	statResponse, err := c.surveyClient.GetSurveyStats(context.Background(), &survey_proto.SurveyId{
		SurveyId: surveyId,
	})
	if err != nil {
		return survey_domain.StatResponse{}, err
	}
	return survey_domain.StatResponse{
		Id:                statResponse.SurveyId,
		QuestionToAverage: DeserializeStringToUint64(statResponse.QuestionToAverage),
	}, nil
}

func (c *Client) Get(surveyId uint64) (survey_domain.Response, error) {
	response, err := c.surveyClient.Get(context.Background(), &survey_proto.SurveyId{
		SurveyId: surveyId,
	})
	if err != nil {
		return survey_domain.Response{}, err
	}
	return survey_domain.Response{
		Id:               response.SurveyId,
		QuestionIdToText: DeserializeAnswers(response.QuestionIdToText),
	}, nil
}

func (c *Client) GetAllStats() ([]survey_domain.StatResponse, error) {
	return []survey_domain.StatResponse{}, nil
}
