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
	// INSERT INTO user (login, pass_hash, secret) VALUES (?, ?, ?)
	_, err := db.conn.Exec("INSERT INTO users (login, pass_hash, secret) VALUES (?, ?, ?)", user.Login, user.PassHash, user.Secret)
	return err
}

func (db *DataBase) GetUser(login string) (entity.User, error) {
	// запросить пользователя из базы по login
	// SELECT * FROM users WHERE login = '?'

	rows, err := db.conn.Query("SELECT * FROM users WHERE login = '?'", login)
	if err != nil {
		return entity.User{}, err
	}
	defer rows.Close()

	var requestUser entity.User
	for rows.Next() {
		var u entity.User
		err := rows.Scan(&u.Login, &u.PassHash, &u.Secret)
		if err != nil {
			return entity.User{}, err
		}
		requestUser = u
	}

	return requestUser, nil
}

func (db *DataBase) GetAllLogins() ([]string, error) {
	// запрос в базу для проверки существования логина
	// SELECT login FROM users
	rows, err := db.conn.Query("SELECT login FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logins []string
	for rows.Next() {
		var login string
		err = rows.Scan(&login)
		if err != nil {
			return nil, err
		}
		logins = append(logins, login)
	}

	return logins, nil
}

func (db *DataBase) GetAllUsers() ([]entity.User, error) {
	// запроса к базе данных
	// select * from users;
	rows, err := db.conn.Query("SELECT * from users")
	if err != nil {
		return nil, err
	}
	users := []entity.User{}
	for rows.Next() {
		u := entity.User{}
		err := rows.Scan(&u.Login, &u.PassHash, &u.Secret)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
