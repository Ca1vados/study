package main

import (
	"temp/config"
	"temp/database"
	"temp/http"
	"temp/usecase"
)

func main() {
	cfg := config.NewConfig()
	db := database.New(cfg)
	u := usecase.New(db, cfg)
	server := http.New(u, cfg)
	server.Run()
}
