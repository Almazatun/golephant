package router

import (
	handler "github.com/Almazatun/golephant/handler"
	"github.com/gorilla/mux"
)

type Handler struct {
	UserHandler   handler.UserHandler
	ResumeHandler handler.ResumeHandler
}

func NewRouter(h Handler) *mux.Router {
	router := mux.NewRouter()

	//User
	user := router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", h.UserHandler.RegisterUser).Methods("POST")
	user.HandleFunc("/login", h.UserHandler.LogIn).Methods("POST")
	user.HandleFunc("/authMe", h.UserHandler.AuthMe).Methods("POST")
	user.HandleFunc("/{id}", h.UserHandler.UpdateUserData).Methods("PATCH")

	// Resume
	resume := router.PathPrefix("/resume").Subrouter()
	resume.HandleFunc("/{userId}", h.ResumeHandler.CreateResume).Methods("POST")

	router.HandleFunc("/helloWorld", handler.HelloWord).Methods("POST")

	return router
}
