package movie

import (
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"flicksfi/internal/config"
	"flicksfi/internal/types"
	"flicksfi/pkg"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

// ----------------------
// Поиск по сюжету фильма
// ----------------------
func OverviewText(text string) ([]types.Movie, error) {
	// Чтение файла
	file, err := os.Open("pkg/movie/db/movies.json.gz")
	if err != nil {
		return nil, fmt.Errorf("error reading movies")
	}
	defer file.Close()

	// Открытие zip файл
	zr, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("error reading movies")
	}
	defer zr.Close()

	// Чтение файла zip
	data, err := io.ReadAll(zr)
	if err != nil {
		return nil, fmt.Errorf("error reading movies")
	}

	// Конвертация данных файла в json
	var movies []types.JsonMovies
	var response []types.Movie

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, fmt.Errorf("error convert data to json")
	}

	for _, movie := range movies {
		for _, movieItem := range movie.Results {
			tfidf := pkg.TF_IDF_MOVIE(movieItem.Overview, text, 0.289)
			fmt.Printf("tfidf для %s: %.4f\n", movieItem.Title, tfidf)

			if tfidf >= 0.4 {
				response = append(response, movieItem)
			}
		}
	}

	return pkg.TruncateArrayMovies(response), nil
}

// -------------------------
// Поиск по сюжету GIGA CHAT
// -------------------------
func GIGA_CHAT_OVERVIEW(text string) ([]types.Movie, error) {
	// Получение токена доступа
	// ------------------------
	url := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	payload := strings.NewReader("scope=GIGACHAT_API_PERS")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ЭТО НЕБЕЗОПАСНО!!
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return []types.Movie{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("RqUID", uuid.NewString())
	bearer := fmt.Sprintf("Basic %s", config.Envs.GIGA_CHAT_AUTH_KEY)
	req.Header.Add("Authorization", bearer)

	resp, err := client.Do(req)
	if err != nil {
		return []types.Movie{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []types.Movie{}, err
	}

	var oAuth types.OAuthResponse
	err = json.Unmarshal(body, &oAuth)
	if err != nil {
		return []types.Movie{}, err
	}

	// Получение ответа на сообщение от Giga Chat
	text_request := fmt.Sprintf(`Output a list of movie titles and enclose it in an array of [], at least 40 movies (don't ask questions, just output only a list of the list) for example ["item1", "item2"] min. 40 movie, where the name of the movie or its plot is similar to this: %s`, text)
	url = "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
	jsonPayload := struct {
		Model    string `json:"model"`
		Messages []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
		Stream         bool `json:"stream"`
		UpdateInterval int  `json:"update_interval"`
	}{
		Model: "GigaChat",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "user", Content: text_request},
		},
		Stream:         false,
		UpdateInterval: 0,
	}

	// Маршалим структуру в JSON
	jsonString, err := json.Marshal(jsonPayload)
	if err != nil {
		return []types.Movie{}, err
	}

	payload = strings.NewReader(string(jsonString))

	client = &http.Client{Transport: tr}
	req, err = http.NewRequest("POST", url, payload)
	if err != nil {
		return []types.Movie{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	bearer = fmt.Sprintf("Bearer %s", oAuth.AccessToken)
	req.Header.Add("Authorization", bearer)

	resp, err = client.Do(req)
	if err != nil {
		return []types.Movie{}, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return []types.Movie{}, err
	}

	var gigaMovie types.ChatCompletionResponse
	var contentMessage []string
	err = json.Unmarshal(body, &gigaMovie)
	if err != nil {
		return []types.Movie{}, err
	}

	fmt.Println("Content " + gigaMovie.Choices[0].Message.Content)

	// Читаем файл (gzip)
	file, err := os.Open("pkg/movie/db/movies.json.gz")
	if err != nil {
		return nil, fmt.Errorf("error open file")
	}
	defer file.Close()

	// Создать декомпрессор gzip
	zr, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("error decompress file to movie")
	}
	defer zr.Close()

	// Читаем массив байтов
	data, err := io.ReadAll(zr)
	if err != nil {
		return nil, fmt.Errorf("error read data movies")
	}

	// Создаем переменную для фильмов
	var movies []types.JsonMovies
	var response []types.Movie

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return nil, fmt.Errorf("error decode movies %s", err)
	}

	err = json.Unmarshal([]byte(gigaMovie.Choices[0].Message.Content), &contentMessage)
	if err != nil {
		return []types.Movie{}, err
	}

	fmt.Println("Array is movie: " + contentMessage[0] + contentMessage[1] + contentMessage[2])

	for _, movie := range movies {
		for _, movieItem := range movie.Results {
			for _, movieTitleGiga := range contentMessage {
				if strings.Contains(movieItem.Title, movieTitleGiga) {
					response = append(response, movieItem)
					break
				}
			}
		}
	}

	return pkg.TruncateArrayMovies(response), nil
}
