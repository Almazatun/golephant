package common

import (
	"encoding/json"
	"log"
	"net/http"

	error_handler "github.com/Almazatun/golephant/pkg/common/error-handler"
)

func JSONError(w http.ResponseWriter, err error, status int) {
	apiErr := error_handler.ErrorMessageHandler(err)
	statusCode, err := apiErr.APIError()

	if statusCode == 600 {
		statusCode = status
	}

	resp := make(map[string]string)

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	resp["error"] = err.Error()
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
