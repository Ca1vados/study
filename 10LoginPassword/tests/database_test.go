package tests

import (
	"LoginPass/database"
	"log"
	"os"
	"testing"
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
	expString := "1"
	res := 1
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
	if data != res {
		t.Errorf("data = %v, expected %v", data, res)
	}

}
