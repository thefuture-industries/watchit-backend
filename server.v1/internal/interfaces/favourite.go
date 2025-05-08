package interfaces

import "flicksfi/internal/types"

type IFavourite interface {
	// Проверка что избранное не существует в БД у uuid
	CheckFavourites(uuid string, movieID int) error

	// Запись избранного в БД
	AddFavourite(favourite types.FavouriteAddPayload) error

	// Удаление избранного фильма
	DeleteFavourite(payload types.FavouriteDeletePayload) error

	// Получение списка избранных фильмов
	Favourites(uuid string) ([]types.Favourites, error)
}
