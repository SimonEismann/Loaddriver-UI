package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"load-director/registry"
	"load-director/shared"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
)

const ()

var (
	sem         = semaphore.NewWeighted(1)
	killChan    = make(chan struct{}, 1)
	killSuccess = make(chan error, 1)
	jobs        []job
	jobQueue    = make(chan job)
	logger      = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
)

type job struct {
	Id            string   `json:"id"`
	Slaves        []string `json:"slaves"`
	ScriptName    string   `json:"scriptName"`
	IntensityFile string   `json:"intensityFile"`
}

func main() {
	log := logger.WithField("func", "main")
	r := mux.NewRouter()
	r.HandleFunc("", handleRoot)
	r.HandleFunc("/", handleRoot)
	r.HandleFunc("/jobs", handleJobPost).Methods(http.MethodPost)
	registry := registry.New()
	go registry.StartCleanUpRoutine()
	go jobConsumer()
	r.Handle("/registry", registry)
	log.Info("Started server")
	http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodDelete, http.MethodGet, http.MethodPut, http.MethodOptions}),
		handlers.AllowCredentials())(r))
}

func handleJobPost(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var postedJob job
	err = json.Unmarshal(b, &postedJob)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postedJob.Id = time.Now().Format("01-02-2006_03:04")
	jobs = append(jobs, postedJob)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("location", fmt.Sprintf("%s/%s", req.URL.Path, postedJob.Id))
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	resp, err := json.MarshalIndent(map[string]string{
		"GET /scripts":             "Lists the available script file names",
		"GET /scripts/{fileName}":  "Returns the script with the specified file name",
		"POST /jobs":               "Queue an experiment with all registered slaves",
		"GET /jobs":                "Lists all jobs",
		"GET /jobs/{jobId}":        "Get the current status of the specified job",
		"GET /jobs/{jobId}/result": "Returns the result of the job",
		"DELETE /api/jobs/current": "Force stop the currently running load generation",
		"GET /registry":            "List all currently registered slaves",
	}, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func handleStop(w http.ResponseWriter, r *http.Request) {
	if sem.TryAcquire(1) {
		sem.Release(1)
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "There is currently no job running!")
		return
	}
	killChan <- struct{}{}
	err := <-killSuccess
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error occured while stopping current job!")
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Job successfully stopped!")
}

func runJob(jobToStart job) error {
	defer sem.Release(1)
	execDir := shared.GetExecDir()
	currentTime := time.Now().Format("01-02-2006_03:04")
	f, err := os.Create(fmt.Sprintf("%s/results%s.csv", execDir, currentTime))
	if err != nil {
		logger.WithField("func", "startLoadDriver").Error("Could not create result file!")
		return err
	}
	defer f.Close()
	cmd := generateCommand(jobToStart.Slaves, jobToStart.ScriptName, jobToStart.IntensityFile)
	done := make(chan error, 1)
	go func() {
		done <- cmd.Run()
	}()
	for {
		select {
		case <-done:
			return nil
		case <-killChan:
			err := cmd.Process.Kill()
			killSuccess <- err
			if err == nil {
				logger.Println("Job successfully stopped!")
				return nil
			}
			logger.WithError(err).Println("Error occured while killing process!")
			return err
		}
	}
}

func jobConsumer() {
	for nextJob := range jobQueue {
		logger.WithField("func", "jobConsumer").Info(fmt.Sprintf("Starting job with ID %s", nextJob.Id))
		err := runJob(nextJob)
		if err != nil {
			logger.WithField("func", "jobConsumer").Error("Job finished with errors")
		}
		logger.WithField("func", "jobConsumer").Info(fmt.Sprintf("Job with ID %s finished", nextJob.Id))
	}
}

func generateCommand(slaves []string, scriptFile string, intensityFile string) *exec.Cmd {
	execDir := shared.GetExecDir()
	commandArgs := []string{
		"-jar", filepath.Join(execDir, "httploadgenerator.jar"),
		"director",
		"-a", fmt.Sprintf("%s/%s", execDir, intensityFile),
		"-l", fmt.Sprintf("%s/%s", execDir, scriptFile)}
	for _, slave := range slaves {
		commandArgs = append(commandArgs, fmt.Sprintf("-s %s", slave))
	}
	return exec.Command(
		"java",
		commandArgs...,
	)
}
