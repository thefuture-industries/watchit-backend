package user

import (
	"context"
	"database/sql"
	"flick_finder/internal/types"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// Пробегаемся по данным по rows
func scanRowDataUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.UUID,
		&user.SecretWord,
		&user.UserName,
		&user.UserNameUpper,
		&user.Email,
		&user.EmailUpper,
		&user.IPAddress,
		&user.Lon,
		&user.Lat,
		&user.Country,
		&user.RegionName,
		&user.Zip,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error scaning data")
	}

	return user, nil
}

func scanRowUser(rows *sql.Rows, user *types.User) error {
	return rows.Scan(
		&user.ID,
		&user.UUID,
		&user.SecretWord,
		&user.UserName,
		&user.UserNameUpper,
		&user.Email,
		&user.EmailUpper,
		&user.IPAddress,
		&user.Lon,
		&user.Lat,
		&user.Country,
		&user.RegionName,
		&user.Zip,
		&user.CreatedAt,
	)
}

// Получения всех пльзователей
// ---------------------------
func (s *Service) GetUsers() ([]types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, "select * from users")
	if err != nil {
		return nil, fmt.Errorf("error execute query to db")
	}

	var users []types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.ID, &user.UUID, &user.UserName, &user.UserNameUpper, &user.Email, &user.EmailUpper, &user.IPAddress, &user.Lat, &user.Lon, &user.Country, &user.RegionName, &user.Zip, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("there are no users")
	}

	return users, nil
}

// Получения пользователя по ID
// ----------------------------
func (s *Service) GetUserById(id int) (*types.User, error) {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по ID
	rows, err := s.db.QueryContext(ctx, "select * from users where id = ?", id)

	// проверка ошибки
	if err != nil {
		return nil, fmt.Errorf("database request error")
	}

	// читаем из результата
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowDataUser(rows)
		if err != nil {
			return nil, err
		}
	}

	// если нет пользователя то ошибка
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	// Отправка пользователя
	return u, nil
}

// Получения пользователя по секретному слову
// ------------------------------------------
func (s *Service) GetUserBySecretWord(word string) (*types.User, error) {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по ID
	rows, err := s.db.QueryContext(ctx, "select * from users where secret_word = ?", word)

	// проверка ошибки
	if err != nil {
		return nil, fmt.Errorf("database request error")
	}

	var user *types.User
	if rows.Next() {
		user = new(types.User)
		if err := scanRowUser(rows, user); err != nil {
			return nil, fmt.Errorf("error scanning user data: %w", err)
		}
	} else if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	// Отправка пользователя
	return user, nil
}

// Получения пользователя по uuid
// ------------------------------
func (s *Service) GetUserByUUID(uuid string) (*types.User, error) {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по ID
	rows, err := s.db.QueryContext(ctx, "select * from users where uuid = ?", uuid)

	// проверка ошибки
	if err != nil {
		return nil, fmt.Errorf("database request error")
	}

	var user *types.User
	if rows.Next() {
		user = new(types.User)
		if err := scanRowUser(rows, user); err != nil {
			return nil, fmt.Errorf("error scanning user data: %w", err)
		}
	} else if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	// Отправка пользователя
	return user, nil
}

// Получения пользователя по Email
// -------------------------------
func (s *Service) GetUserByEmail(email string) (*types.User, error) {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по Email
	rows, err := s.db.QueryContext(ctx, "select * from users where email = ?", email)

	// проверка ошибки
	if err != nil {
		return nil, fmt.Errorf("database request error")
	}

	// читаем из результата
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowDataUser(rows)
		if err != nil {
			return nil, err
		}
	}

	// Проверка на существование
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	// Отправка пользователя
	return u, nil
}

// Создание пользователя
// ---------------------
func (s *Service) CreateUser(user types.User) error {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9000) + 1000
	username := fmt.Sprintf("Гость%d", randomNumber)

	// Запрос к БД на создания пользователя
	_, err := s.db.ExecContext(ctx, "insert into users (uuid, secret_word, username, username_upper, ip_address, latitude, longitude, country, regionName, zip, createdAt) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", user.UUID, user.SecretWord, username, strings.ToUpper(username), user.IPAddress, user.Lat, user.Lon, user.Country, user.RegionName, user.Zip, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("database insert error")
	}

	// Запрос к БД на создания лимитов
	_, err = s.db.ExecContext(ctx, "insert into limiter (uuid, update_at) values (?, ?)", user.UUID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return fmt.Errorf("database insert error")
	}

	return nil
}

// Проверка на существования пользователя в БД
// -------------------------------------------
func (s *Service) CheckUser(user types.LoginUserPayload) (*types.User, error) {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на вывод пользователя по ID
	rows, err := s.db.QueryContext(ctx, "select * from users where uuid = ?", user.UUID)

	// проверка ошибки
	if err != nil {
		return nil, fmt.Errorf("database request error")
	}

	// читаем из результата
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowDataUser(rows)
		if err != nil {
			return nil, err
		}
	}

	// если нет пользователя то ошибка
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

// Запись избранного в БД
// ----------------------
func (s *Service) AddFavourite(favourite types.FavouriteAddPayload, uuid string) error {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	// Запрос к БД на создания пользователя
	_, err := s.db.ExecContext(ctx, "insert into favourites (uuid, movieId, moviePoster, createdAt) values (?, ?, ?, ?)", uuid, favourite.MovieID, favourite.MoviePoster, time.Now().Format("2006-01-02 15:04:05"))

	// Обработка ошибки
	if err != nil {
		return fmt.Errorf("database insert error")
	}

	return nil
}

// ----------------------------------
// Получение списка избранных фильмов
// ----------------------------------
func (s *Service) Favourites(uuid string) ([]types.Favourites, error) {
	// определение контекста времени
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, "select * from favourites where uuid = ?", uuid)
	if err != nil {
		return nil, fmt.Errorf("error execute query to db")
	}

	var favourites []types.Favourites
	for rows.Next() {
		var f types.Favourites
		err := rows.Scan(&f.ID, &f.UUID, &f.MovieID, &f.MoviePoster, &f.CreatedAt)
		if err != nil {
			return nil, err
		}

		favourites = append(favourites, f)
	}

	if len(favourites) == 0 {
		return nil, fmt.Errorf("you don't have any favourites!")
	}

	return favourites, nil
}
