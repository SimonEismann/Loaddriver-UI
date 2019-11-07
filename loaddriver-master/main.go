package main

import (
	"loaddriver-master/api"
	"loaddriver-master/consoleEcho"
	log "loaddriver-master/logger"
	"loaddriver-master/pkg/jobs"
	"loaddriver-master/pkg/registry"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const ()

var (
	consoleChan   = make(chan []byte)
	slaveRegistry = registry.New()
	logger        = log.NewLogger()
)

func main() {
	r := mux.NewRouter()
	hub := consoleEcho.NewHub(logger, consoleChan)
	jobRunner := jobs.NewJobRunner(r.PathPrefix("/jobs").Subrouter(), consoleChan, slaveRegistry)
	r.HandleFunc("", api.PrintDescription)
	r.HandleFunc("/", api.PrintDescription)
	r.Handle("/jobs/current/output", hub)
	r.Handle("/registry", slaveRegistry)

	// r.HandleFunc("/jobConfigs", handlePostJob).Methods(http.MethodPost)
	// r.HandleFunc("/jobConfigs", handleGetJobs).Methods(http.MethodGet)

	go hub.Run()
	go jobRunner.Start()
	go slaveRegistry.StartCleanUpRoutine()
	logger.WithField("func", "main").Info("Started server")
	err := http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodDelete, http.MethodGet, http.MethodPut, http.MethodOptions}),
		handlers.AllowCredentials())(r))
	if err != nil {
		logger.WithField("func", "main").WithError(err).Error("Error starting server")
	}
}
