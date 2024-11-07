package main

import (
	"temphttp/config"
	"temphttp/database"
	httpserver "temphttp/http"
	"temphttp/usecase"
)

// http_server - только обработка Http
//
// usecase
//
// database - только обработка запросов к базе данных

func main() {
	cfg := config.NewConfig()
	db := database.NewDatabase(cfg)
	u := usecase.NewUsecase(db, cfg)
	server := httpserver.NewHttpServer(u, cfg)
	server.Start()
}
