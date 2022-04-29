package handler

import (
	"encoding/json"
	"net/http"

	repository "github.com/Almazatun/golephant/internal/infrastucture"
	logger "github.com/Almazatun/golephant/pkg/logger"
)

type resumeStatusHandler struct {
	resumeStatusRepo repository.ResumeStatusRepo
}

type ResumeStatusHandler interface {
	ResumeStatuses(w http.ResponseWriter, r *http.Request)
}

func NewResumeStatusHandler(resumeStatusRepo repository.ResumeStatusRepo) ResumeStatusHandler {
	return &resumeStatusHandler{
		resumeStatusRepo: resumeStatusRepo,
	}
}

func (h *resumeStatusHandler) ResumeStatuses(w http.ResponseWriter, r *http.Request) {
	res, err := h.resumeStatusRepo.List()

	if err != nil {
		logger.InfoError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
