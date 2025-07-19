package movie

import (
	"net/http"
	"sort"
	"time"
	"watchit/httpx/infra/constants"
	"watchit/httpx/infra/store/postgres/models"
	"watchit/httpx/infra/types"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"
	"watchit/httpx/pkg/machinelearning"
)

func (h *Handler) GetMoviesSuggestHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	authToken := r.Context().Value("identity").(string)

	var payload *SuggestPayload

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

	if err := machinelearning.ShuffleArray(*movies); err != nil {
		return httperr.InternalServerError(err.Error())
	}

	// check user limit
	// ----------------
	limits, err := h.Store.Users.Get_UserLimitByUuid(ctx, authToken)
	if err != nil {
		return httperr.Db(ctx, err)
	}

	if limits == nil {
		return httperr.Forbidden("log in to your account")
	}

	switch limits.LimitId {
	case constants.LimitFree:
		if len(payload.Text) > constants.FreeMaxQueryLengthUsage {
			return httperr.BadRequest("request length exceeded in the free version")
		}

		if limits.DailySearchLimitUsage > constants.FreeDailySearchLimitUsage {
			return httperr.BadRequest("you have reached the daily search limit")
		}

		limitMovie := len(*movies) * 85 / 100
		*movies = (*movies)[:limitMovie]

		time.Sleep(2 * time.Second)

	case constants.LimitPay:
		if len(payload.Text) > constants.PayMaxQueryLengthUsage {
			return httperr.BadRequest("oops, request length exceeded")
		}

		if limits.DailySearchLimitUsage > constants.PayDailySearchLimitUsage {
			return httperr.BadRequest("you have reached the daily search limit")
		}
	}
	// ----------------
	// check user limit

	matrix, docs := lsaBuilder.AnalyzeByMovie(*movies, payload.Text)
	if matrix == nil {
		return httperr.NotFound("we didn't find any movies")
	}

	rows, _ := matrix.Dims()
	inputVec := matrix.RawRowView(rows - 1)

	sims := make([]types.LSASimilarity, 0, rows-1)
	for i := 0; i < rows-1; i++ {
		rowVec := matrix.RawRowView(i)
		sim := lsaBuilder.CosineSimilarity(rowVec, inputVec)
		sims = append(sims, types.LSASimilarity{Index: i, Similarity: sim})
	}

	sort.Slice(sims, func(i, j int) bool {
		return sims[i].Similarity > sims[j].Similarity
	})

	if len(sims) < maxCountMovie {
		maxCountMovie = len(sims)
	}

	var moviesTop []models.Movie

	for i := 0; i < maxCountMovie; i++ {
		idx := sims[i].Index
		movie := docs[idx]

		moviesTop = append(moviesTop, movie)
	}

	// db: update limits usage
	if err := h.Store.Users.Update_UserLimitIncrementUsageByUuid(ctx, limits.UserUUID); err != nil {
		return httperr.Db(ctx, err)
	}

	httpx.HttpResponse(w, r, http.StatusOK, moviesTop)
	return nil
}
