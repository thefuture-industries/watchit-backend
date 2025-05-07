package recommendation

import (
	"go-movie-service/internal/common/database/action"
	"go-movie-service/internal/common/database/schema"
	"go-movie-service/internal/common/utils"
	"go-movie-service/internal/packages"
	"go-movie-service/internal/types"
	"net/http"

	"github.com/noneandundefined/vision-go"
)

type Handler struct {
	monitor *vision.Vision
	errors  *packages.Errors
}

func NewHandler(monitor *vision.Vision, errors *packages.Errors) *Handler {
	return &Handler{
		monitor: monitor,
		errors:  errors,
	}
}

func (h Handler) RecommendationGetHandler(w http.ResponseWriter, r *http.Request) {}

func (h Handler) RecommendationAddHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("identity").(*schema.Users)

	var payload *types.RecommendationAddPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, err.Error())
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, "Not all fields are filled in")
		return
	}

	for _, genreID := range payload.Genres {
		updated, err := action.UpdateRecommendation(user.UUID, genreID)
		if err != nil {
			utils.WriteJSON(w, r, http.StatusBadRequest, err.Error())
			return
		}

		if !updated {
			if err := action.CreateRecommendation(user.UUID, genreID); err != nil {
				utils.WriteJSON(w, r, http.StatusBadRequest, err.Error())
				return
			}
		}
	}

	utils.WriteJSON(w, r, http.StatusNoContent, "")
}
