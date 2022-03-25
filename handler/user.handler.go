package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	loggerinfo "github.com/Almazatun/golephant/common/loggerInfo"
	usecase "github.com/Almazatun/golephant/domain"
	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/Almazatun/golephant/presentation/input"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

type UserHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

func (h *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registerUserInput *entity.User
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

func (h *userHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var logInInput input.LogIn
	err := json.NewDecoder(r.Body).Decode(&logInInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(logInInput)

	tokenString, err := h.userUseCase.LogIn(logInInput)

	if err != nil {
		addHttpBodyResponse(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(tokenString)

}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func addHttpBodyResponse(w http.ResponseWriter, err error) {
	resp := make(map[string]string)

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	resp["message"] = err.Error()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
