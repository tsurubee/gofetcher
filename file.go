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
	Contents string `json:"contents"`
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	f := &File{}
	f.Path = fmt.Sprintf("/home/%s/.ssh/%s", vars["username"], UserFiles[vars["fileType"]])
	f.Open()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(f); err != nil {
		logrus.Info(err)
	}
}

func (f *File) Open() {
	data, err := ioutil.ReadFile(f.Path)
	if err != nil {
		logrus.Info(err)
	}
	f.Contents = string(data)
}
