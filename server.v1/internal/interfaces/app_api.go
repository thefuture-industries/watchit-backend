package interfaces

type IAppAPI interface {
	// Создание api_key
	CreateAPIKEY() string
	// Запись в БД api_key
	InsertAPIKEY(api_key string, uuid string) error
}
