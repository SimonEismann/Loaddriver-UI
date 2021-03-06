package main

import (
	"bytes"
	"io/ioutil"
	"loaddriver-slave/shared"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
)

var (
	registryHost string
	registryPort int
	stopChan     = make(chan struct{}, 1)
	success      = make(chan error, 1)
	sem          = semaphore.NewWeighted(1)
	logger       = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
)

const ()

func initEnvs() {
	registryHost = os.Getenv("REGISTRY_HOST")
	port, err := strconv.Atoi(os.Getenv("REGISTRY_PORT"))
	if err != nil {
		logger.Panicf("Invalid Registry Port: %s", os.Getenv("REGISTRY_PORT"))
	}
	registryPort = port
}

func main() {
	initEnvs()
	r := mux.NewRouter()
	r.HandleFunc("/start", handleStart)
	r.HandleFunc("/stop", handleStop)
	r.HandleFunc("/logs", handleGetLogs)
	go startRegistryHeartbeat()
	http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodDelete, http.MethodGet, http.MethodOptions}),
		handlers.AllowCredentials())(r))
}

func handleGetLogs(w http.ResponseWriter, r *http.Request) {
	logFile, err := ioutil.ReadFile(filepath.Join(shared.GetExecDir(), "logs.txt"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	buffer := bytes.NewBuffer(logFile)
	w.Header().Set("Content-type", "text/plain")
	if _, err := buffer.WriteTo(w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
