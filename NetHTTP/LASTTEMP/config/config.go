package config

type Config struct {
	Url          string
	LogLvl       string
	DataBaseName string
}

func New() *Config {
	url := ":8080"
	loglvl := "debug"
	name := "MyDb.json"
	cfg := Config{
		Url:          url,
		LogLvl:       loglvl,
		DataBaseName: name,
	}
	return &cfg
}
