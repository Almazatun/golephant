package handler

import (
	"encoding/json"
	"net/http"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
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
		logger.InfoError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
