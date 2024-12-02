package errors

import (
	"flick_finder/internal/types"
	"fmt"
)

// Если фильм не найден по Title, то ошибка
func SearchNotFound(parametrs map[string]string, checked bool) error {
	if parametrs["search"] != "" && !checked {
		return fmt.Errorf("search film %s not found", parametrs["search"])
	}

	return nil
}

// Если фильм не найден по ReleaseDate, то ошибка
func DateNotFound(parametrs map[string]string, checked bool) error {
	if parametrs["release_date"] != "" && !checked {
		return fmt.Errorf("release date is %s not found", parametrs["release_date"])
	}

	return nil
}

// Ошибка страницы
func ErrorPage(page int, response []types.Movie) error {
	if page > len(response) {
		return fmt.Errorf("page %d not found", (page / 10))
	}

	return nil
}
