package database

import (
	"temp/config"
)

type DataBase struct {
	name string
}

func New(cfg *config.Config) *DataBase {
	db := DataBase{
		name: cfg.DataBaseName,
	}
	return &db
}
