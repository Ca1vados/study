package tests

import "testing"

/*
тест на чтение json файла.
Для этого создаем переменную с данными структуры User
и создаем на её основе файл Json
затем загнать файл в функцию ReadJsonFile и проверить получим ли мы на выходе
такую же структуру
*/

func TestReadJsonFile(t *testing.T) {
	var NewUser = User{Login: user3, PassHash: somePass, Secret: someSecret}

}
