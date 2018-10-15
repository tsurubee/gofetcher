package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"fmt"
)

var UserFiles = map[string]string{
	"authKey": "authorized_keys",
}

type File struct {
	Path     string `json:"path"`
	Output   string `json:"output"`
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := r.FormValue("user")
	f := &File{}
	f.Path = fmt.Sprintf("/home/%s/.ssh/%s", user, UserFiles[vars["fileType"]])
	f.Fetch()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(f); err != nil {
		logrus.Info(err)
	}
}

func (f *File) Fetch() {
	data, err := ioutil.ReadFile(f.Path)
	if err != nil {
		logrus.Info(err)
	}
	f.Output = string(data)
}
