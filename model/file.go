package model

import (
	"io/ioutil"
	"path"
)

var UserFiles = map[string]string {
	"authKey": ".ssh/authorized_keys",
}

type File struct {
	Path     string `json:"path"`
	Output   string `json:"output"`
}

func makeUserFilePath(user, file string) string {
	return path.Join("/home", user, file)
}

func (f *File) Fetch(user string) error {
	data, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}
	f.Output = string(data)
	return nil
}
