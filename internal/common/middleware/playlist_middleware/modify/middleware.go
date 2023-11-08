package modify_playlist

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"main/internal/common/default_handlers"
	common_handler "main/internal/common/handler"
	"main/internal/common/response"
	"main/internal/pkg/playlist"
	"main/internal/pkg/session"
	"net/http"
	"strconv"
)

type Middleware struct {
	playlistUseCase playlist.UseCase
	sessionUseCase  session.UseCase
	logger          *logrus.Logger
}

func NewMiddleware(pu playlist.UseCase, su session.UseCase, logger *logrus.Logger) Middleware {
	return Middleware{
		playlistUseCase: pu,
		sessionUseCase:  su,
		logger:          logger,
	}
}

func (m *Middleware) ModifyPlaylistAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
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

		playlistId, err := strconv.Atoi(mux.Vars(request)["id"])
		if err != nil {
			common_handler.Handler{H: default_handlers.BadRequest}.ServeHTTP(w, request)
			return
		}
		m.logger.Infoln("Parsed playlistId from Vars")

		if hasAccess, err := m.playlistUseCase.HasModifyAccess(userId, uint64(playlistId)); err != nil || !hasAccess {
			common_handler.Handler{H: default_handlers.Forbidden}.ServeHTTP(w, request)
			return
		}

		next.ServeHTTP(w, request)
	})
}
