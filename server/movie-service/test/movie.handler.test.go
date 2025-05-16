package movie_test

import (
	"compress/gzip"
	"encoding/json"
	"go-movie-service/internal/module/movie"
	"go-movie-service/internal/packages"
	"go-movie-service/internal/types"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/noneandundefined/vision-go"
)

func setupTestMovieData(t *testing.T, id int, title string) {
	testData := types.Movies{
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
	testData := types.Movies{
		Page:         1,
		Results:      []types.Movie{},
		TotalPages:   0,
		TotalResults: 0,
	}

	createTestDataFile(t, testData)
}

func createTestDataFile(t *testing.T, testData types.Movies) {

	dataDir := filepath.Join("internal", "data")
	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create data directory: %v", err)
	}

	filePath := filepath.Join(dataDir, "movies.json.gz")
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test data file: %v", err)
	}
	defer file.Close()

	gzWriter := gzip.NewWriter(file)
	defer gzWriter.Close()

	encoder := json.NewEncoder(gzWriter)
	err = encoder.Encode(testData)
	if err != nil {
		t.Fatalf("Failed to encode test data: %v", err)
	}
}

func cleanupHandlerTestData(t *testing.T) {
	filePath := filepath.Join("internal", "data", "movies.json.gz")
	err := os.Remove(filePath)
	if err != nil && !os.IsNotExist(err) {
		t.Logf("Warning: Failed to remove test data file: %v", err)
	}
}

func TestMovieDetailsHandler(t *testing.T) {

	defer cleanupHandlerTestData(t)

	monitoring := vision.NewVision()
	errors := packages.NewErrors(monitoring)
	handler := movie.NewHandler(monitoring, errors)

	tests := []struct {
		name           string
		movieID        string
		setupMockData  func()
		wantStatusCode int
		wantTitle      string
		wantHeaders    map[string]string
	}{
		{
			name:    "Valid movie ID 1034541",
			movieID: "1034541",
			setupMockData: func() {
				setupTestMovieData(t, 1034541, "Terrifier 3")
			},
			wantStatusCode: http.StatusOK,
			wantTitle:      "Terrifier 3",
			wantHeaders: map[string]string{
				"Cache-Control": "public, max-age=3600",
				"Content-Type":  "application/json",
			},
		},
		{
			name:    "Invalid movie ID format",
			movieID: "invalid",
			setupMockData: func() {

			},
			wantStatusCode: http.StatusBadRequest,
			wantTitle:      "",
			wantHeaders:    map[string]string{},
		},
		{
			name:    "Movie not found",
			movieID: "999999",
			setupMockData: func() {
				setupEmptyTestData(t)
			},
			wantStatusCode: http.StatusNotFound,
			wantTitle:      "",
			wantHeaders:    map[string]string{},
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

				for key, value := range tt.wantHeaders {
					if rr.Header().Get(key) != value {
						t.Errorf("handler returned wrong %s header: got %v want %v",
							key, rr.Header().Get(key), value)
					}
				}

				expires := rr.Header().Get("Expires")
				if expires == "" {
					t.Errorf("handler did not set Expires header")
				}

				expireTime, err := time.Parse(time.RFC1123, expires)
				if err != nil {
					t.Errorf("invalid Expires header format: %v", err)
				} else if expireTime.Before(time.Now()) {
					t.Errorf("Expires header is set to past time: %v", expireTime)
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
