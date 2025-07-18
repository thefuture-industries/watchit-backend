package favourite

import (
	"net/http"
	"watchit/httpx/infra/store/postgres/models"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"
)

func (h *Handler) AddFavouriteHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	authToken := r.Context().Value("identity").(string)

	var payload *FavouriteAddPayload

	if err := httpx.HttpParse(r, &payload); err != nil {
		return httperr.BadRequest(err.Error())
	}

	if err := httpx.Validate.Struct(payload); err != nil {
		return httperr.BadRequest("not all fields are filled in")
	}

	exists, err := h.Store.Favourites.Get_FavouriteByUuidByMovieId(ctx, authToken, payload.MovieId)
	if err != nil {
		return httperr.Db(ctx, err)
	}

	if exists != nil {
		return httperr.BadRequest("the movie is already in favorites")
	}

	if err := h.Store.Favourites.Create_Favourite(ctx, &models.Favourite{
		UserUUID:    authToken,
		MovieId:     payload.MovieId,
		MoviePoster: payload.MoviePoster,
	}); err != nil {
		return httperr.Db(ctx, err)
	}

	httpx.HttpResponse(w, r, http.StatusOK, "the movie is saved to favorites")
	return nil
}
