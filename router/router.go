package router

import (
	"net/http"

	handler "github.com/Almazatun/golephant/handler"
	"github.com/Almazatun/golephant/middleware"
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
	user.HandleFunc("/login", h.UserHandler.LogIn).Methods("PUT")
	user.HandleFunc("/authMe", h.UserHandler.AuthMe).Methods("POST")
	user.Handle("/{id}", middleware.Authentication(http.HandlerFunc(h.UserHandler.UpdateUserData))).Methods("PATCH")

	// Resume
	resume := router.PathPrefix("/resume").Subrouter()
	resume.Handle("/{userId}", middleware.Authentication(http.HandlerFunc(h.ResumeHandler.CreateResume))).Methods("POST")
	resume.Handle("/{resumeId}", middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteResume))).Methods("DELETE")

	// User education in resume
	resume.Handle("/{resumeId}/userEducation/{userEducationId}", middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteUserEducationInResume))).Methods("DELETE")

	// User experience in resume
	resume.Handle("/{resumeId}/userExperience/{userExperienceId}", middleware.Authentication(http.HandlerFunc(h.ResumeHandler.DeleteUserExperienceInResume))).Methods("DELETE")

	// Test
	router.HandleFunc("/helloWorld", handler.HelloWord).Methods("GET")

	return router
}
