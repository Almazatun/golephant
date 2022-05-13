package error_handler

import (
	"net/http"

	error_message "github.com/Almazatun/golephant/pkg/common/error-message"
)

type sentineApiError struct {
	status int
	err    error
}

func (e sentineApiError) APIError() (int, error) {
	return e.status, e.err
}

func ErrorMessageHandler(err error) sentineApiError {
	switch err.Error() {
	case error_message.UNAUTHORIZED:
		return sentineApiError{
			status: http.StatusUnauthorized,
			err:    err,
		}
	case error_message.USER_NOT_FOUND:
		return sentineApiError{
			status: http.StatusNotFound,
			err:    err,
		}
	case error_message.COMPANY_NOT_FOUND:
		return sentineApiError{
			status: http.StatusNotFound,
			err:    err,
		}
	case error_message.RESUME_NOT_FOUND:
		return sentineApiError{
			status: http.StatusNotFound,
			err:    err,
		}
	case error_message.USER_EDUCATION_NOT_FOUND:
		return sentineApiError{
			status: http.StatusNotFound,
			err:    err,
		}
	case error_message.USER_EXPERIENCE_NOT_FOUND:
		return sentineApiError{
			status: http.StatusNotFound,
			err:    err,
		}
	case error_message.COMPANY_ADDRESS_NOT_FOUND:
		return sentineApiError{
			status: http.StatusNotFound,
			err:    err,
		}
	case error_message.BAD_REGUEST:
		return sentineApiError{
			status: http.StatusBadRequest,
			err:    err,
		}
	case error_message.INCCORECT_PASSWORD:
		return sentineApiError{
			status: http.StatusBadRequest,
			err:    err,
		}
	case error_message.INVALID_RESET_PASSWORD_TOKEN:
		return sentineApiError{
			status: http.StatusBadRequest,
			err:    err,
		}
	default:
		return sentineApiError{
			status: 600,
			err:    err,
		}
	}
}
