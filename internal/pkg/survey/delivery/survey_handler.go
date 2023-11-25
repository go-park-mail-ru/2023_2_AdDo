package survey_delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/pkg/session"
	"net/http"
	"strconv"
)

type SurveyHandler struct {
	surveyUseCase  survey.UseCase
	sessionUseCase session.UseCase
}

func NewHandler(survey survey.UseCase, session session.UseCase) SurveyHandler {
	return SurveyHandler{
		surveyUseCase:  survey,
		sessionUseCase: session,
	}
}

// /submit/id
func (handler *SurveyHandler) Submit(w http.ResponseWriter, r *http.Request) error {
	surveyId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	sessionId, err := response.GetCookie(r)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	userId, err := handler.sessionUseCase.GetUserId(sessionId)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	var Survey map[int]string
	if err = json.NewDecoder(r.Body).Decode(&Survey); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	//handler.surveyUseCase.
}
