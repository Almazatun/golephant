package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	usecase "github.com/Almazatun/golephant/internal/domain"
	repository "github.com/Almazatun/golephant/internal/infrastucture"
	"github.com/Almazatun/golephant/pkg/common"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
	logger "github.com/Almazatun/golephant/pkg/logger"
	"github.com/gorilla/mux"
)

type resumeHandler struct {
	resumeUseCase    usecase.ResumeUseCase
	resumeRepository repository.ResumeRepo
}

type ResumeHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	UpdateBasicInfo(w http.ResponseWriter, r *http.Request)
	UpdateAboutMe(w http.ResponseWriter, r *http.Request)
	UpdateCitizenship(w http.ResponseWriter, r *http.Request)
	UpdateTags(w http.ResponseWriter, r *http.Request)
	UpdateUserEducation(w http.ResponseWriter, r *http.Request)
	UpdateUserExperience(w http.ResponseWriter, r *http.Request)
	UpdateDesiredPosition(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	DeleteUserEducation(w http.ResponseWriter, r *http.Request)
	DeleteUserExperience(w http.ResponseWriter, r *http.Request)
}

func NewResumeHandler(
	resumeUseCase usecase.ResumeUseCase,
	resumeRepository repository.ResumeRepo,
) ResumeHandler {
	return &resumeHandler{
		resumeUseCase:    resumeUseCase,
		resumeRepository: resumeRepository,
	}
}

func (h *resumeHandler) List(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.resumeRepository.ListByUserId(params["userId"])

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *resumeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createResumeInput input.CreateResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&createResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.Create(params["userId"], createResumeInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *resumeHandler) UpdateBasicInfo(w http.ResponseWriter, r *http.Request) {
	var updateBasicInfoResumeInput input.UpdateBasicInfoResume
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateBasicInfoResumeInput)

	fmt.Println(updateBasicInfoResumeInput)

	if err != nil {
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateBasicInfo(
		params["userId"],
		params["resumeId"],
		updateBasicInfoResumeInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateAboutMe(w http.ResponseWriter, r *http.Request) {
	var updateAboutMeResumeInput input.UpdateAboutMeResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateAboutMeResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateAboutMe(
		params["userId"],
		params["resumeId"],
		updateAboutMeResumeInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateCitizenship(w http.ResponseWriter, r *http.Request) {
	var updateCitizenshipResumeInput input.UpdateCitizenshipResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateCitizenshipResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateCitizenship(
		params["userId"],
		params["resumeId"],
		updateCitizenshipResumeInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateTags(w http.ResponseWriter, r *http.Request) {
	var udateTagsResumeInput input.UdateTagsResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&udateTagsResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateTags(
		params["userId"],
		params["resumeId"],
		udateTagsResumeInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateUserEducation(w http.ResponseWriter, r *http.Request) {
	var updateUserEducationResumeInput input.UpdateUserEducationResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateUserEducationResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateUserEducation(
		params["userId"],
		params["resumeId"],
		updateUserEducationResumeInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateUserExperience(w http.ResponseWriter, r *http.Request) {
	var updateUserExperienceResumeInput input.UpdateUserExperienceResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateUserExperienceResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateUserExperience(
		params["userId"],
		params["resumeId"],
		updateUserExperienceResumeInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) UpdateDesiredPosition(w http.ResponseWriter, r *http.Request) {
	var updateDesiredPositionResumeInput input.UpdateDesiredPositionResumeInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateDesiredPositionResumeInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.resumeUseCase.UpdateDesiredPosition(
		params["userId"],
		params["resumeId"],
		updateDesiredPositionResumeInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (h *resumeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.resumeUseCase.Delete(params["resumeId"])

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *resumeHandler) DeleteUserEducation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.resumeUseCase.DeleteUserEducation(
		params["resumeId"],
		params["userEducationId"],
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *resumeHandler) DeleteUserExperience(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.resumeUseCase.DeleteUserExperience(
		params["resumeId"],
		params["userExperienceId"],
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res)
}
