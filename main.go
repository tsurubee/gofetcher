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
	confFile := "./config.toml"
	c, err := loadConfig(confFile)
	if err != nil {
		logrus.Fatal(err)
	}

	r := mux.NewRouter()

	r.Path("/users/{resourceType}/{resourceName}/{user}").
	  HandlerFunc(UsersHandler)

	logrus.Info("Start Listening on ", c.ListenAddr)
	logrus.Fatal(http.ListenAndServe(c.ListenAddr, r))
}
