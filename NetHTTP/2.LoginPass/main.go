package main

import (
	"loginpass/database"
	"loginpass/httpserver"
	"loginpass/usecase"
)

// http <-> usecase <-> database
//   ^        ^           ^
// 			entity

// добавить 2 маршрута
// singin - в body логин и пароль (json) - в ответе secret Или ошибка
// register - в body логин, пароль, секрет  - в ответе либо OK либо error

func main() {
	database_path := "./database.db"
	db := database.New(database_path)

	u := usecase.New(db)
	server := httpserver.New(u)

	server.Start()
}
