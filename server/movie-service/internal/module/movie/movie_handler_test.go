package movie_test

import (
	"encoding/json"
	"go-movie-service/internal/module/movie"
	"go-movie-service/internal/packages"
	"go-movie-service/internal/types"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/noneandundefined/vision-go"
)

// TestMovieDetailsHandler тестирует обработчик получения деталей фильма
func TestMovieDetailsHandler(t *testing.T) {
	monitoring := vision.NewVision()
	errors := packages.NewErrors(monitoring)
	handler := movie.NewHandler(monitoring, errors)

	tests := []struct {
		name           string
		movieID        string
		setupMockData  func()
		wantStatusCode int
		wantTitle      string
	}{
		{
			name:    "Valid movie ID 1034541",
			movieID: "1034541",
			setupMockData: func() {
				setupTestMovieData(t, 1034541, "Terrifier 3")
			},
			wantStatusCode: http.StatusOK,
			wantTitle:      "Terrifier 3",
		},
		{
			name:    "Invalid movie ID format",
			movieID: "invalid",
			setupMockData: func() {
			},
			wantStatusCode: http.StatusBadRequest,
			wantTitle:      "",
		},
		{
			name:    "Movie not found",
			movieID: "999999",
			setupMockData: func() {
				setupEmptyTestData(t)
			},
			wantStatusCode: http.StatusNotFound,
			wantTitle:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupMockData != nil {
				tt.setupMockData()
			}

			req, err := http.NewRequest("GET", "/movie/"+tt.movieID, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/movie/{id}", handler.MovieDetailsHandler).Methods("GET")

			router.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.wantStatusCode)
			}

			if rr.Code == http.StatusOK {
				cacheControl := rr.Header().Get("Cache-Control")
				if cacheControl != "public, max-age=3600" {
					t.Errorf("handler returned wrong Cache-Control header: got %v want %v",
						cacheControl, "public, max-age=3600")
				}

				expires := rr.Header().Get("Expires")
				if expires == "" {
					t.Errorf("handler did not set Expires header")
				}

				var movie types.Movie
				err = json.NewDecoder(rr.Body).Decode(&movie)
				if err != nil {
					t.Errorf("error decoding response body: %v", err)
				}

				if movie.Title != tt.wantTitle {
					t.Errorf("handler returned wrong movie title: got %v want %v",
						movie.Title, tt.wantTitle)
				}
			}
		})
	}
}

func setupTestMovieData(t *testing.T, id int, title string) {
	testData := types.JsonMovies{
		Page: 1,
		Results: []types.Movie{
			{
				Id:               id,
				Title:            title,
				OriginalTitle:    title,
				OriginalLanguage: "en",
				Overview:         "Test overview",
				ReleaseDate:      "2024-10-09",
				VoteAverage:      7.201,
			},
		},
		TotalPages:   1,
		TotalResults: 1,
	}

	createTestDataFile(t, testData)
}

func setupEmptyTestData(t *testing.T) {
	testData := types.JsonMovies{
		Page:         1,
		Results:      []types.Movie{},
		TotalPages:   0,
		TotalResults: 0,
	}

	createTestDataFile(t, testData)
}

func createTestDataFile(t *testing.T, testData types.JsonMovies) {
	setupTestDataWithContent(t, testData)
}
