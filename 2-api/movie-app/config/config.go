package config

// AppConfig application config
type AppConfig struct {
	MovieDBConfig MovieDBConfig
	Addr          string
}

type MovieDBConfig struct {
	APIKey  string
	BaseURL string
}
