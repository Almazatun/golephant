package handler

import (
	"encoding/json"
	"net/http"

	usecase "github.com/Almazatun/golephant/internal/domain"
	"github.com/Almazatun/golephant/pkg/common"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
	logger "github.com/Almazatun/golephant/pkg/logger"
	"github.com/gorilla/mux"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

type UserHandler interface {
	UpdateUserData(w http.ResponseWriter, r *http.Request)
	GetLinkResetPassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

func (h *userHandler) UpdateUserData(w http.ResponseWriter, r *http.Request) {
	var updateUserDataInput input.UpdateUserDataInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateUserDataInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.userUseCase.UpdateData(params["userId"], updateUserDataInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *userHandler) GetLinkResetPassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.userUseCase.GetLinkResetPassword(params["userId"])

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *userHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var resetUserPasswordInput input.ResetUserPasswordInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&resetUserPasswordInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.userUseCase.ResetPassword(
		params["userId"],
		params["token"],
		resetUserPasswordInput.Password,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
