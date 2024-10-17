package tests

import (
	"LoginPass/database"
	"LoginPass/entity"
	"log"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

/*
тест на чтение json файла.
Для этого создаем переменную с данными структуры User
и создаем на её основе файл Json
затем загнать файл в функцию ReadJsonFile и проверить получим ли мы на выходе
такую же структуру
*/

func TestReadJsonFile(t *testing.T) {
	filename := "TestReadJsonFile.json"
	expString := `[
		{
			"login": "user1",
			"hash": "e6c3da5b206634d7f3f3586d747ffdb36b5c675757b380c6a5fe5c570c714349",
			"secret": "secret1"
		},
		{
			"login": "user2",
			"hash": "1ba3d16e9881959f8c9a9762854f72c6e6321cdd44358a10a4e939033117eab9",
			"secret": "secret2"
		}
	]`
	res := []entity.User{}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // флаги, разрешения, которые мы даем файлу
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}

	// Запись строки в файл
	if _, err := file.WriteString(expString); err != nil {
		log.Fatalf("Ошибка при записи в файл: %v", err)
	}
	file.Close()
	data, err := database.ReadJsonFile(filename)
	if err != nil {
		t.Errorf("func ReadJsonFile return error")
	}
	answer := cmp.Equal(data, res)
	if !answer { // изначально было if data != res. ругалось на то что нельзя две структуры так сравнить
		t.Errorf("data = %v, expected %v", data, res)
	}

}
