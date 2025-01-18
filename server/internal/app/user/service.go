package user

import (
	"context"
	"database/sql"
	"flicksfi/cmd/configuration"
	"flicksfi/internal/types"
	"flicksfi/internal/utils"
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
	defer rows.Close()

	var users []types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.UUID, &user.UserName, &user.UserNameUpper, &user.Email, &user.EmailUpper, &user.IPAddress, &user.Country, &user.RegionName, &user.Zip, &user.CreatedAt)
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
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)

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

// Получения пользователя по IP
// ----------------------------
func (s *Service) GetUserByIP(ip string) (*types.User, error) {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по ID
	queryStart := time.Now()
	row := s.db.QueryRowContext(ctx, "select * from users where ip_address = ? limit 1", ip)

	// читаем из результата
	u := new(types.User)
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("scan user",
			zap.String("ip", ip),
			zap.Error(err))

		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// мониторинг времени запроса
	s.monitor.TrackDBQuery(time.Since(queryStart))

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
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)

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
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)

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
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)

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

	// Начало транзакции
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("transaction begin error",
			zap.String("uuid", user.UUID),
			zap.Error(err))

		return fmt.Errorf("start transaction: %w", err)
	}

	// Функция отката при ошибке или панике
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			s.monitor.TrackDBError()
			s.monitor.TrackError(fmt.Errorf("panic during transaction: %v", r))
			s.logger.Error("transaction rolled back due to panic",
				zap.String("uuid", user.UUID),
				zap.Any("panic", r))
		}
	}()

	// Запрос к БД на создания пользователя
	queryStart := time.Now()
	_, err = tx.ExecContext(ctx, "insert into users (uuid, secret_word, username, username_upper, ip_address, country, regionName, zip, createdAt) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", user.UUID, user.SecretWord, username, strings.ToUpper(username), user.IPAddress, user.Country, user.RegionName, user.Zip, user.CreatedAt)
	if err != nil {
		tx.Rollback()
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", user.UUID),
			zap.Error(err))

		return fmt.Errorf("iterate user: %w", err)
	}

	// Запрос к БД на создания лимитов
	_, err = tx.ExecContext(ctx, "insert into limiter (uuid, update_at) values (?, ?)", user.UUID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		tx.Rollback()
		// Логирование ошибки
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error",
			zap.String("uuid", user.UUID),
			zap.Error(err))

		return fmt.Errorf("iterate limiter: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("transaction commit error",
			zap.String("uuid", user.UUID),
			zap.Error(err))
		return fmt.Errorf("commit transaction: %w", err)
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
	err := row.Scan(&u.ID, &u.UUID, &u.SecretWord, &u.UserName, &u.UserNameUpper, &u.Email, &u.EmailUpper, &u.IPAddress, &u.Country, &u.RegionName, &u.Zip, &u.CreatedAt)
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

// Обновление данных пользователя
// ------------------------------
func (s *Service) UserUpdate(user types.UserUpdate) error {
	start := time.Now()
	defer func() {
		s.monitor.TrackRequest(time.Since(start))
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	preparing_query := "update users set "
	params := []interface{}{}

	// Обновление username
	if user.Username != nil && *user.Username != "" {
		preparing_query += "username = ?, username_upper = ?, "
		params = append(params, *user.Username, strings.ToUpper(*user.Username))
	}

	// Обновление email
	if user.Email != nil && *user.Email != "" {
		preparing_query += "email = ?, email_upper = ?, "
		params = append(params, *user.Email, strings.ToUpper(*user.Email))
	}

	// Обновление secret_word
	if user.SecretWord != nil && *user.SecretWord != "" {
		// Шифруем секретное слово
		encrypt, err := utils.Encrypt(*user.SecretWordOld)
		if err != nil {
			s.logger.Error("encrypt secret word", zap.Error(err))
			s.monitor.TrackError(err)
			return fmt.Errorf("failed to encrypt secret word: %w", err)
		}

		// Проверяем существование пользователя с таким секретным словом
		u, err := s.GetUserBySecretWord(encrypt)
		if err != nil {
			s.logger.Error("Failed to get user by secret_word", zap.Error(err), zap.String("secret_word", encrypt))
			s.monitor.TrackError(err)
			return fmt.Errorf("failed to get user by secret_word: %w", err)
		}

		// Если пользователь не найден, то возвращаем ошибку
		if u == nil {
			s.logger.Error("secret word not found/incorrect", zap.String("secret_word", encrypt))
			s.monitor.TrackError(err)
			return fmt.Errorf("secret word not found/incorrect")
		}

		// Если все ок, то меняем секретное слово
		encrypted_secret_word, err := utils.Encrypt(*user.SecretWord)
		if err != nil {
			s.logger.Error("encrypt secret word", zap.Error(err))
			s.monitor.TrackError(err)
			return fmt.Errorf("failed to encrypt secret word: %w", err)
		}

		preparing_query += "secret_word = ?, "
		params = append(params, encrypted_secret_word)
	}

	// Удаляем последнюю запятую и пробел
	preparing_query = preparing_query[:len(preparing_query)-2]
	preparing_query += " where uuid = ?"
	params = append(params, user.UUID)

	queryStart := time.Now()
	_, err := s.db.ExecContext(ctx, preparing_query, params...)
	if err != nil {
		s.monitor.TrackDBError()
		s.monitor.TrackError(err)
		s.logger.Error("database error", zap.String("uuid", user.UUID), zap.Error(err))
		return fmt.Errorf("failed to update user: %w", err)
	}

	s.monitor.TrackDBQuery(time.Since(queryStart))
	return nil
}
