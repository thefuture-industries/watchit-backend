package store

import (
	"context"
	"database/sql"
	"errors"
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
		SELECT title, overview, release_date, original_language, popularity, vote_average, poster_path, backdrop_path, video, adult FROM movie
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

func (s *MovieStore) Get_MovieById(ctx context.Context, id int) (*models.Movie, error) {
	movie := &models.Movie{}

	query := `
		SELECT id, title, overview, release_date, original_language, popularity, vote_average, poster_path, backdrop_path, video, adult FROM movie
		WHERE id = $1 LIMIT 1
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&movie.ID,
	)
}
