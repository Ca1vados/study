package main

import (
	"LoginPass/database"
	"LoginPass/entity"
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
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644) // флаги, разрешения, которые мы даем файлу
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}

	// Запись строки в файл
	if _, err := file.WriteString(expString); err != nil {
		log.Fatalf("Ошибка при записи в файл: %v", err)
	}
	file.Close()

	u1 := entity.User{
		Login:    "user1",
		PassHash: "e6c3da5b206634d7f3f3586d747ffdb36b5c675757b380c6a5fe5c570c714349",
		Secret:   "secret1",
	}
	u2 := entity.User{
		Login:    "user2",
		PassHash: "1ba3d16e9881959f8c9a9762854f72c6e6321cdd44358a10a4e939033117eab9",
		Secret:   "secret2",
	}

	res := make([]entity.User, 2)
	res[0] = u1
	res[1] = u2

	data, err := database.ReadJsonFile(filename)
	if err != nil {
		t.Errorf("func ReadJsonFile return error")
	}
	// os.Remove(filename)
	if len(data) != len(res) {
		t.Errorf("func ReadJsonFile return error")
	}

	for i, u := range res {
		if u.Login != data[i].Login || u.PassHash != data[i].PassHash || u.Secret != data[i].Secret {
			t.Errorf("Ожидалось %+v\nРезультат %+v\n", u, data[i])
		}
	}
}

func TestVeryficationPass(t *testing.T) {
	// создаю пользователя, передаю пароль, и сверяю пароль с должным. Сделать две проверки, на true и false. При правильном пароле вернуть тру, при неправильном пароле фолс
	UserPassTrue := entity.User{Login: "user1", PassHash: "pass1", Secret: ""}
	expectedAnswer1 := true

	db := database.New()

	for _, u := range db.Users {
		funcAnswer1 := u.VeryficationPass(UserPassTrue.PassHash)
		if funcAnswer1 != expectedAnswer1 {
			t.Errorf("answer = %v, expected %v", funcAnswer1, expectedAnswer1)
		}
	}

	expectedAnswer2 := false
	UserPassFalse := entity.User{Login: "user2", PassHash: "wrongPass", Secret: ""}
	for _, u := range db.Users {
		funcAnswer2 := u.VeryficationPass(UserPassFalse.PassHash)
		if funcAnswer2 != expectedAnswer2 {
			t.Errorf("answer = %v, expected %v", funcAnswer2, expectedAnswer2)
		}
	}
}

func TestWriteToJSONFile(t *testing.T) {
	filename := "TestWriteToJsonFile.json"
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
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	entity.WriteToJSONFile(file, expString)
}

// создаю список пользователей, записываю в файл, считываю в виде строки, сравниваю со строкой которая должна быть.
