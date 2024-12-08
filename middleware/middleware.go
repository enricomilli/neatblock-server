package middleware

import (
	"log/slog"
	"net/http"

	apiutil "github.com/enricomilli/neat-server/api/api-utils"
)

func Make(h func(w http.ResponseWriter, r *http.Request) error) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := h(w, r); err != nil {
				apiutil.ResponseWithError(w, http.StatusBadRequest, err)
				slog.Error("Internal server error", "Error", err, "Path", r.URL.Path)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
