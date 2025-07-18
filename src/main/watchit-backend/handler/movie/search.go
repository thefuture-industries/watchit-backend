package movie

import (
	"net/http"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"
)

func (h *Handler) GetMoviesBySearchHandler(w http.ResponseWriter, r *http.Request) error {
	var payload *SearchPayload

	if err := httpx.HttpParse(r, &payload); err != nil {
		return httperr.BadRequest(err.Error())
	}

	if err := httpx.Validate.Struct(payload); err != nil {
		return httperr.BadRequest("not all fields are filled in")
	}

	sims :=

		httpx.HttpResponse(w, r, http.StatusOK, "MOVIEs")
	return nil
}
