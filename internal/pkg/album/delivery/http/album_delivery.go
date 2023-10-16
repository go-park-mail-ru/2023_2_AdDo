package album_delivery

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"main/internal/pkg/album"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	"main/internal/pkg/session"
	"main/internal/pkg/track"
	"net/http"
	"strconv"
)

type AlbumHandler struct {
	trackUseCase   track.UseCase
	albumUseCase   album.UseCase
	sessionUseCase session.UseCase
}

func NewHandler(trackUseCase track.UseCase, albumUseCase album.UseCase, session session.UseCase) AlbumHandler {
	return AlbumHandler{
		trackUseCase:   trackUseCase,
		albumUseCase:   albumUseCase,
		sessionUseCase: session,
	}
}

// Feed
//
//	@Description	return all albums
//	@Tags			album
//	@Produce		json
//	@Success		200
//	@Failure		500	{string}	errMsg
//	@Router			/feed [get]
func (handler *AlbumHandler) Feed(w http.ResponseWriter, r *http.Request) error {
	albums, err := handler.albumUseCase.GetRandom()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return handler.handleQuery(albums, w, r)
}

// New
//
//	@Description	return new albums
//	@Tags			album
//	@Produce		json
//	@Success		200
//	@Failure		500	{string}	errMsg
//	@Router			/new [get]
func (handler *AlbumHandler) New(w http.ResponseWriter, r *http.Request) error {
	albums, err := handler.albumUseCase.GetNew()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return handler.handleQuery(albums, w, r)
}

// MostLiked
//
//	@Description	return albums sorted by likes count
//	@Tags			album
//	@Produce		json
//	@Success		200
//	@Failure		500	{string}	errMsg
//	@Router			/most_liked [get]
func (handler *AlbumHandler) MostLiked(w http.ResponseWriter, r *http.Request) error {
	albums, err := handler.albumUseCase.GetMostLiked()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return handler.handleQuery(albums, w, r)
}

// Popular
//
//	@Description	return albums sorted by listen count
//	@Tags			album
//	@Produce		json
//	@Success		200
//	@Failure		500	{string}	errMsg
//	@Router			/popular [get]
func (handler *AlbumHandler) Popular(w http.ResponseWriter, r *http.Request) error {
	albums, err := handler.albumUseCase.GetPopular()
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return handler.handleQuery(albums, w, r)
}

// AlbumTracks
//
//	@Description	return all tracks from album
//	@Tags			album
//	@Produce		json
//	@Param			id	path	integer	true	"album id"
//	@Success		200
//	@Failure		400	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Header			200	{string}	X-CSRFTOKEN	"csrf token"
//	@Router			/album/{id} [get]
func (handler *AlbumHandler) AlbumTracks(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(session.XCsrfToken, csrf.Token(r))
	albumId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	result, err := handler.albumUseCase.GetAlbum(uint64(albumId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	err = response.RenderJSON(w, result)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return nil
}

func (handler *AlbumHandler) handleQuery(albums []album.Response, w http.ResponseWriter, r *http.Request) error {
	err := response.RenderJSON(w, albums)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	return nil
}
