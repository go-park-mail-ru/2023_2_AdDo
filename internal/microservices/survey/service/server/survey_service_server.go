package grpc_survey_server

import (
	"context"
	google_proto "github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	survey_proto "main/internal/microservices/survey/proto"
	"main/internal/pkg/survey"
)

type SurveyManager struct {
	repoSurvey survey.Repository
	logger     *logrus.Logger
	survey_proto.UnimplementedSurveyServiceServer
}

func NewSurveyManager(rs survey.Repository, logger *logrus.Logger) SurveyManager {
	return SurveyManager{
		repoSurvey: rs,
		logger:     logger,
	}
}

func DeserializeAnswers(in []*survey_proto.Uint64ToString) map[uint64]string {
	result := make(map[uint64]string)
	for _, pair := range in {
		result[pair.GetKey()] = pair.GetValue()
	}
	return result
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

func SerializeStringToUint64(in map[string]uint64) []*survey_proto.StringToUint64 {
	answers := make([]*survey_proto.StringToUint64, len(in))
	for answerText, questionId := range in {
		answers = append(answers, &survey_proto.StringToUint64{
			Key:   answerText,
			Value: questionId,
		})
	}
	return answers
}

func SerializeStatResponse(in survey.StatResponse) *survey_proto.StatResponse {
	return &survey_proto.StatResponse{
		SurveyId:          0,
		QuestionToAverage: SerializeStringToUint64(in.QuestionToAverage),
	}
}

func (sm *SurveyManager) SubmitSurvey(ctx context.Context, s *survey_proto.Survey) (*google_proto.Empty, error) {
	sm.logger.Infoln("Survey Manager SubmitSurvey entered")

	err := sm.repoSurvey.SubmitSurveyAnswers(s.GetUserSurvey().GetUserId(), s.GetUserSurvey().GetSurveyId(), DeserializeAnswers(s.GetAnswers()))
	if err != nil {
		return nil, err
	}
	return &google_proto.Empty{}, nil
}

func (sm *SurveyManager) IsSubmit(ctx context.Context, us *survey_proto.UserSurvey) (*survey_proto.IsOk, error) {
	sm.logger.Infoln("Survey Manager IsSubmit entered")

	isSubmit, err := sm.repoSurvey.IsUserSubmitSurvey(us.GetUserId(), us.GetSurveyId())
	if err != nil {
		return nil, err
	}

	return &survey_proto.IsOk{Ok: isSubmit}, nil
}

func (sm *SurveyManager) GetSurveyStats(ctx context.Context, id *survey_proto.SurveyId) (*survey_proto.StatResponse, error) {
	sm.logger.Infoln("Survey Manager GetSurveyStats entered")

	stat, err := sm.repoSurvey.GetSurveyStats(id.GetSurveyId())
	if err != nil {
		return nil, err
	}

	return SerializeStatResponse(stat), nil
}

func (sm *SurveyManager) Get(ctx context.Context, id *survey_proto.SurveyId) (*survey_proto.Response, error) {
	sm.logger.Infoln("Survey Manager Get entered")

	result, err := sm.repoSurvey.Get(id.GetSurveyId())
	if err != nil {
		return nil, err
	}

	return &survey_proto.Response{
		SurveyId:         id.GetSurveyId(),
		QuestionIdToText: SerializeUint64ToString(result.QuestionIdToText),
	}, nil
}

func (sm *SurveyManager) GetAllStats(ctx context.Context, empty *google_proto.Empty) (*survey_proto.StatResponses, error) {
	sm.logger.Infoln("Survey Manager GetAllStats entered")

	return nil, nil
}
