package handler

import (
	"encoding/json"
	"fmt"
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
	UpdateBasicInfoResume(w http.ResponseWriter, r *http.Request)
	UpdateAboutMeResume(w http.ResponseWriter, r *http.Request)
	UpdateCitizenshipResume(w http.ResponseWriter, r *http.Request)
	UpdateTagsResumeInput(w http.ResponseWriter, r *http.Request)
	UpdateUserEducationResume(w http.ResponseWriter, r *http.Request)
	UpdateUserExperiencesResume(w http.ResponseWriter, r *http.Request)
	UpdateDesiredPositionResume(w http.ResponseWriter, r *http.Request)
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

	res, err := h.resumeUseCase.CreateResume(params["userId"], createResumeInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *resumeHandler) UpdateBasicInfoResume(w http.ResponseWriter, r *http.Request) {
	var updateBasicInfoResumeInput input.UpdateBasicInfoResume
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateBasicInfoResumeInput)

	fmt.Println(updateBasicInfoResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateBasicInfoResume(params["userId"], params["resumeId"], updateBasicInfoResumeInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateAboutMeResume(w http.ResponseWriter, r *http.Request) {
	var updateAboutMeResumeInput input.UpdateAboutMeResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateAboutMeResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateAboutMeResume(params["userId"], params["resumeId"], updateAboutMeResumeInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateCitizenshipResume(w http.ResponseWriter, r *http.Request) {
	var updateCitizenshipResumeInput input.UpdateCitizenshipResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateCitizenshipResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateCitizenshipResume(params["userId"], params["resumeId"], updateCitizenshipResumeInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateTagsResumeInput(w http.ResponseWriter, r *http.Request) {
	var udateTagsResumeInput input.UdateTagsResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&udateTagsResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateTagsResumeInput(params["userId"], params["resumeId"], udateTagsResumeInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateUserEducationResume(w http.ResponseWriter, r *http.Request) {
	var updateUserEducationsResumeInput input.UpdateUserEducationsResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateUserEducationsResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateUserEducationResume(params["userId"], params["resumeId"], updateUserEducationsResumeInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateUserExperiencesResume(w http.ResponseWriter, r *http.Request) {
	var updateUserExperiencesResumeInput input.UpdateUserExperiencesResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateUserExperiencesResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateUserExperiencesResume(params["userId"], params["resumeId"], updateUserExperiencesResumeInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateDesiredPositionResume(w http.ResponseWriter, r *http.Request) {
	var updateDesiredPositionResumeInput input.UpdateDesiredPositionResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateDesiredPositionResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateDesiredPositionResume(params["userId"], params["resumeId"], updateDesiredPositionResumeInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

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
