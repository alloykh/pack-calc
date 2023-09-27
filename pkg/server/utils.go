package server

import (
	"encoding/json"
	"net/http"
)

type HTTPStatusCodeError interface {
	error
	StatusCode() int
}

func WriteHttpStatusCode(w http.ResponseWriter, err error, defaultStatus int) {
	status := defaultStatus
	if errStatus, ok := err.(HTTPStatusCodeError); ok {
		status = errStatus.StatusCode()
	}
	w.WriteHeader(status)
}

type JSONError struct {
	ErrorString string `json:"error"`
}

func (j JSONError) Error() string {
	return j.ErrorString
}

func WrapJSONError(err error) JSONError {
	return JSONError{
		ErrorString: err.Error(),
	}
}

func WriteJSONError(w http.ResponseWriter, err error) {
	jsonErr := WrapJSONError(err)

	if err := json.NewEncoder(w).Encode(jsonErr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("X-Error-Description", err.Error())
		return
	}
}
