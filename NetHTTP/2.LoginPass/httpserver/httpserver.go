package httpserver

import (
	"loginpass/usecase"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type HttpServer struct {
	// какие-то поля
	u      *usecase.UseCase
	router *mux.Router
}

func New(u *usecase.UseCase) *HttpServer {
	router := mux.NewRouter()
	hs := HttpServer{u: u, router: router}
	hs.router.HandleFunc("/register", hs.Register).Methods("POST")
	hs.router.HandleFunc("/login", hs.Login).Methods("POST")
	hs.router.HandleFunc("/hello", hs.Hello).Methods("POST")

	hs.router.HandleFunc("/getallusers", hs.GetAllUsers).Methods("GET")

	hs.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return &hs
}

func (hs *HttpServer) Start() {
	//log.Fatal(http.ListenAndServe(":8080", hs.router))
	http.ListenAndServe(":8080", hs.router)
}
