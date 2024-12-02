package translate

import (
	"encoding/json"
	"flick_finder/internal/types"
	"fmt"
	"io"
	"net/http"
)

// Перевод текста на английский язык
// ---------------------------------
func EN(text string) (string, error) {
	url := fmt.Sprintf("https://ftapi.pythonanywhere.com/translate?dl=en&text=%s", text)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response types.TranslationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return string(response.DestinationText), nil
}
