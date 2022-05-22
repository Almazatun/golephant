package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	usecase "github.com/Almazatun/golephant/internal/domain"
	"github.com/Almazatun/golephant/pkg/common"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
	jwt_gl "github.com/Almazatun/golephant/pkg/jwt_gl"
	logger "github.com/Almazatun/golephant/pkg/logger"
	"github.com/dgrijalva/jwt-go"
)

type authHandler struct {
	companyUseCase usecase.CompanyUseCase
	userUseCase    usecase.UserUseCase
}

type AuthHandler interface {
	RegisterCompany(w http.ResponseWriter, r *http.Request)
	LogInCompany(w http.ResponseWriter, r *http.Request)
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LogInUser(w http.ResponseWriter, r *http.Request)
	Me(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(
	companyUseCase usecase.CompanyUseCase,
	userUseCase usecase.UserUseCase,
) AuthHandler {
	return &authHandler{
		companyUseCase: companyUseCase,
		userUseCase:    userUseCase,
	}
}

func (h *authHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registerUserInput input.RegisterUserInput
	err := json.NewDecoder(r.Body).Decode(&registerUserInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.Register(registerUserInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Successfuly register user." + "Email:" + user.Email)
}

func (h *authHandler) LogInUser(w http.ResponseWriter, r *http.Request) {
	var logInUserInput input.LogInUserInput
	err := json.NewDecoder(r.Body).Decode(&logInUserInput)

	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.userUseCase.LogIn(logInUserInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
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

func (h *authHandler) RegisterCompany(w http.ResponseWriter, r *http.Request) {
	var registerCompanyInput input.RegisterCompanyInput
	err := json.NewDecoder(r.Body).Decode(&registerCompanyInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	company, err := h.companyUseCase.Register(registerCompanyInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Successfuly register " + company.Email + "company")
}

func (h *authHandler) LogInCompany(w http.ResponseWriter, r *http.Request) {
	var logInCompanyInput input.LogInCompanyInput
	err := json.NewDecoder(r.Body).Decode(&logInCompanyInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	res, err := h.companyUseCase.LogIn(logInCompanyInput)

	if err != nil {
		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
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

func (h *authHandler) Me(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(jwt_gl.HTTP_COOKIE)

	if err != nil {
		if err == http.ErrNoCookie {
			logger.Error(err)
			newErr := errors.New(error_message.UNAUTHORIZED)
			common.JSONError(w, newErr, http.StatusUnauthorized)
			return
		}

		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value
	claims := &jwt_gl.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwt_gl.JWT_KEY_BYTE, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			newErr := errors.New(error_message.UNAUTHORIZED)
			common.JSONError(w, newErr, http.StatusUnauthorized)
			return
		}

		logger.Error(err)
		common.JSONError(w, err, http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// w.Write([]byte(fmt.Sprintf("You are successfully authorized: =>, %s", claims)))
	json.NewEncoder(w).Encode(claims)

}
