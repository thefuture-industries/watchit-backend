package movie

type SuggestPayload struct {
	Text string `json:"text" validate:"required"`
}

type SearchPayload struct {
	Text string `json:"text" validate:"required"`
}
