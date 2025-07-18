package favourite

import (
	"net/http"
	"watchit/httpx/pkg/httpx"
)

func (h *Handler) GetFavouritesHandler(w http.ResponseWriter, r *http.Request) error {
	httpx.HttpResponse(w, r, http.StatusOK, "FAVs")
	return nil
}
