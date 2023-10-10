package artist_delivery

import (
	"github.com/gorilla/csrf"
	"main/internal/pkg/artist"
	common_handler "main/internal/pkg/common/handler"
	"main/internal/pkg/common/response"
	"main/internal/pkg/session"
	"net/http"
	"strconv"
)

type ArtistHandler struct {
	ArtistUseCase artist.UseCase
}

func NewHandler(artistUseCase artist.UseCase) ArtistHandler {
	return ArtistHandler{
		ArtistUseCase: artistUseCase,
	}
}

func (handler *ArtistHandler) ArtistInfo(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(session.XCsrfToken, csrf.Token(r))

	artistId, err := strconv.Atoi(r.URL.Query().Get("id"))
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
