package movie

import (
	"net/http"
	"watchit/httpx/pkg/httpx"
)

func (h *Handler) GetMoviesBySearchHandler(w http.ResponseWriter, r *http.Request) error {
	var payload

	httpx.HttpResponse(w, r, http.StatusOK, "MOVIEs")
	return nil
}
