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
Валюты:
BTC
LTC
USDT
ETH

При каждом запросе GET мы сохраняем в "бд" значение которое получили.
в функции в юскейсе добавить запись в базу данных + время записи

Создать хендлерна роутер history который вернет нам все из файла "бд"

Создать хендлер Convert , принимает в r.Body json, имеющий поля: from, to, amount (raw => json)
*/

/*
from to amount -> запрос
USDT BTC 500

amount FROM res TO
500 USDT = x BTC
*/
