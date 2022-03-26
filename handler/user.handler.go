package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	common "github.com/Almazatun/golephant/common"
	error_message "github.com/Almazatun/golephant/common/error-message"
	loggerinfo "github.com/Almazatun/golephant/common/loggerInfo"
	usecase "github.com/Almazatun/golephant/domain"
	"github.com/Almazatun/golephant/infrastucture/entity"
	"github.com/Almazatun/golephant/presentation/input"
	"github.com/dgrijalva/jwt-go"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

type UserHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
	AuthMe(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

func (h *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registerUserInput *entity.User
	err := json.NewDecoder(r.Body).Decode(&registerUserInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userUseCase.RegisterUser(registerUserInput)

	if err != nil {
		loggerinfo.LoggerError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *userHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	var logInInput input.LogIn
	err := json.NewDecoder(r.Body).Decode(&logInInput)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.userUseCase.LogIn(logInInput)

	if err != nil {
		httpResponseBody(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   res.Token,
			Expires: res.ExperationTimeJWT,
		})

	json.NewEncoder(w).Encode("Successfuly log in")

}

func (h *userHandler) AuthMe(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			newErr := errors.New(error_message.UNAUTHORIZED)
			httpResponseBody(w, newErr)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := cookie.Value
	claims := &common.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return common.JWT_KEY_BYTE, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			newErr := errors.New(error_message.UNAUTHORIZED)
			httpResponseBody(w, newErr)

			w.WriteHeader(http.StatusUnauthorized)
			return
		}
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

func HelloWord(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func httpResponseBody(w http.ResponseWriter, err error) {
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
