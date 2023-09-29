package track_delivery

import (
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	"net/http"
	"strconv"
)

type TrackHandler struct {
	trackUseCase   track.UseCase
	sessionUseCase session.Usecase
}

func NewHandler(track track.UseCase, session session.Usecase) TrackHandler {
	return TrackHandler{
		trackUseCase:   track,
		sessionUseCase: session,
	}
}

func (handler *TrackHandler) Music(w http.ResponseWriter, r *http.Request) error {
	tracks, err := handler.trackUseCase.GetAll()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	err = response.RenderJSON(w, tracks)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	cookie, err := r.Cookie(session.CookieName)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	isAuth, err := handler.sessionUseCase.CheckSession(cookie.Value, uint64(userId))

	if err != nil || !isAuth {
		return common_handler.StatusError{Code: http.StatusUnauthorized, Err: err}
	}

	userTracks, err := handler.trackUseCase.GetFavourite(uint64(userId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	err = response.RenderJSON(w, userTracks)
	return err
}
