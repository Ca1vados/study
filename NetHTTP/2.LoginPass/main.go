package main

import (
	"fmt"
	"loginpass/database"
)

// добавить 2 маршрута
// singin - в body логин и пароль (json) - в ответе secret Или ошибка
// register - в body логин, пароль, секрет  - в ответе либо OK либо error
func main() {
	db := database.New()
	fmt.Print(db)
	fmt.Println("LoginPass")
}
