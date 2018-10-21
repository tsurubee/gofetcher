package main

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"path"
	"github.com/sirupsen/logrus"
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
		//ToDo
		logrus.Info(err)
	}

	err = f.Fetch(vars["user"])
	if err != nil {
		//ToDo
		logrus.Info(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(f); err != nil {
		logrus.Info(err)
	}
}

func (f *File) Fetch(user string) error {
	data, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}
	f.Output = string(data)
	return nil
}
