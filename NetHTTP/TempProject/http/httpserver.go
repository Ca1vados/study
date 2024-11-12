package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"temp/config"
	"temp/usecase"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	Url    string
	router *mux.Router
	u      *usecase.UseCase
}

func New(u *usecase.UseCase, cfg *config.Config) *HttpServer {
	router := mux.NewRouter()
	hs := HttpServer{
		Url:    cfg.URL,
		router: router,
		u:      u,
	}
	hs.router.HandleFunc("/", nil)
	hs.router.HandleFunc("/data", hs.GetData).Methods("GET")
	return &hs
}

func (hs *HttpServer) Run() {
	http.ListenAndServe(hs.Url, hs.router)
}

func (hs *HttpServer) GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*httpserver - (hs *HttpServer) GetData")
	data := hs.u.GetData()
	json.NewEncoder(w).Encode(data)
}
