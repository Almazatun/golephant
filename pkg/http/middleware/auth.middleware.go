package middleware

import (
	"errors"
	"net/http"

	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	handler "github.com/Almazatun/golephant/pkg/http/handler"
	jwt_gl "github.com/Almazatun/golephant/pkg/jwt_gl"
	logger "github.com/Almazatun/golephant/pkg/logger"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(jwt_gl.HTTP_COOKIE)

		if err != nil {
			if err == http.ErrNoCookie {
				newErr := errors.New(error_message.UNAUTHORIZED)
				handler.HttpResponseBody(w, newErr)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			logger.InfoError(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := cookie.Value

		res, err := jwt_gl.IsValidJWTStr(tokenString)

		if err != nil && !res {
			logger.InfoError(err)
			handler.HttpResponseBody(w, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
