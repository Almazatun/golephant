package middleware

import (
	"errors"
	"net/http"

	"github.com/Almazatun/golephant/pkg/common"
	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
	jwt_gl "github.com/Almazatun/golephant/pkg/jwt_gl"
	logger "github.com/Almazatun/golephant/pkg/logger"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(jwt_gl.HTTP_COOKIE)

		if err != nil {
			if err == http.ErrNoCookie {
				newErr := errors.New(error_message.UNAUTHORIZED)
				common.JSONError(w, newErr, http.StatusUnauthorized)
				return
			}

			logger.Error(err)
			common.JSONError(w, err, http.StatusBadRequest)
			return
		}

		tokenString := cookie.Value

		res, err := jwt_gl.IsValidJWTStr(tokenString)

		if err != nil && !res {
			logger.Error(err)
			common.JSONError(w, err, http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
