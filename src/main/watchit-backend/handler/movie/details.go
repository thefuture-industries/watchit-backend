package movie

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"
)

func (h *Handler) GetDetailsMovieHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	idParam := mux.Vars(r)["id"]
	if idParam == "" {
		return httperr.NotFound("movie with the id was not found")
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return httperr.BadRequest("movie id was not an integer")
	}

	movie, err := h.Store.Movies.Get_MovieById(ctx, id)
	if err != nil {
		return httperr.Db(ctx, err)
	}

	if movie == nil {
		return httperr.NotFound("movie not found")
	}

	httpx.HttpResponse(w, r, http.StatusOK, "MOVIEs")
	return nil
}
