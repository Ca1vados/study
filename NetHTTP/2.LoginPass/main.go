package main

import (
	"loginpass/database"
	_ "loginpass/docs"
	"loginpass/httpserver"
	"loginpass/usecase"
)

// http <-> usecase <-> database
//   ^        ^           ^
// 			entity

// добавить 2 маршрута
// singin - в body логин и пароль (json) - в ответе secret Или ошибка
// register - в body логин, пароль, секрет  - в ответе либо OK либо error

// @title Login password service
// @version 1.0
// @description service for authorization by login password
// @termsOfService http://swagger.io/terms/

// @host localhost:8080

func main() {
	database_path := "./database.db"
	db := database.New(database_path)

	u := usecase.New(db)
	server := httpserver.New(u)

	server.Start()
}
