package handler

import (
	"encoding/json"
	"net/http"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	"github.com/Almazatun/golephant/pkg/common"
	logger "github.com/Almazatun/golephant/pkg/logger"
)

type positionTypeHandler struct {
	positionTypeRepo repository.PositionTypeRepo
}

type PositionTypeHandler interface {
	List(w http.ResponseWriter, r *http.Request)
}

func NewPositionTypeHandler(positionTypeRepo repository.PositionTypeRepo) PositionTypeHandler {
	return &positionTypeHandler{
		positionTypeRepo: positionTypeRepo,
	}
}

func (h *positionTypeHandler) List(w http.ResponseWriter, r *http.Request) {
	res, err := h.positionTypeRepo.List()

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
