package track_delivery

import (
	"github.com/gorilla/csrf"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	"net/http"
)

type TrackHandler struct {
	trackUseCase   track.UseCase
	sessionUseCase session.UseCase
}

func NewHandler(track track.UseCase, session session.UseCase) TrackHandler {
	return TrackHandler{
		trackUseCase:   track,
		sessionUseCase: session,
	}
}

//{
//	"Id":1,
//	"Name":"Stargazing",
//	"Artist":[{
//		"Id":1,
//		"Name":"Travis Scott",
//		"Avatar":"http://82.146.45.164:9000/images/avatars/artists/Travis_Scott.jpg"
//	}],
//	"Album":[{
//		"Id":1,
//		"Name":"Astroworld",
//		"Preview":"http://82.146.45.164:9000/audio/Travis_Scott/albums/Astroworld"
//	}],
//	"Preview":"http://82.146.45.164:9000/images/tracks/Travis_Scott/albums/Astroworld.jpg",
//	"Content":"http://82.146.45.164:9000/audio/Travis_Scott/albums/Astroworld/Stargazing.mp3"
//}

// @Description	return all tracks
// @Tags			track
// @Produce		json
// @Param			id	query	int	false	"user id"
// @Security		cookieAuth
// @Security		csrfToken
// @Success		200	{array}		track.Response
// @Failure		400	{string}	errMsg
// @Failure		401	{string}	errMsg
// @Failure		500	{string}	errMsg
// @Router			/music [get]

func (handler *TrackHandler) Music(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(session.XCsrfToken, csrf.Token(r))
	tracks, err := handler.trackUseCase.GetAll()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	err = response.RenderJSON(w, tracks)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	// TODO На текущий момент мы не возвращаем пользовательскую музыку, потому что нет смысла возвращать то, что у пользователя в добавленных
	// TODO Поэтому к первому рк можно сделать так, что мы возвращаем просто общую музыку
	// TODO После того как придумаем простейшую систему рекомендаций, будет смысл возвращать что-то из пользовательских подборок
	return nil
}
