package router

import (
	handler "github.com/Almazatun/golephant/handler"
	"github.com/gorilla/mux"
)

type Handler struct {
	UserHandler handler.UserHandler
}

func NewRouter(h Handler) *mux.Router {
	router := mux.NewRouter()

	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", h.UserHandler.CreateUser).Methods("POST")

	router.HandleFunc("/helloWorld", handler.HelloWord).Methods("GET")

	return router
}