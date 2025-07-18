package favourite

type FavouriteAddPayload struct {
	MovieId     int    `json:"movie_id"`
	MoviePoster string `json:"movie_poster"`
}
