package favourite

type FavouriteAddPayload struct {
	MovieId     int    `json:"movie_id" validate:"required"`
	MoviePoster string `json:"movie_poster" validate:"required"`
}
