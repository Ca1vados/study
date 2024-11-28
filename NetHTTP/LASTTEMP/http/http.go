package http

import (
	"net/http"
	"temp/config"
	"temp/usecase"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// handlers: get () convert history
type HttpServer struct {
	url    string
	router *mux.Router
	u      *usecase.UseCase
}

func New(usecase *usecase.UseCase, cfg *config.Config) *HttpServer {
	router := mux.NewRouter()
	hs := HttpServer{
		url:    cfg.Url,
		router: router,
		u:      usecase,
	}
	hs.router.HandleFunc("/hello", hs.Hello).Methods(http.MethodPost)
	hs.router.HandleFunc("/get", hs.GetCryptoRataes).Methods(http.MethodGet)

	hs.router.HandleFunc("/history", hs.History).Methods(http.MethodGet)
	hs.router.HandleFunc("/convert", hs.Convert).Methods(http.MethodPost)

	hs.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return &hs
}

func (hs *HttpServer) Start() {
	http.ListenAndServe(hs.url, hs.router)
}
