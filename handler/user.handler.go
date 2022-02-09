package handler

import (
	"encoding/json"
	"net/http"

	loggerinfo "github.com/Almazatun/golephant/common/loggerInfo"
	usecase "github.com/Almazatun/golephant/domain"
	"github.com/Almazatun/golephant/infrastucture/model"
	"github.com/Almazatun/golephant/presentation/input"
)

type handler struct {
	userUseCase usecase.UserUseCase
}

type UserHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &handler{
		userUseCase: userUseCase,
	}
}

func (h *handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registerUserInput *model.User
	err := json.NewDecoder(r.Body).Decode(&registerUserInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.RegisterUser(registerUserInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var logInInput *input.LogIn
	err := json.NewDecoder(r.Body).Decode(&logInInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	str, err := h.userUseCase.LogIn(logInInput)

	json.NewEncoder(w).Encode(str)

}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
