package api

import (
	"encoding/json"
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/tsurubee/gofetcher/model"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewJsonResponse(w http.ResponseWriter, statusCode int, f model.Fetcher) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(f); err != nil {
		logrus.Fatal(err)
	}
}

func NewHTTPError(w http.ResponseWriter, statusCode int, err error) {
	logrus.Error(err)
	w.WriteHeader(statusCode)

	e := &HTTPError{
		Code:    statusCode,
		Message: err.Error(),
	}
	if err := json.NewEncoder(w).Encode(e); err != nil {
		logrus.Fatal(err)
	}
}
