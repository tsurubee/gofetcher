package main

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/Gurpartap/logrus-stack"
	"github.com/julienschmidt/httprouter"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	stackLevels := []logrus.Level{logrus.PanicLevel, logrus.FatalLevel}
	logrus.AddHook(logrus_stack.NewHook(stackLevels, stackLevels))
}

func main() {
	router := httprouter.New()
	router.GET("/files/keys/:username", FetchAuthorizedKeys)
	logrus.Info("Start Listening ...")
	logrus.Fatal(http.ListenAndServe(":8888", router))
}
