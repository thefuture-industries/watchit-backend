package movie

import (
	"net/http"
	"sort"
	"watchit/httpx/infra/store/postgres/models"
	"watchit/httpx/infra/types"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"
	"watchit/httpx/pkg/machinelearning"
)

var maxCountMovie int = 50

func (h *Handler) GetMoviesSuggestHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

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

	if err := machinelearning.ShuffleArray(moviesTop); err != nil {
		return httperr.InternalServerError(err.Error())
	}

	httpx.HttpResponse(w, r, http.StatusOK, moviesTop)
	return nil
}
