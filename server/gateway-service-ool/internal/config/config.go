package config

type Config struct {
	Security string

	ServerAddr string
	ServerPort string
	// microservces
	UserAddr string
	UserPort string
}

func NewConfig() *Config {
	return &Config{
		Security:   "http",
		ServerAddr: "127.0.0.1",
		ServerPort: ":8080",
		UserAddr:   "127.0.0.1",
		UserPort:   ":8001",
	}
}
