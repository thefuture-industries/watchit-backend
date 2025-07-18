package store

import (
	"context"
	"database/sql"
	"watchit/httpx/infra/logger"
	"watchit/httpx/infra/store/postgres/models"
)

type Storage struct {
	Users interface {
		Create_UserCore(ctx context.Context, tx *sql.Tx, user *models.UserCore) error
		Create_UserProfile(ctx context.Context, tx *sql.Tx, user *models.UserProfile) error
		Create_UserLimit(ctx context.Context, tx *sql.Tx, user *models.UserLimit) error

		Get_UserProfileByUsername(ctx context.Context, username string) (*models.UserProfile, error)
		Get_UserLimitByUuid(ctx context.Context, uuid string) (*models.UserLimit, error)
	}
	Movies interface {
		Get_Movies(ctx context.Context) (*[]models.Movie, error)
		Get_MovieById(ctx context.Context, id int) (*models.MovieWithGenres, error)
	}
}

func NewStorage(db *sql.DB, logger *logger.Logger) Storage {
	return Storage{
		Users:  &UserStore{db: db, logger: logger},
		Movies: &MovieStore{db: db, logger: logger},
	}
}

func WithTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
