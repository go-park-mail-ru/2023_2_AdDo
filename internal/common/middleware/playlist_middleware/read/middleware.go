package read_playlist

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"main/internal/common/default_handlers"
	common_handler "main/internal/common/handler"
	"main/internal/pkg/playlist"
	"net/http"
	"strconv"
)

type Middleware struct {
	playlistUseCase playlist.UseCase
	logger          *logrus.Logger
}

func NewMiddleware(pu playlist.UseCase, logger *logrus.Logger) Middleware {
	return Middleware{
		playlistUseCase: pu,
		logger:          logger,
	}
}

func (m *Middleware) ReadPlaylistAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		playlistId, err := strconv.Atoi(mux.Vars(request)["id"])
		if err != nil {
			common_handler.Handler{H: default_handlers.BadRequest}.ServeHTTP(w, request)
			return
		}
		m.logger.Infoln("Parsed playlistId from Vars")

		if hasAccess, err := m.playlistUseCase.HasReadAccess(uint64(playlistId)); err != nil || !hasAccess {
			common_handler.Handler{H: default_handlers.Forbidden}.ServeHTTP(w, request)
			return
		}
		m.logger.Infoln("Ready to read")

		next.ServeHTTP(w, request)
	})
}
