package track_delivery

import (
	"encoding/json"
	"main/internal/pkg/album"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	"net/http"
)

type TrackHandler struct {
	trackUseCase   track.UseCase
	albumUseCase   album.UseCase
	sessionUseCase session.UseCase
}

func NewHandler(track track.UseCase, session session.UseCase) TrackHandler {
	return TrackHandler{
		trackUseCase:   track,
		sessionUseCase: session,
	}
}

// Listen
//
//	@Description	listen track
//	@Tags			track
//	@Accept			json
//	@Param			trackId	body	integer	true	"track id"
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Success		200
//	@Failure		400	{string}	errMsg
//	@Failure		403	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/listen [post]
func (handler *TrackHandler) Listen(w http.ResponseWriter, r *http.Request) error {
	var trackId track.Id
	if err := json.NewDecoder(r.Body).Decode(&trackId); err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	err := handler.trackUseCase.Listen(trackId.Id)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	return nil
}

// Like
//
//	@Description	like track
//	@Tags			track
//	@Security		csrfToken
//	@Security		cookieCsrfToken
//	@Accept			json
//	@Param			trackId	body		integer	true	"track id"
//	@Success		200		{object}	track.Response
//	@Failure		400		{string}	errMsg
//	@Failure		401		{string}	errMsg
//	@Failure		403		{string}	errMsg
//	@Failure		500		{string}	errMsg
//	@Router			/like [post]
func (handler *TrackHandler) Like(w http.ResponseWriter, r *http.Request) error {
	var trackId track.Id
	if err := json.NewDecoder(r.Body).Decode(&trackId); err != nil {
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

	err = handler.trackUseCase.Like(userId, trackId.Id)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	return nil
}
