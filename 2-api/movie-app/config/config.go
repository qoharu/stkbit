package config

// AppConfig application config
type AppConfig struct {
	MovieServiceClient MovieServiceClientConfig
	Addr               string
}

// MovieServiceClientConfig movie service client config
type MovieServiceClientConfig struct {
	Host    string
	Port    string
}
