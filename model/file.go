package model

import (
	"net/http"
	"io/ioutil"
	"path"
	"github.com/gorilla/mux"
)

var UserFiles = map[string]string{
	"authKey": ".ssh/authorized_keys",
}

type File struct {
	Path     string `json:"path"`
	Output   string `json:"output"`
}

func makeUserFilePath(user, file string) string {
	return path.Join("/home", user, file)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	f, err := NewResource(vars)
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

func (f *File) Fetch(user string) error {
	data, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}
	f.Output = string(data)
	return nil
}
