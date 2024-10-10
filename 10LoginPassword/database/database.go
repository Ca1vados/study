package database

import (
	"LoginPass/entity"
	"encoding/json"
	"os"
)

type Database struct {
	Users []entity.User
}

func New() *Database {
	users, err := ReadJsonFile("database/LogPass.json")
	if err != nil {
		panic(1) // вызывает эту панику
	}

	return &Database{Users: users}
}

// разобраться
func ReadJsonFile(file_name string) ([]entity.User, error) {
	file, err := os.Open(file_name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var users []entity.User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	return users, err
}
