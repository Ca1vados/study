package main

import (
	"fmt"
	"io/ioutil"
	"loginpass/database"
	"loginpass/entity"
	"net/http"
	"strings"
)

// добавить 2 маршрута
// singin - в body логин и пароль (json) - в ответе secret Или ошибка
// register - в body логин, пароль, секрет  - в ответе либо OK либо error

func Register(w http.ResponseWriter, r *http.Request) {
	byteLogPassSecret, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	stringLogPassSecret := string(byteLogPassSecret)
	arrLogPassSecret := strings.Split(stringLogPassSecret, " ")
	var u entity.User
}

func main() {
	db := database.New()
	fmt.Print(db)
	fmt.Println("LoginPass")
	http.HandleFunc("/singin", smfunc)
	http.HandleFunc("/register", Register)
	http.ListenAndServe(":8080", nil)
}
