package main

import (
	"net/http"
	"os"
	"github.com/sirupsen/logrus"
	"github.com/Gurpartap/logrus-stack"
	"github.com/gorilla/mux"
	"github.com/tsurubee/gofetcher/api"
	"github.com/tsurubee/gofetcher/config"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	stackLevels := []logrus.Level{logrus.PanicLevel, logrus.FatalLevel}
	logrus.AddHook(logrus_stack.NewHook(stackLevels, stackLevels))
}

func main() {
	confFile := os.Getenv("GOFETCHER_CONFIG_PATH")

	if confFile == "" {
		confFile = "./environment/dev/config.toml"
	}

	c, err := config.LoadConfig(confFile)
	if err != nil {
		logrus.Fatal(err)
	}

	r := mux.NewRouter()

	r.Path("/users/{resourceType}/{resourceName}/{user}").
	  HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  	err = api.UsersHandler(w, r, c)
	  	if err != nil {
	  		logrus.Fatal(err)
		}
	})

	logrus.Info("Start Listening on ", c.ListenAddr)
	logrus.Fatal(http.ListenAndServe(c.ListenAddr, r))
}
