package database

import (
	"database/sql"
	"loginpass/entity"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	conn *sql.DB
}

func (db *DataBase) AddUser(user entity.User) error {
	// ... добавить польователя в базу данных

	return nil
}

func (db *DataBase) GetUser(login string) (entity.User, error) {
	// запросить пользователя из базы по login

	return entity.User{}, nil
}

func New() *DataBase {
	// подключение к базе данных
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		os.Exit(1)
	}

	_, err = db.Exec(`create table if not exists users(
	id INT primary key,
	login TEXT,
	pass_hash TEXT, 
	secret TEXT)
	`)

	database := DataBase{conn: db}
	return &database
}
