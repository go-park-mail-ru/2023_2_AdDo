package survey_middleware

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"main/internal/common/default_handlers"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/pkg/session"
	"main/internal/pkg/survey"
	"net/http"
	"strconv"
)

type Middleware struct {
	surveyUseCase  survey.UseCase
	sessionUseCase session.UseCase
	logger         *logrus.Logger
}

func NewMiddleware(survey survey.UseCase, session session.UseCase, logger *logrus.Logger) Middleware {
	return Middleware{
		surveyUseCase:  survey,
		sessionUseCase: session,
		logger:         logger,
	}
}

func (m *Middleware) CheckUserVote(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		surveyId, err := strconv.Atoi(mux.Vars(request)["id"])
		if err != nil {
			common_handler.Handler{H: default_handlers.BadRequest}.ServeHTTP(w, request)
			return
		}
		m.logger.Infoln("Parsed surveyId from Vars")

		sessionId, err := response.GetCookie(request)
		if err != nil {
			common_handler.Handler{H: default_handlers.NoAuth}.ServeHTTP(w, request)
			return
		}
		m.logger.Infoln("Got user cookie")

		userId, err := m.sessionUseCase.GetUserId(sessionId)
		if err != nil {
			common_handler.Handler{H: default_handlers.NoAuth}.ServeHTTP(w, request)
			return
		}
		m.logger.Infoln("Got user id")

		isSubmit, err := m.surveyUseCase.IsSubmit(userId, uint64(surveyId))
		if err != nil {
			common_handler.Handler{H: default_handlers.Internal}.ServeHTTP(w, request)
		}
		if isSubmit {
			common_handler.Handler{H: default_handlers.Forbidden}.ServeHTTP(w, request)
			return
		}
		next.ServeHTTP(w, request)
	})
}
