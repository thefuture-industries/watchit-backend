package youtube

import (
	"context"
	"flicksfi/internal/config"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Обьект с категориями и их индексами
var category_ids map[string]string = map[string]string{
	"Film":          "1",
	"Animation":     "1",
	"Autos":         "2",
	"Vehicles":      "2",
	"Music":         "10",
	"Pets":          "15",
	"Animals":       "15",
	"Sports":        "17",
	"Travel":        "19",
	"Events":        "19",
	"Gaming":        "20",
	"People":        "22",
	"Blogs":         "22",
	"Entertainment": "24",
	"Howto":         "26",
	"Style":         "26",
	"Education":     "27",
	"Science":       "28",
	"Technology":    "28",
	"Nonprofits":    "29",
	"Activism":      "29",
}

// -------------------------------
// Отправка запроса на YouTube API
// -------------------------------
func Request(body map[string]string, max int64) ([]*youtube.SearchResult, error) {
	// Создание контекста
	ctx := context.Background()
	// Переменная для записи отпарвки данных на API
	var call *youtube.SearchListCall

	// Создания клиента
	client, err := youtube.NewService(ctx, option.WithAPIKey(config.Envs.YOUTUBE_KEY_API))
	if err != nil {
		return nil, err
	}

	// Получение категории
	category_id, ok := category_ids[body["category"]] // category

	if !ok {
		// Если не найдена отправляем категорию в поиск
		call = client.Search.List([]string{"snippet"}).
			Q(body["category"] + body["search"] + body["channel"]).
			MaxResults(max).
			Type("video")
	} else {
		// Если найдена отправляем категорию в параметр
		call = client.Search.List([]string{"snippet"}).
			Q(body["search"] + body["channel"]).
			VideoCategoryId(category_id).
			MaxResults(max).
			Type("video")
	}

	// Отправка запроса на API
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	// Возвращяем данные
	return response.Items, nil
}

// -----------------------------------
// Получение популярных видео с YouTube
// -----------------------------------
func GetPopular() ([]*youtube.SearchResult, error) {
	// Создание контекста
	ctx := context.Background()
	// Переменная для записи отпарвки данных на API
	var call *youtube.SearchListCall

	// Создания клиента
	client, err := youtube.NewService(ctx, option.WithAPIKey(config.Envs.YOUTUBE_KEY_API))
	if err != nil {
		return nil, err
	}

	call = client.Search.List([]string{"snippet"}).
		MaxResults(10).
		Type("video").
		PublishedAfter("2024-01-01T00:00:00Z").
		Order("viewCount").
		RegionCode("US")

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	// Возвращяем данные
	return response.Items, nil
}
