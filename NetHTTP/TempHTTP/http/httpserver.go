package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"temphttp/config"
	"temphttp/usecase"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	url    string
	router *mux.Router
	u      *usecase.Usecase
}

func NewHttpServer(u *usecase.Usecase, cfg *config.Config) *HttpServer {
	router := mux.NewRouter()
	hs := HttpServer{
		url:    cfg.HttpURL,
		router: router,
		u:      u,
	}
	hs.router.HandleFunc("/", nil)
	hs.router.HandleFunc("/data", hs.GetData).Methods("GET")
	hs.router.HandleFunc("/data", hs.PutData).Methods("PUT")
	return &hs
}

func (hs *HttpServer) Start() {
	http.ListenAndServe(hs.url, hs.router)
}

func (hs *HttpServer) GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*httpserver - (hs *HttpServer) GetData")
	data := hs.u.GetData()
	json.NewEncoder(w).Encode(data)
}

func (hs *HttpServer) PutData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*httpserver - (hs *HttpServer) PutData")
	data, _ := io.ReadAll(r.Body)
	hs.u.PutData(string(data))
}
