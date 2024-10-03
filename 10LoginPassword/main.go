package main

import (
	"bufio"
	"fmt"
	"os"

	"LoginPass/database"
)

func ReadString(msg string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	input, err := reader.ReadString('\n') // !
	if err != nil {
		return input, err
	}
	return input, nil
}

func main() {
	db := database.New()
	login, _ := ReadString("Введите логин:")
	pass, _ := ReadString("Введите пароль:")

	for _, u := range db.Users {
		if login == u.Login {
			if u.VeryficationPass(pass) {
				fmt.Printf("Secret: %s !\n", u.Secret)
				return
			} else {
				break
			}
		}
	}

	fmt.Println("Неверный логин или пароль ...")
}
