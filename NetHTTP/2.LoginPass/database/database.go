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

func New(database_path string) *DataBase {
	// подключение к базе данных
	db, err := sql.Open("sqlite3", database_path)
	if err != nil {
		os.Exit(1)
	}

	_, err = db.Exec(`create table if not exists users (
	login TEXT primary key,
	pass_hash TEXT, 
	secret TEXT)
	`)

	database := DataBase{conn: db}
	return &database
}

func (db *DataBase) CreateUser(user entity.User) error {
	// ... добавить польователя в базу данных
	// db.conn.Exec("INSERT ...")
	return nil
}

func (db *DataBase) GetUser(login string) (entity.User, error) {
	// запросить пользователя из базы по login

	return entity.User{}, nil
}

func (db *DataBase) GetAllLogins() ([]string, error) {
	// запрос в базу для проверки существования логина

	return nil, nil
}

func (db *DataBase) GetAllUsers() ([]entity.User, error) {
	// запроса к базе данных
	// select * from users;
	users := []entity.User{}

	return users, nil
}
