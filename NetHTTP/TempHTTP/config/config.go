package config

type Config struct {
	HttpURL      string
	DataBaseName string
	LogLevel     string
}

func NewConfig() *Config {
	http_url := ":8080"
	database_name := "mydatabase"
	log_level := "ERROR"
	return &Config{
		HttpURL:      http_url,
		DataBaseName: database_name,
		LogLevel:     log_level,
	}
}
