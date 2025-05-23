package types

type RecommendationAddPayload struct {
	Genres []uint16 `json:"genres" validate:"required"`
}

type MoviePage struct {
	Page uint16 `json:"page"`
}

type Movies struct {
	Page         uint16  `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   uint32  `json:"total_pages"`
	TotalResults uint32  `json:"total_results"`
}

type Movie struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"`
	GenreIds         []uint16 `json:"genre_ids"`
	Id               uint32   `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title"`
	Overview         string   `json:"overview"`
	Popularity       float32  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	ReleaseDate      string   `json:"release_date"`
	Title            string   `json:"title"`
	Video            bool     `json:"video"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        uint16   `json:"vote_count"`
}

type IndexEntry struct {
	Page   int32
	Offset int64
}

type WC struct {
	Word  string
	Count int
}

type DocSimilarity struct {
	index      int
	similarity float64
}
