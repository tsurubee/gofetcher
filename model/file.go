package model

import (
	"os/user"
	"io/ioutil"
	"path"
)

var UserFiles = map[string]string{
	"authKey": ".ssh/authorized_keys",
}

type File struct {
	Path     string `json:"path"`
	Output   string `json:"output"`
}

func makeUserFilePath(username, file string) (string, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return "", err
	}

	return path.Join(u.HomeDir, file), nil
}

func (f *File) Fetch(user string) error {
	data, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}
	f.Output = string(data)
	return nil
}
