package handler

import (
	"encoding/json"
	"net/http"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	"github.com/Almazatun/golephant/pkg/common"
	logger "github.com/Almazatun/golephant/pkg/logger"
)

type specializationHandler struct {
	specializationRepo repository.SpecializationRepo
}

type SpecializationHandler interface {
	List(w http.ResponseWriter, r *http.Request)
}

func NewSpecializationHandler(specializationRepo repository.SpecializationRepo) SpecializationHandler {
	return &specializationHandler{
		specializationRepo: specializationRepo,
	}
}

func (h *specializationHandler) List(w http.ResponseWriter, r *http.Request) {
	res, err := h.specializationRepo.List()

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
