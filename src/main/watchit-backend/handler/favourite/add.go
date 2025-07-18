package favourite

import (
	"net/http"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"
)

func (h *Handler) AddFavouriteHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var payload *FavouriteAddPayload

	if err := httpx.HttpParse(r, &payload); err != nil {
		return httperr.BadRequest(err.Error())
	}

	if err := httpx.Validate.Struct(payload); err != nil {
		return httperr.BadRequest("not all fields are filled in")
	}

	exists, err := h.Store.Favourites.Get_FavouriteByUuidByMovieId(ctx)

	httpx.HttpResponse(w, r, http.StatusOK, "FAVs")
	return nil
}
