package survey_delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/common/utils"
	"main/internal/pkg/session"
	"main/internal/pkg/survey"
	"net/http"
	"strconv"
)

type SurveyHandler struct {
	surveyUseCase  survey.UseCase
	sessionUseCase session.UseCase
	logger         *logrus.Logger
}

func NewHandler(survey survey.UseCase, session session.UseCase, logger *logrus.Logger) SurveyHandler {
	return SurveyHandler{
		surveyUseCase:  survey,
		sessionUseCase: session,
		logger:         logger,
	}
}

func (handler *SurveyHandler) Submit(w http.ResponseWriter, r *http.Request) error {
	handler.logger.WithFields(logrus.Fields{
		"request_id": utils.GenReqId(r.RequestURI + r.Method),
	}).Infoln("Submit Handler entered")

	surveyId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Parsed surveyId from Vars")

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Got Cookie")

	userId, err := handler.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}
	handler.logger.Infoln("Got user id")

	var answer map[int]string
	if err = json.NewDecoder(r.Body).Decode(&answer); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	handler.logger.Infoln("Got user answer from body")

	if err = handler.surveyUseCase.SubmitSurvey(userId, uint64(surveyId), answer); err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
