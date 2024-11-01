package main

import (
	"loginpass/database"
	"loginpass/internal/httpserver"
	"loginpass/internal/usecase"
)

// http <-> usecase <-> database
//   ^        ^           ^
// 			entity

// добавить 2 маршрута
// singin - в body логин и пароль (json) - в ответе secret Или ошибка
// register - в body логин, пароль, секрет  - в ответе либо OK либо error

func main() {
	db := database.New()        // создает файл database с таблицей в нем посредством функции database.New
	u := usecase.New(db)        // ...
	server := httpserver.New(u) // запускает http server с помощью httpserver.New
	server.Start()              // запускаем сервер (почему нет ссылки на пакет httpserver?)
}
