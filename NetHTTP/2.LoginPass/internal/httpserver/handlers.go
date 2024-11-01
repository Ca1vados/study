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

func (hs *HttpServer) SingIn(w http.ResponseWriter, r *http.Request) {

}
