package movie

type SuggestPayload struct {
	Text string `json:"text"`
}

type SearchPayload struct {
	Text string `json:"text"`
}
