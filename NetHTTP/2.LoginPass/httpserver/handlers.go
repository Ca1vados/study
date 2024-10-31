package httpserver

import (
	"encoding/json"
	"loginpass/entity"
	"net/http"
)

func (hs *HttpServer) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err := hs.u.RegisterUser(user)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
}

func (hs *HttpServer) Login(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err := hs.u.Login(user)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
}

func (hs *HttpServer) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func (hs *HttpServer) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// 1. Обработка входных данных и продготовка данных для передачи в UseCase
	// тут нет обработки request, так как это просто запрос данных

	// 2. Вызов метода UseCase - внутри вся логика
	users, err := hs.u.GetAllUsers() // вызов UseCase

	// 3. Отправка ответа клиенту
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(users)
	}
}
