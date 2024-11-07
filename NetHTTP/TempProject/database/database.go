package database

import "temp/config"

type DataBase struct {
	name string
}

func New(cfg *config.Config) *DataBase {
	u := DataBase{
		name: cfg.DataBaseName,
	}
	return &u
}
