package favourite

import (
	"net/http"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"
)

func (h *Handler) GetFavouritesHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	authToken := r.Context().Value("identity").(string)

	favourites, err := h.Store.Favourites.Get_FavouritesByUuid(ctx, authToken)
	if err != nil {
		return httperr.Db(ctx, err)
	}

	httpx.HttpResponse(w, r, http.StatusOK, favourites)
	return nil
}
