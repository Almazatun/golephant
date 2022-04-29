package handler

import (
	"encoding/json"
	"net/http"

	usecase "github.com/Almazatun/golephant/internal/domain"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
	jwt_gl "github.com/Almazatun/golephant/pkg/jwt_gl"
	logger "github.com/Almazatun/golephant/pkg/logger"
	"github.com/gorilla/mux"
)

type companyHandler struct {
	companyUseCase usecase.CompanyUseCase
}

type CompanyHandler interface {
	RegisterCompany(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
	AddCompanyAddress(w http.ResponseWriter, r *http.Request)
	DeleteCompanyAddress(w http.ResponseWriter, r *http.Request)
}

func NewCompanyHandler(
	companyUseCase usecase.CompanyUseCase,
) CompanyHandler {
	return &companyHandler{
		companyUseCase: companyUseCase,
	}
}

func (h *companyHandler) RegisterCompany(w http.ResponseWriter, r *http.Request) {
	var registerCompanyInput input.RegisterCompanyInput
	err := json.NewDecoder(r.Body).Decode(&registerCompanyInput)

	if err != nil {
		logger.InfoError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	company, err := h.companyUseCase.RegisterCompany(registerCompanyInput)

	if err != nil {
		logger.InfoError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Successfuly register " + company.Email + "company")
}

func (h *companyHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var logInCompanyInput input.LogInCompanyInput
	err := json.NewDecoder(r.Body).Decode(&logInCompanyInput)

	if err != nil {
		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.companyUseCase.LogIn(logInCompanyInput)

	if err != nil {
		logger.InfoError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:    jwt_gl.HTTP_COOKIE,
		Value:   res.Token,
		Expires: res.ExperationTimeJWT,
		Path:    jwt_gl.SET_COOKIE_PATH,
	}

	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode(res.LogInEntityData)
}

func (h *companyHandler) AddCompanyAddress(w http.ResponseWriter, r *http.Request) {
	var createCompanyAddressInput input.CreateCompanyAddressInput
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&createCompanyAddressInput)

	if err != nil {
		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.companyUseCase.AddCompanyAddress(
		params["companyId"],
		createCompanyAddressInput,
	)

	if err != nil {
		logger.InfoError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *companyHandler) DeleteCompanyAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.companyUseCase.DeleteCompanyAddress(
		params["companyId"],
		params["companyAddressId"],
	)

	if err != nil {
		logger.InfoError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}
