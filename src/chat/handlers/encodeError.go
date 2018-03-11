package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

// EncodeResponse return response for client
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var code int
	var stateCode string

	switch err {
	default:
		code = http.StatusInternalServerError
		stateCode = "ERR_INTERNAL_SERVER"
	}

	w.WriteHeader(code)

	json.NewEncoder(w).Encode(StandardResponse{
		Message:        err.Error(),
		HTTPStatusCode: code,
		Success:        false,
		StateCode:      stateCode,
	})
}

type errorer interface {
	error() error
}
