package main

import (
	"garden/database"
	"garden/httpserver"
	"garden/usecase"
)

func main() {
	database_path := "./Garden.json"
	db := database.New(database_path)

	u := usecase.New(db)
	server := httpserver.New(u)

	server.Start()
}
