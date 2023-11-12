package csrf

import (
	"github.com/gorilla/csrf"
	"main/internal/pkg/session"
	"net/http"
)

func GetCSRF(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(session.XCsrfToken, csrf.Token(r))

	w.WriteHeader(http.StatusNoContent)
	return nil
}
