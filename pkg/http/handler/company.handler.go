package handler

import (
	"encoding/json"
	"net/http"

	usecase "github.com/Almazatun/golephant/internal/domain"
	common "github.com/Almazatun/golephant/pkg/common"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
	logger "github.com/Almazatun/golephant/pkg/logger"
)

type companyHandler struct {
	companyUseCase usecase.CompanyUseCase
}

type CompanyHandler interface {
	RegisterCompany(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
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

	json.NewEncoder(w).Encode("Successfuly register" + company.Email + "company")
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
		Name:    common.HTTP_COOKIE,
		Value:   res.Token,
		Expires: res.ExperationTimeJWT,
		Path:    common.SET_COOKIE_PATH,
	}

	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode(res.LogInEntityData)
}
