package handler

import (
	"encoding/json"
	"net/http"

	usecase "github.com/Almazatun/golephant/domain"
)

type handler struct {
	userUseCase usecase.UserUseCase
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func InitUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &handler{
		userUseCase: userUseCase,
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	h.userUseCase.CreateUser()
	// fmt.Println(ctx)
}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
