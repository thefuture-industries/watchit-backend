package types

type Limiter struct {
	ID           int    `json:"id"`
	UUID         string `json:"uuid"`
	TextLimiter  int    `json:"text_limit"`
	YoutubeLimit int    `json:"youtube_limit"`
	UpdateAt     string `json:"update_at"`
}

type Favourites struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid"`
	MovieID     int    `json:"movieId"`
	MoviePoster string `json:"moviePoster"`
	CreatedAt   string `json:"createdAt"`
}

type Recommendations struct {
	ID    int    `json:"id"`
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

type JsonMovies struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

// Movie: Модель для JsonMovies данные для Result json
type Movie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}
