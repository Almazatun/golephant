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

type companyHandler struct {
	companyUseCase usecase.CompanyUseCase
}

type CompanyHandler interface {
	AddAddress(w http.ResponseWriter, r *http.Request)
	DeleteAddress(w http.ResponseWriter, r *http.Request)
	AddPosition(w http.ResponseWriter, r *http.Request)
	UpdatePositionStatus(w http.ResponseWriter, r *http.Request)
	UpdatePositionResponsibilities(w http.ResponseWriter, r *http.Request)
	UpdatePositionRequirements(w http.ResponseWriter, r *http.Request)
	UpdatePosition(w http.ResponseWriter, r *http.Request)
	DeletePosition(w http.ResponseWriter, r *http.Request)
}

func NewCompanyHandler(
	companyUseCase usecase.CompanyUseCase,
) CompanyHandler {
	return &companyHandler{
		companyUseCase: companyUseCase,
	}
}

func (h *companyHandler) AddAddress(w http.ResponseWriter, r *http.Request) {
	var createCompanyAddressInput input.CreateCompanyAddressInput
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&createCompanyAddressInput)

	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.companyUseCase.AddAddress(
		params["companyId"],
		createCompanyAddressInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *companyHandler) AddPosition(w http.ResponseWriter, r *http.Request) {

	var createPositionInput input.CreatePositionInput
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&createPositionInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.companyUseCase.AddPosition(
		params["companyId"],
		createPositionInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *companyHandler) UpdatePositionResponsibilities(w http.ResponseWriter, r *http.Request) {

	var updatePositionResponsobilitesInput input.UpdatePositionResponsobilitesInput
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&updatePositionResponsobilitesInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.companyUseCase.UpdatePositionResponsibilities(
		params["companyId"],
		params["positionId"],
		updatePositionResponsobilitesInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *companyHandler) UpdatePositionRequirements(w http.ResponseWriter, r *http.Request) {

	var updatePositionRequirementsInput input.UpdatePositionRequirementsInput
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&updatePositionRequirementsInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.companyUseCase.UpdatePositionRequirements(
		params["companyId"],
		params["positionId"],
		updatePositionRequirementsInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *companyHandler) UpdatePosition(w http.ResponseWriter, r *http.Request) {

	var updatePositionInput input.UpdatePositionInput
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&updatePositionInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.companyUseCase.UpdatePosition(
		params["companyId"],
		params["positionId"],
		updatePositionInput,
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *companyHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.companyUseCase.DeleteAddress(
		params["companyId"],
		params["companyAddressId"],
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *companyHandler) UpdatePositionStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.companyUseCase.UpdatePositionStatus(
		params["companyId"],
		params["positionId"],
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *companyHandler) DeletePosition(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.companyUseCase.DeletePosition(
		params["companyId"],
		params["positionId"],
	)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}
