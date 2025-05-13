package movie_test

import (
	"compress/gzip"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"go-movie-service/internal/module/movie"
	"go-movie-service/internal/types"
)

func setupTestData(t *testing.T) {
	testData := types.JsonMovies{
		Page: 1,
		Results: []types.Movie{
			{
				Id:               1034541,
				Title:            "Terrifier 3",
				OriginalTitle:    "Terrifier 3",
				OriginalLanguage: "en",
				Overview:         "Test overview",
				ReleaseDate:      "2024-10-09",
				VoteAverage:      7.201,
			},
			{
				Id:               1037855,
				Title:            "Alien: Romulus",
				OriginalTitle:    "Alien: Romulus",
				OriginalLanguage: "en",
				Overview:         "Test overview 2",
				ReleaseDate:      "2024-08-15",
				VoteAverage:      7.3,
			},
		},
		TotalPages:   1,
		TotalResults: 2,
	}

	setupTestDataWithContent(t, testData)
}

func setupTestDataWithContent(t *testing.T, testData types.JsonMovies) {
	dataDir := filepath.Join("test_data")
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

func TestMovieDetails(t *testing.T) {
	setupTestData(t)

	tests := []struct {
		name     string
		id       int
		wantID   int
		wantName string
		wantErr  bool
	}{
		{
			name:     "Valid movie ID 1034541",
			id:       1034541,
			wantID:   1034541,
			wantName: "Terrifier 3",
			wantErr:  false,
		},
		{
			name:     "Valid movie ID 1037855",
			id:       1037855,
			wantID:   1037855,
			wantName: "Alien: Romulus",
			wantErr:  false,
		},
		{
			name:     "Invalid movie ID",
			id:       999999,
			wantID:   0,
			wantName: "",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := movie.MovieDetails(tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("MovieDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Id != tt.wantID {
				t.Errorf("MovieDetails() got ID = %v, want %v", got.Id, tt.wantID)
			}

			if got.Title != tt.wantName {
				t.Errorf("MovieDetails() got Title = %v, want %v", got.Title, tt.wantName)
			}
		})
	}
}
