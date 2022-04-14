package handler

import (
	"encoding/json"
	"net/http"

	loggerinfo "github.com/Almazatun/golephant/common/loggerInfo"
	repository "github.com/Almazatun/golephant/infrastucture"
)

type positionTypeHanler struct {
	positionTypeRepo repository.PositionTypeRepo
}

type PositionTypeHandler interface {
	PositionTypes(w http.ResponseWriter, r *http.Request)
}

func NewPositionTypeHandler(positionTypeRepo repository.PositionTypeRepo) PositionTypeHandler {
	return &positionTypeHanler{
		positionTypeRepo: positionTypeRepo,
	}
}

func (h *positionTypeHanler) PositionTypes(w http.ResponseWriter, r *http.Request) {
	res, err := h.positionTypeRepo.List()

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
