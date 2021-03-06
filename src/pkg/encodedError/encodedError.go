package encodedError

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	BadRequest          = errors.New("invalid payload")
	InternalServerError = errors.New("Internal Server Error")
)

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(getCode(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": "invalid payload",
	})
}

func getCode(err error) int {
	switch err {
	case BadRequest:
		return http.StatusBadRequest
	case InternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
