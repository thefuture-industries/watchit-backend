package types

import (
	"sync"
	"time"
)

// User: Моделька пользователя в системе БД
// Id, UUID, UserName, UserNameUpper, Email, EmailUpper, IPAddress, Lat, Lon, Country, RegionName, Zip, CreatedAt
type User struct {
	ID            int     `json:"id"`
	UUID          string  `json:"uuid"`
	SecretWord    string  `json:"secret_word"`
	UserName      string  `json:"username"`
	UserNameUpper string  `json:"username_upper"`
	Email         *string `json:"email"`
	EmailUpper    *string `json:"email_upper"`
	IPAddress     string  `json:"ip_address"`
	Lat           string  `json:"latitude"`
	Lon           string  `json:"longitude"`
	Country       string  `json:"country"`
	RegionName    string  `json:"regionName"`
	Zip           string  `json:"zip"`
	CreatedAt     string  `json:"createdAt"`
}

// Favourites: Моделька избранных фильмов в системе БД
// Id, UUID, MovieID, MoviePoster, CreatedAt
type Favourites struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid"`
	MovieID     int    `json:"movieId"`
	MoviePoster string `json:"moviePoster"`
	CreatedAt   string `json:"createdAt"`
}

// API_KEYS: Моделька api keys в системе БД
// Id, UUID, ApiKey, CreatedAt
type API_KEYS struct {
	ID        int    `json:"id"`
	UUID      string `json:"uuid"`
	ApiKey    string `json:"api_key"`
	CreatedAt string `json:"createdAt"`
}

// Limiter: Моделька лимитов действий в системе БД
// Id, UUID, TextLimiter, YoutubeLimit, UpdateAt
type Limiter struct {
	ID           int    `json:"id"`
	UUID         string `json:"uuid"`
	TextLimiter  int    `json:"text_limit"`
	YoutubeLimit int    `json:"youtube_limit"`
	UpdateAt     string `json:"update_at"`
}

// Recommendations: Моделька предпочтений пользователя в системе БД
// Id, UUID, Title, Genre (string[] "27,28")
type Recommendations struct {
	ID    int    `json:"id"`
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

// JsonMovies: Модель данных фильмов
// Структурирует модель данных json из запроса к tmdb
type JsonMovies struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

// Movie: Модель для JsonMovies данные для Result json
type Movie struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	ReleaseDate      string  `json:"release_date"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

// Структура для всего JSON
type KeywordList struct {
	Keywords []KeywordCategory `json:"keywords"`
}

// Структура для категории ключевых слов
type KeywordCategory struct {
	GenreID int      `json:"genre_id"`
	Words   []string `json:"words"`
}

// Получение авторизационных данных Giga Chat
type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

// ChatCompletionResponse описывает структуру ответа GIGA CHAT.
type ChatCompletionResponse struct {
	Choices []Choice `json:"choices"`
	Created int64    `json:"created"` // Время создания в формате Unix Timestamp
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Object  string   `json:"object"`
}

// Choice описывает выбор в рамках ответа.
type Choice struct {
	Message          Message       `json:"message"`
	Index            int           `json:"index"`
	FinishReason     string        `json:"finish_reason"`
	DataForContext   []interface{} `json:"data_for_context"`
	FunctionsStateID string        `json:"functions_state_id,omitempty"`
	FunctionCall     FunctionCall  `json:"function_call,omitempty"`
}

// Message описывает сообщение внутри выбора.
type Message struct {
	Role             string       `json:"role"`
	Content          string       `json:"content"`
	FunctionsStateID string       `json:"functions_state_id,omitempty"`
	FunctionCall     FunctionCall `json:"function_call,omitempty"`
}

// FunctionCall описывает вызов функции.
type FunctionCall struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// Usage описывает использование токенов.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Type для api translate
type TranslationResponse struct {
	SourceLanguage      string        `json:"source-language"`
	SourceText          string        `json:"source-text"`
	DestinationLanguage string        `json:"destination-language"`
	DestinationText     string        `json:"destination-text"`
	Pronunciation       Pronunciation `json:"pronunciation"`
	Translations        Translations  `json:"translations"`
	Definitions         []Definition  `json:"definitions"`
	SeeAlso             interface{}   `json:"see-also"`
}

type Pronunciation struct {
	SourceTextPhonetic   string `json:"source-text-phonetic"`
	SourceTextAudio      string `json:"source-text-audio"`
	DestinationTextAudio string `json:"destination-text-audio"`
}

type Translations struct {
	AllTranslations      [][]interface{} `json:"all-translations"`
	PossibleTranslations []string        `json:"possible-translations"`
	PossibleMistakes     interface{}     `json:"possible-mistakes"`
}

type Definition struct {
	PartOfSpeech  string      `json:"part-of-speech"`
	Definition    string      `json:"definition"`
	Example       string      `json:"example"`
	OtherExamples interface{} `json:"other-examples"`
	Synonyms      interface{} `json:"synonyms"`
}

// Моделька для мониторинга приложения
type MonitoringStats struct {
	sync.Mutex
	RequestCount   int64
	ErrorCount     int64
	TotalLatency   time.Duration
	DBQueryCount   int64
	DBErrorCount   int64
	DBTotalLatency time.Duration
	LastErrors     []ErrorLog
}

// Моделька для ошибок в приложении
type ErrorLog struct {
	Timestamp time.Time `json:"timestamp"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Error     string    `json:"error"`
}

// Моделька для ответа мониторинга приложения
type MonitoringResponse struct {
	Requests struct {
		Total        int64   `json:"total"`
		Errors       int64   `json:"errors"`
		SuccessRate  float64 `json:"success_rate"`
		AvgLatencyMs float64 `json:"avg_latency_ms"`
	} `json:"requests"`
	Database struct {
		TotalQueries int64   `json:"total_queries"`
		Errors       int64   `json:"errors"`
		AvgLatencyMs float64 `json:"avg_latency_ms"`
	} `json:"database"`
	LastErrors []ErrorLog `json:"last_errors,omitempty"`
}

// тип DTO от пользователя
// Требуется для входа в систему
type LoginUserPayload struct {
	UUID string `json:"uuid"`
}

// Тип DTO для обновления данных пользователя
type UserUpdate struct {
	UUID       string  `json:"uuid"`
	Username   *string `json:"username"`
	Email      *string `json:"email"`
	SecretWord *string `json:"secret_word"`
}

// тип DTO от пользователя
// Требуется для регестрации в системе
type RegisterUserPayload struct {
	UserName   string `json:"username"`
	Email      string `json:"email"`
	IPAddress  string `json:"ip_address" validate:"required"`
	Lat        string `json:"latitude" validate:"required"`
	Lon        string `json:"longitude" validate:"required"`
	Country    string `json:"country" validate:"required"`
	RegionName string `json:"regionName" validate:"required"`
	Zip        string `json:"zip" validate:"required"`
}

// CreateAPIKEYPayload: Модель DTO данные от пользователя
// Создание API_KEY
type CreateAPIKEYPayload struct {
	Email string `json:"email" validate:"required"`
}

type FavouriteAddPayload struct {
	UUID        string `json:"uuid" validate:"required"`
	MovieID     int    `json:"movieId" validate:"required"`
	MoviePoster string `json:"moviePoster" validate:"required"`
}

type FavouriteDeletePayload struct {
	UUID    string `json:"uuid" validate:"required"`
	MovieID int    `json:"movieId" validate:"required"`
}

type RecommendationAddPayload struct {
	UUID  string `json:"uuid" validate:"required"`
	Title string `json:"title" validate:"required"`
	Genre string `json:"genre" validate:"required"`
}

type TextMoviePayload struct {
	UUID string `json:"uuid" validate:"required"`
	Text string `json:"text" validate:"required"`
	Lege string `json:"lege" validate:"required"`
}
