package main

import (
	"bufio"
	"fmt"
	"os"

	"./user"
)

func main() {
	var UserInput user.User
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите логин:")
	Login, _ := reader.ReadString('\n')
	UserInput.Login = Login
	fmt.Println("Введите пароль:")
	Password, _ := reader.ReadString('\n')
	UserInput.Pass = Password
	fmt.Printf("Структура пользователя: %+v\n", UserInput)
}
