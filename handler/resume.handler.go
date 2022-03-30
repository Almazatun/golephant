package handler

import (
	"encoding/json"
	"net/http"

	loggerinfo "github.com/Almazatun/golephant/common/loggerInfo"
	usecase "github.com/Almazatun/golephant/domain"
	"github.com/Almazatun/golephant/presentation/input"
	"github.com/gorilla/mux"
)

type resumeHandler struct {
	resumeUseCase usecase.ResumeUseCase
}

type ResumeHandler interface {
	CreateResume(w http.ResponseWriter, r *http.Request)
	DeleteResume(w http.ResponseWriter, r *http.Request)
	DeleteUserEducationInResume(w http.ResponseWriter, r *http.Request)
	DeleteUserExperienceInResume(w http.ResponseWriter, r *http.Request)
}

func NewResumeHandler(resumeUseCase usecase.ResumeUseCase) ResumeHandler {
	return &resumeHandler{
		resumeUseCase: resumeUseCase,
	}
}

func (h *resumeHandler) CreateResume(w http.ResponseWriter, r *http.Request) {
	var createResumeInput input.CreateResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&createResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.resumeUseCase.CreateResume(params["userId"], createResumeInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *resumeHandler) DeleteResume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.resumeUseCase.DeleteResume(params["resumeId"])

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *resumeHandler) DeleteUserEducationInResume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.resumeUseCase.DeleteUserEducationInResume(params["resumeId"], params["userEducationId"])

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *resumeHandler) DeleteUserExperienceInResume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.resumeUseCase.DeleteUserExperienceInResume(params["resumeId"], params["userExperienceId"])

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
