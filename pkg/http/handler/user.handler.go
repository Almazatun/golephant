package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	usecase "github.com/Almazatun/golephant/internal/domain"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	"github.com/Almazatun/golephant/pkg/http/presentation/input"
	jwt_gl "github.com/Almazatun/golephant/pkg/jwt_gl"
	logger "github.com/Almazatun/golephant/pkg/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

type UserHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
	AuthMe(w http.ResponseWriter, r *http.Request)
	UpdateUserData(w http.ResponseWriter, r *http.Request)
	GetLinkResetPassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

func (h *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registerUserInput input.RegisterUserInput
	err := json.NewDecoder(r.Body).Decode(&registerUserInput)

	if err != nil {
		logger.InfoError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.Register(registerUserInput)

	if err != nil {
		logger.InfoError(err)
		HttpResponseBody(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Successfuly register user." + "Email:" + user.Email)
}

func (h *userHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var logInUserInput input.LogInUserInput
	err := json.NewDecoder(r.Body).Decode(&logInUserInput)

	if err != nil {
		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.userUseCase.LogIn(logInUserInput)

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

func (h *userHandler) AuthMe(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(jwt_gl.HTTP_COOKIE)

	if err != nil {
		if err == http.ErrNoCookie {
			newErr := errors.New(error_message.UNAUTHORIZED)
			HttpResponseBody(w, newErr)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
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
			HttpResponseBody(w, newErr)

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// w.Write([]byte(fmt.Sprintf("You are successfully authorized: =>, %s", claims)))
	json.NewEncoder(w).Encode(claims)

}

func (h *userHandler) UpdateUserData(w http.ResponseWriter, r *http.Request) {
	var updateUserDataInput input.UpdateUserDataInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updateUserDataInput)

	if err != nil {
		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.userUseCase.UpdateData(params["userId"], updateUserDataInput)

	if err != nil {
		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *userHandler) GetLinkResetPassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := h.userUseCase.GetLinkResetPassword(params["userId"])

	if err != nil {
		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (h *userHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var resetUserPasswordInput input.ResetUserPasswordInput
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&resetUserPasswordInput)

	if err != nil {
		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.userUseCase.ResetPassword(
		params["userId"],
		params["token"],
		resetUserPasswordInput.Password,
	)

	if err != nil {
		logger.InfoError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func HelloWord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func HttpResponseBody(w http.ResponseWriter, err error) {
	resp := make(map[string]string)

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	resp["message"] = err.Error()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
