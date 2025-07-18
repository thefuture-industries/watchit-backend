package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"time"
	"watchit/httpx/infra/logger"
	"watchit/httpx/infra/store/postgres/models"
)

type MovieStore struct {
	db     *sql.DB
	logger *logger.Logger
}

func (s *MovieStore) Get_Movies(ctx context.Context) (*[]models.Movie, error) {
	movies := []models.Movie{}

	query := `
		SELECT id, title, overview, release_date, original_language, popularity, vote_average, poster_path, backdrop_path, video, adult FROM movie
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		movie := models.Movie{}

		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Overview,
			&movie.ReleaseDate,
			&movie.OriginalLanguage,
			&movie.Popularity,
			&movie.VoteAverage,
			&movie.PosterPath,
			&movie.BackdropPath,
			&movie.Video,
			&movie.Adult,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &movies, nil
}

func (s *MovieStore) Get_MovieById(ctx context.Context, id int) (*models.MovieWithGenres, error) {
	movie := models.MovieWithGenres{}

	query := `
		SELECT
		    movie.id,
		    movie.title,
		    movie.overview,
		    movie.release_date,
		    movie.original_language,
		    movie.popularity,
		    movie.vote_average,
		    movie.poster_path,
		    movie.backdrop_path,
		    movie.video,
		    movie.adult,
		    array_agg(genres.genre_name ORDER BY genres.genre_name) AS genres
		FROM movie
		LEFT JOIN movie_genres ON movie.id = movie_genres.movie_id
		LEFT JOIN genres ON movie_genres.genre_id = genres.genre_id
		WHERE movie.id = $1 GROUP BY movie.id
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Overview,
		&movie.ReleaseDate,
		&movie.OriginalLanguage,
		&movie.Popularity,
		&movie.VoteAverage,
		&movie.PosterPath,
		&movie.BackdropPath,
		&movie.Video,
		&movie.Adult,
		pq.Array(&movie.Genres),
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		fmt.Println(err.Error())
		return nil, err
	}

	return &movie, nil
}
