package handler

import (
	"encoding/json"
	"net/http"

	repository "github.com/Almazatun/golephant/infrastucture"
)

type handler struct {
	userRepo repository.UserRepo
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func InitUserHandler(userRepo repository.UserRepo) UserHandler {
	return &handler{
		userRepo: userRepo,
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	h.userRepo.CreateUser()
	// fmt.Println(ctx)
}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
