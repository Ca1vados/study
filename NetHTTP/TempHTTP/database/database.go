package database

import (
	"fmt"
	"temphttp/config"
)

type Database struct {
	Name string
	// conn *sql.Connect
}

func NewDatabase(cfg *config.Config) *Database {
	return &Database{
		Name: cfg.DataBaseName,
	}
}

func (db *Database) GetData() string {
	fmt.Println("(database) Вызвали функцию GetData")
	return "From Data Base: " + db.Name
}

func (db *Database) CreateData(s string) {
	fmt.Println("(database) Вызвали функцию CreateData")
	fmt.Printf("\tDatabase (%s): %s\n", db.Name, s)
}
