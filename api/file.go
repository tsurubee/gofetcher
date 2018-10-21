package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tsurubee/gofetcher/model"
	"github.com/tsurubee/gofetcher/config"
)

func UsersHandler(w http.ResponseWriter, r *http.Request, c *config.Config) error {
	vars := mux.Vars(r)

	f, err := model.NewResource(vars, c)
	if err != nil {
		return NewHTTPError(w, http.StatusInternalServerError, err)
	}

	if err = f.Fetch(vars["user"]); err != nil {
		return NewHTTPError(w, http.StatusInternalServerError, err)
	}

	return NewJsonResponse(w, http.StatusOK, f)
}
