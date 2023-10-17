package artist_delivery

import (
	"main/internal/pkg/artist"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	"main/internal/pkg/session"
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

type ArtistHandler struct {
	ArtistUseCase artist.UseCase
}

func NewHandler(artistUseCase artist.UseCase) ArtistHandler {
	return ArtistHandler{
		ArtistUseCase: artistUseCase,
	}
}

// ArtistInfo
//
//	@Description	return artist info
//	@Tags			artist
//	@Produce		json
//	@Param			id	path		integer	true	"artist id"
//	@Success		200	{object}	artist.Response
//	@Failure		400	{string}	errMsg
//	@Failure		500	{string}	errMsg
//	@Router			/artist/{id} [get]
func (handler *ArtistHandler) ArtistInfo(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(session.XCsrfToken, csrf.Token(r))

	artistId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return common_handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	artistInfo, err := handler.ArtistUseCase.GetArtistInfo(uint64(artistId))
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	err = response.RenderJSON(w, artistInfo)
	if err != nil {
		return common_handler.StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	return nil
}
