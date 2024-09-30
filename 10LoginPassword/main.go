package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Login    string `yaml:"Id"`
	Pass     string `json:"Pass"`
	PassHash string `json:""`
	Secret   string `json:"Secret"`
}

func main() {
	var UserInput User
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите логин:")
	Login, _ := reader.ReadString('\n')
	UserInput.Login = Login
	fmt.Println("Введите пароль:")
	Password, _ := reader.ReadString('\n')
	UserInput.Pass = Password
	fmt.Printf("Структура пользователя: %+v\n", UserInput)
}
