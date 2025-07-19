package movie

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"watchit/httpx/pkg/httpx/httperr"
	"watchit/httpx/pkg/httpx/httpprox"

	"github.com/gorilla/mux"
)

func (h *Handler) GetMovieImageHandler(w http.ResponseWriter, r *http.Request) error {
	image := mux.Vars(r)["image"]
	if image == "" {
		return httperr.NotFound("image not found")
	}

	url := fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s?api_key=%s", image, os.Getenv("TMDB_KEY_API"))

	httpConfig, err := httpprox.GetProxy()
	if err != nil {
		return httperr.InternalServerError(err.Error())
	}

	client := &http.Client{
		Transport: httpConfig,
		Timeout:   15 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return httperr.New(err.Error(), http.StatusBadGateway)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return httperr.New("error send request to get image", http.StatusBadGateway)
	}

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(http.StatusOK)

	_, _ = io.Copy(w, resp.Body)
	return nil
}
