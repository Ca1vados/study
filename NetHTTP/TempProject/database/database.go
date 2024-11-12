package database

import (
	"fmt"
	"temp/config"
)

type DataBase struct {
	name string
}

func New(cfg *config.Config) *DataBase {
	u := DataBase{
		name: cfg.DataBaseName,
	}
	return &u
}

func (db *DataBase) GetData() string {
	fmt.Println("(database) Вызвали функцию GetData")
	return "From Data Base: " + db.name
}
