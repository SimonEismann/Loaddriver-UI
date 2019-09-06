package main

import (
	"load-director/logger"
	"load-director/registry"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const ()

func main() {
	log := logger.NewLogger().WithField("method", "main")
	r := mux.NewRouter()
	registry := registry.New()
	go registry.StartCleanUpRoutine()
	r.Handle("/registry", registry)
	log.Info("Started server")
	http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodDelete, http.MethodGet, http.MethodPut, http.MethodOptions}),
		handlers.AllowCredentials())(r))
}
