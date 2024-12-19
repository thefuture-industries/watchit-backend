package user

import (
	"context"
	"database/sql"
	"flicksfi/cmd/configuration"
	"flicksfi/internal/types"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	db      *sql.DB
	logger  *zap.Logger
	monitor *configuration.Track
}

func NewService(db *sql.DB, logger *zap.Logger, monitor *configuration.Track) *Service {
	return &Service{
		db:      db,
		logger:  logger,
		monitor: monitor,
	}
}

// Получения всех пльзователей
// ---------------------------
func (s *Service) GetUsers() ([]types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, "select * from users")
	if err != nil {
		// Логирование ошибки
		s.logger.Error("database error",
			zap.Error(err))
		return nil, fmt.Errorf("error execute query to db")
	}

	var users []types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.UUID, &user.UserName, &user.UserNameUpper, &user.Email, &user.EmailUpper, &user.IPAddress, &user.Lat, &user.Lon, &user.Country, &user.RegionName, &user.Zip, &user.CreatedAt)
		if err != nil {
			// Логирование ошибки
			s.logger.Error("scan user",
				zap.Error(err))
			return nil, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return []types.User{}, nil
	}

	return users, nil
}

// Получения пользователя по ID
// ----------------------------
func (s *Service) GetUserById(id int) (*types.User, error) {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по ID
	queryStart := time.Now()
	row := s.db.QueryRowContext(ctx, "select * from users where id = ? limit 1", id)

	// читаем из результата
	u := new(types.User)
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Lat, &u.Lon, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("scan user",
			zap.Int("userID", id),
			zap.Error(err))

		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	// Отправка пользователя
	return u, nil
}

// Получения пользователя по секретному слову
// ------------------------------------------
func (s *Service) GetUserBySecretWord(word string) (*types.User, error) {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по ID
	queryStart := time.Now()
	row := s.db.QueryRowContext(ctx, "select * from users where secret_word = ? limit 1", word)

	// читаем из результата
	u := new(types.User)
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Lat, &u.Lon, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("scan user",
			zap.String("secret_word", word),
			zap.Error(err))

		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	// Отправка пользователя
	return u, nil
}

// Получения пользователя по uuid
// ------------------------------
func (s *Service) GetUserByUUID(uuid string) (*types.User, error) {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по ID
	queryStart := time.Now()
	row := s.db.QueryRowContext(ctx, "select * from users where uuid = ?", uuid)

	// читаем из результата
	u := new(types.User)
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Lat, &u.Lon, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("scan user",
			zap.String("uuid", uuid),
			zap.Error(err))

		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	// Отправка пользователя
	return u, nil
}

// Получения пользователя по Email
// -------------------------------
func (s *Service) GetUserByEmail(email string) (*types.User, error) {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по Email
	queryStart := time.Now()
	row := s.db.QueryRowContext(ctx, "select * from users where email = ?", email)

	// читаем из результата
	u := new(types.User)
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Lat, &u.Lon, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("scan user",
			zap.String("email", email),
			zap.Error(err))

		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	// Отправка пользователя
	return u, nil
}

// Создание пользователя
// ---------------------
func (s *Service) CreateUser(user types.User) error {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9000) + 1000
	username := fmt.Sprintf("Гость%d", randomNumber)

	// Запрос к БД на создания пользователя
	queryStart := time.Now()
	_, err := s.db.ExecContext(ctx, "insert into users (uuid, secret_word, username, username_upper, ip_address, latitude, longitude, country, regionName, zip, createdAt) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", user.UUID, user.SecretWord, username, strings.ToUpper(username), user.IPAddress, user.Lat, user.Lon, user.Country, user.RegionName, user.Zip, user.CreatedAt)
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", user.UUID),
			zap.Error(err))

		return fmt.Errorf("iterate user: %w", err)
	}

	// Запрос к БД на создания лимитов
	_, err = s.db.ExecContext(ctx, "insert into limiter (uuid, update_at) values (?, ?)", user.UUID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", user.UUID),
			zap.Error(err))

		return fmt.Errorf("iterate limiter: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

	return nil
}

// Проверка на существования пользователя в БД
// -------------------------------------------
func (s *Service) CheckUser(user types.LoginUserPayload) (*types.User, error) {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по ID
	row := s.db.QueryRowContext(ctx, "select * from users where uuid = ?", user.UUID)
	u := new(types.User)

	// читаем из результата
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Lat, &u.Lon, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}

		// Логирование ошибки
		s.logger.Error("scan user",
			zap.String("uuid", user.UUID),
			zap.Error(err))

		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return u, nil
}

// func (s *Service) UserUpdate(user types.UserUpdate, uuid string) error {
// 	// определение контекста времени
// 	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
// 	defer cancel()

// 	query := "update users set "
// 	args := make(map[string]interface{})
// 	first := true

// 	if *user.Username != "" {
// 		if !first {
// 			query += ", "
// 		}

// 		query += "username = :username"
// 		args["username"] = user.Username
// 		first = false
// 	}

// 	if *user.Email != "" {
// 		if !first {
// 			query += ", "
// 		}

// 		query += "email = :email"
// 		args["email"] = user.Email
// 		first = false
// 	}

// 	if !first {
// 		query += " where uuid = :uuid"
// 		args["uuid"] = uuid
// 	} else {
// 		return fmt.Errorf("no fields provided for update")
// 	}

// 	result, err := s.db.NamedRe(ctx, "update users set username = ?, username_apper = ?, email = ?, email_upper = ?, secret_word = ?")

// 	return nil
// }
