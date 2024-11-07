package config

type Config struct {
	DataBaseName string
	URL          string
	LogLvl       string
}

func NewConfig() *Config {
	name := "db"
	url := ":8080"
	loglevel := "123"
	u := Config{
		DataBaseName: name,
		URL:          url,
		LogLvl:       loglevel,
	}
	return &u
}
