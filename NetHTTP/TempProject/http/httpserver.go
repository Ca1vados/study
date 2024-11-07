package http

import (
	"net/http"
	"temp/config"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	Url    string
	router *mux.Router
}

func New(cfg *config.Config) *HttpServer {
	router := mux.NewRouter()
	u := HttpServer{
		Url:    cfg.URL,
		router: router,
	}
	return &u
}

func (hs *HttpServer) Run() {
	http.ListenAndServe(hs.Url, hs.router)
}

func GetData(w http.ResponseWriter, r *http.Request) {

}
