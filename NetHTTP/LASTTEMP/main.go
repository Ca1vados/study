package main

import (
	"temp/config"
	"temp/database"
	"temp/http"
	"temp/usecase"
)

func main() {
	cfg := config.New()
	db := database.New(cfg)
	usecase := usecase.New(db, cfg)
	server := http.New(usecase, cfg)
	server.Start()
}

/* Задание:
При каждом запросе GET мы сохраняем в "бд" значение которое получили. в функции в юскейсе добавить запись в базу данных + время записи
Создать хендлерна роутер histroy который вернет нам все из файла "бд"
Создать хендлер Convert , принимает в r.Body json, имеющий поля: from, to, amount (raw => json)
*/
