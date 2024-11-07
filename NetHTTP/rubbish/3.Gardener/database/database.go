package database

import (
	"encoding/json"
	"fmt"
	"garden/entity"
	"os"
)

type DataBase struct {
	Plunts []entity.Plunt
}

func New(u string) *Database {
	plunts, err := ReadJsonFile(u)
	if err != nil {
		fmt.Println(err.Error())
		panic(1)
	}

	return &Database{Plunts: plunts}
}

func ReadJsonFile(file_name string) ([]entity.Plunt, error) {
	file, err := os.Open(file_name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var plunts []entity.Plunt
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&plunts)
	return plunts, err
}
