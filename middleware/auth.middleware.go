package middleware

import (
	"errors"
	"net/http"

	"github.com/Almazatun/golephant/common"
	error_message "github.com/Almazatun/golephant/common/error-message"
	loggerinfo "github.com/Almazatun/golephant/common/loggerInfo"
	handler "github.com/Almazatun/golephant/handler"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(common.HTTP_COOKIE)

		if err != nil {
			if err == http.ErrNoCookie {
				newErr := errors.New(error_message.UNAUTHORIZED)
				handler.HttpResponseBody(w, newErr)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			loggerinfo.LoggerError(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := cookie.Value

		res, err := common.IsValidJWTStr(tokenString)

		if err != nil && !res {
			loggerinfo.LoggerError(err)
			handler.HttpResponseBody(w, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
