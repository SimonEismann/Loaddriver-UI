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
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type filePayload struct {
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
	r.HandleFunc("/scripts/{fileName}", handleDeleteScript).Methods(http.MethodDelete)
	r.HandleFunc("/intensities", handleGetAllIntensityFileNames).Queries("names-only", "true").Methods(http.MethodGet)
	r.HandleFunc("/intensities", handleGetAllIntensities).Methods(http.MethodGet)
	r.HandleFunc("/intensities", handlePostIntensity).Methods(http.MethodPost)
	r.HandleFunc("/intensities/{fileName}", handleDeleteIntensity).Methods(http.MethodDelete)
	r.PathPrefix("/intensities/").Handler(http.StripPrefix("/intensities/", http.FileServer(http.Dir("./intensities/"))))
	r.PathPrefix("/scripts/").Handler(http.StripPrefix("/scripts/", http.FileServer(http.Dir("./scripts/"))))
	go hub.Run()
	go jobRunner.Start()
	go slaveRegistry.StartCleanUpRoutine()
	logger.WithField("func", "main").Info("Started server")
	err := http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodDelete, http.MethodGet, http.MethodPut, http.MethodOptions}),
		handlers.AllowCredentials())(handlers.LoggingHandler(os.Stdout, r)))
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
	var result []filePayload
	for _, file := range files {
		result = append(result, filePayload{Name: file.Name()})
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
	result := make([]filePayload, 0)
	for _, file := range files {
		content, err := ioutil.ReadFile(filepath.Join(shared.GetExecDir(), "scripts", file.Name()))
		if err != nil {
			logger.WithField("func", "handleGetAllScripts").WithError(err).Errorf("Could not read file named %s", file.Name())
		}
		result = append(result, filePayload{Name: file.Name(), Content: string(content)})
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
	postedScript := filePayload{}
	err := json.NewDecoder(r.Body).Decode(&postedScript)
	defer r.Body.Close()
	if err != nil {
		logger.WithField("func", "handlePostScript").WithError(err).Error("Could not parse body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ioutil.WriteFile(filepath.Join(shared.GetExecDir(), "scripts", postedScript.Name), []byte(postedScript.Content), 0666)
	if err != nil {
		logger.WithField("func", "handlePostScript").WithError(err).Error("Could not create file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("location", fmt.Sprintf("%s/%s", r.URL.Path, postedScript.Name))
}

func handleDeleteScript(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	var _, err = os.Stat(filepath.Join(shared.GetExecDir(), "scripts", fileName))
	if os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = os.Remove(filepath.Join(shared.GetExecDir(), "scripts", fileName))
	if err != nil {
		logger.WithField("func", "handleDeleteScript").WithError(err).Error("Could not delete file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handleGetAllIntensityFileNames(w http.ResponseWriter, r *http.Request) {
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
		logger.WithField("func", "handleGetAllIntensities").WithError(err).Errorf("Marshalling failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func handleGetAllIntensities(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(filepath.Join(shared.GetExecDir(), "intensities"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result := make([]filePayload, 0)
	for _, file := range files {
		content, err := ioutil.ReadFile(filepath.Join(shared.GetExecDir(), "intensities", file.Name()))
		if err != nil {
			logger.WithField("func", "handleGetAllIntensities").WithError(err).Errorf("Could not read file named %s", file.Name())
		}
		result = append(result, filePayload{Name: file.Name(), Content: string(content)})
	}
	resp, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		logger.WithField("func", "handleGetAllIntensities").WithError(err).Errorf("Marshalling failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func handlePostIntensity(w http.ResponseWriter, r *http.Request) {
	postedIntensity := filePayload{}
	err := json.NewDecoder(r.Body).Decode(&postedIntensity)
	defer r.Body.Close()
	if err != nil {
		logger.WithField("func", "handlePostIntensity").WithError(err).Error("Could not parse body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = ioutil.WriteFile(filepath.Join(shared.GetExecDir(), "intensities", postedIntensity.Name), []byte(postedIntensity.Content), 0666)
	if err != nil {
		logger.WithField("func", "handlePostIntensity").WithError(err).Error("Could not create file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("location", fmt.Sprintf("%s/%s", r.URL.Path, postedIntensity.Name))
}

func handleDeleteIntensity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	var _, err = os.Stat(filepath.Join(shared.GetExecDir(), "intensities", fileName))
	if os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = os.Remove(filepath.Join(shared.GetExecDir(), "intensities", fileName))
	if err != nil {
		logger.WithField("func", "handleDeleteIntensity").WithError(err).Error("Could not delete file")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
