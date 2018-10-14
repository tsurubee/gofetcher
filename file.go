package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type File struct {
	FileName string `json:"filename"`
	Contents string `json:"contents"`
}

func FetchAuthorizedKeys(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	f := File{}
	f.FileName = "authorized_keys"

	data, err := ioutil.ReadFile(fmt.Sprintf("/home/%s/.ssh/authorized_keys", username))
	if err != nil {
		logrus.Info(err)
	}
	f.Contents = string(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(f)
}
