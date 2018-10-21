package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tsurubee/gofetcher/model"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	f, err := model.NewResource(vars)
	if err != nil {
		NewHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	if err = f.Fetch(vars["user"]); err != nil {
		NewHTTPError(w, http.StatusInternalServerError, err)
		return
	}

	NewJsonResponse(w, http.StatusOK, f)
}
