package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type QueryResolverFunc[I, O any] func(context.Context, I) (O, error)

var ErrExpectedPostRequest = errors.New("expected POST method with body")

func HandleJSONPost[I, O any](fn QueryResolverFunc[I, O]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var input I

		if r.Method != http.MethodPost || r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			WriteJSONError(w, ErrExpectedPostRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			WriteJSONError(w, fmt.Errorf("failed to decode request: %w", err))
			return
		}

		if err := Validate(input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			WriteJSONError(w, fmt.Errorf("request validation failed: %w", err))
			return
		}

		output, err := fn(r.Context(), input)
		if err != nil {
			WriteHttpStatusCode(w, err, http.StatusInternalServerError)
			WriteJSONError(w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(output); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			WriteJSONError(w, fmt.Errorf("failed to encode response: %w", err))
			return
		}
	}
}
