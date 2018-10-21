package model

import (
	"net/http"
)

type Command struct {
	Command string `json:"command"`
	Output  string `json:"output"`
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	//ToDo
}

func (c *Command) Fetch() {
	//ToDo
}
