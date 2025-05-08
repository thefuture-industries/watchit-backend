package interfaces

import "flicksfi/internal/types"

type IUser interface {
	// Получения всех пльзователей
	GetUsers() ([]types.User, error)

	// Получения пользователя по секретному слову
	GetUserBySecretWord(word string) (*types.User, error)

	// Получения пользователя по GetUserByIPAddress
	GetUserByUUID(uuid string) (*types.User, error)

	// Получения пользователя по ID
	GetUserById(id int) (*types.User, error)

	// Получения пользователя по IP
	GetUserByIP(ip string) (*types.User, error)

	// Получения пользователя по Email
	GetUserByEmail(email string) (*types.User, error)

	// Создание пользователя
	CreateUser(user types.User) error

	// Проверка на существования пользователя в БД
	CheckUser(user types.LoginUserPayload) (*types.User, error)

	// Обновление данных пользователя
	UserUpdate(user types.UserUpdate) error
}
