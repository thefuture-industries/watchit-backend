package interfaces

import "flick_finder/internal/types"

type IUser interface {
	// Получения всех пльзователей
	GetUsers() ([]types.User, error)

	// Получения пользователя по секретному слову
	GetUserBySecretWord(word string) (*types.User, error)

	// Получения пользователя по GetUserByIPAddress
	GetUserByUUID(uuid string) (*types.User, error)

	// Получения пользователя по ID
	GetUserById(id int) (*types.User, error)

	// Получения пользователя по Email
	GetUserByEmail(email string) (*types.User, error)

	// Создание пользователя
	CreateUser(user types.User) error

	// Проверка на существования пользователя в БД
	CheckUser(user types.LoginUserPayload) (*types.User, error)

	// Запись избранного в БД
	AddFavourite(favourite types.FavouriteAddPayload, uuid string) error

	// Получение списка избранных фильмов
	Favourites(uuid string) ([]types.Favourites, error)
}
