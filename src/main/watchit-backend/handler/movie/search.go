package movie

import (
	"net/http"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"
)

func (h *Handler) GetMoviesBySearchHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var payload *SearchPayload

	if err := httpx.HttpParse(r, &payload); err != nil {
		return httperr.BadRequest(err.Error())
	}

	if err := httpx.Validate.Struct(payload); err != nil {
		return httperr.BadRequest("not all fields are filled in")
	}

	movies, err := h.Store.Movies.Get_Movies(ctx)
	if err != nil {
		return httperr.Db(ctx, err)
	}

	if movies == nil {
		return httperr.NotFound("sorry, we couldn't find the movie")
	}

	sims := lsaBuilder.AnalyzeByCosine(movies, payload.Text)

	httpx.HttpResponse(w, r, http.StatusOK, "MOVIEs")
	return nil
}
