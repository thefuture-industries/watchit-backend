package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"watchit/httpx/infra/logger"
	"watchit/httpx/infra/store/postgres/models"
)

type UserStore struct {
	db     *sql.DB
	logger *logger.Logger
}

func (s *UserStore) Create_UserCore(ctx context.Context, tx *sql.Tx, user *models.UserCore) error {
	query := `
		INSERT INTO user_cores (user_uuid) VALUES ($1)
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, user.UserUUID)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) Create_UserProfile(ctx context.Context, tx *sql.Tx, user *models.UserProfile) error {
	query := `
		INSERT INTO user_profiles (user_uuid, username) VALUES ($1, $2)
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, user.UserUUID, user.Username)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) Create_UserLimit(ctx context.Context, tx *sql.Tx, user *models.UserLimit) error {
	query := `
		INSERT INTO user_limits (user_uuid, limit_id) VALUES ($1, $2)
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, user.UserUUID, user.LimitId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) Get_UserProfileByUsername(ctx context.Context, username string) (*models.UserProfile, error) {
	user := models.UserProfile{}

	query := `
		SELECT id, user_uuid, avatar, username, bio FROM user_profiles WHERE username = $1 LIMIT 1
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := s.db.QueryRowContext(ctx, query, username)

	if err := row.Scan(&user.ID, &user.UserUUID, &user.Avatar, &user.Username, &user.Bio); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (s *UserStore) Get_UserLimitByUuid(ctx context.Context, uuid string) (*models.UserLimit, error) {
	user := models.UserLimit{}

	query := `
		SELECT id, user_uuid, limit_id, daily_search_limit_usage
		FROM user_limits WHERE user_uuid = $1 LIMIT 1
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	row := s.db.QueryRowContext(ctx, query, uuid)

	if err := row.Scan(&user.ID, &user.UserUUID, &user.LimitId, &user.DailySearchLimitUsage); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (s *UserStore) Update_UserLimitIncrementUsageByUuid(ctx context.Context, uuid string) error {
	query := `
		SELECT increment_user_limits_usage($1)
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, uuid)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) Update_UserLimitReset(ctx context.Context) error {
	query := `
		UPDATE user_limits SET daily_search_limit_usage = 0
	`

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	fmt.Println("All limits resets!)")
	return nil
}
