package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"loaddriver-master/api"
	"loaddriver-master/consoleEcho"
	log "loaddriver-master/logger"
	"loaddriver-master/pkg/jobs"
	"loaddriver-master/pkg/registry"
	"loaddriver-master/shared"
	"net/http"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type scriptPayload struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

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
	r.HandleFunc("/scripts", handleGetAllScriptNames).Queries("names-only", "true").Methods(http.MethodGet)
	r.HandleFunc("/scripts", handleGetAllScripts).Methods(http.MethodGet)
	r.HandleFunc("/scripts", handlePostScript).Methods(http.MethodPost)
	r.HandleFunc("/intensities", handleGetAllIntensities).Methods(http.MethodGet)
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

func handleGetAllScriptNames(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(filepath.Join(shared.GetExecDir(), "scripts"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var result []scriptPayload
	for _, file := range files {
		result = append(result, scriptPayload{Name: file.Name()})
	}
	resp, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func handleGetAllScripts(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(filepath.Join(shared.GetExecDir(), "scripts"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var result []scriptPayload
	for _, file := range files {
		content, err := ioutil.ReadFile(filepath.Join(shared.GetExecDir(), "scripts", file.Name()))
		if err != nil {
			logger.WithField("func", "handleGetAllScripts").WithError(err).Errorf("Could not read file named %s", file.Name())
		}
		result = append(result, scriptPayload{Name: file.Name(), Content: string(content)})
	}
	resp, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		logger.WithField("func", "handleGetAllScripts").WithError(err).Errorf("Marshalling failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func handlePostScript(w http.ResponseWriter, r *http.Request) {
	postedScript := scriptPayload{}
	err := json.NewDecoder(r.Body).Decode(&postedScript)
	defer r.Body.Close()
	if err != nil {
		logger.WithField("func", "handlePostScript").WithError(err).Error("Could not parse body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(postedScript.Name)
	err = ioutil.WriteFile(filepath.Join(shared.GetExecDir(), "scripts", postedScript.Name), []byte(postedScript.Content), 0777)
	if err != nil {
		logger.WithField("func", "handlePostScript").WithError(err).Error("Could not create file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("location", fmt.Sprintf("%s/%s", r.URL.Path, postedScript.Name))
}

func handleGetAllIntensities(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(filepath.Join(shared.GetExecDir(), "intensities"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}
	resp, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
