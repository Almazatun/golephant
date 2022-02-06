package handler

import (
	"encoding/json"
	"net/http"

	loggerinfo "github.com/Almazatun/golephant/common/loggerInfo"
	usecase "github.com/Almazatun/golephant/domain"
	"github.com/Almazatun/golephant/infrastucture/model"
)

type handler struct {
	userUseCase usecase.UserUseCase
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &handler{
		userUseCase: userUseCase,
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserInput *model.User
	err := json.NewDecoder(r.Body).Decode(&createUserInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.CreateUser(createUserInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
