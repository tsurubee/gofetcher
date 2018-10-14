package main

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/Gurpartap/logrus-stack"
	"github.com/gorilla/mux"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	stackLevels := []logrus.Level{logrus.PanicLevel, logrus.FatalLevel}
	logrus.AddHook(logrus_stack.NewHook(stackLevels, stackLevels))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/files/{fileType}/{username}", FileHandler)
	logrus.Info("Start Listening ...")
	logrus.Fatal(http.ListenAndServe(":8888", r))
}
