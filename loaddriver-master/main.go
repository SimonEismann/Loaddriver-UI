package main

import (
	"encoding/json"
	"fmt"
	"io"
	"loaddriver-master/api"
	"loaddriver-master/registry"
	"loaddriver-master/shared"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const ()

var (
	killChan      = make(chan struct{}, 1)
	killSuccess   = make(chan error, 1)
	jobs          []job
	jobQueue      = make(chan job)
	slaveRegistry = registry.New()
	logger        = &logrus.Logger{
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
	r.HandleFunc("", api.PrintDescription)
	r.HandleFunc("/", api.PrintDescription)
	r.HandleFunc("/jobs/current", handleStop).Methods(http.MethodDelete)
	r.HandleFunc("/jobs/default", handleJobPostDefault).Methods(http.MethodPost)
	r.HandleFunc("/jobs", handleJobPost).Methods(http.MethodPost)
	r.HandleFunc("/jobs", handleJobsGet).Methods(http.MethodGet)
	go slaveRegistry.StartCleanUpRoutine()
	go jobConsumer()
	r.Handle("/registry", slaveRegistry)
	log.Info("Started server")
	http.ListenAndServe(":80", handlers.CORS(handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodDelete, http.MethodGet, http.MethodPut, http.MethodOptions}),
		handlers.AllowCredentials())(r))
}

func handleJobsGet(w http.ResponseWriter, req *http.Request) {
	resp, err := json.MarshalIndent(jobs, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func handleJobPost(w http.ResponseWriter, req *http.Request) {
	var postedJob job
	err := json.NewDecoder(req.Body).Decode(&postedJob)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	postedJob.Id = time.Now().Format("01-02-2006_03:04:05")
	jobs = append(jobs, postedJob)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("location", fmt.Sprintf("%s/%s", req.URL.Path, postedJob.Id))
}

func handleJobPostDefault(w http.ResponseWriter, req *http.Request) {
	var allSlaves []string
	for _, slave := range slaveRegistry.Slaves {
		allSlaves = append(allSlaves, string(slave.Location))
	}
	newJob := job{
		Id:            time.Now().Format("01-02-2006 03:04:05"),
		Slaves:        allSlaves,
		ScriptName:    "teastore_browse.lua",
		IntensityFile: "defaultIntensity.csv",
	}
	jobs = append(jobs, newJob)
	jobQueue <- newJob
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("location", fmt.Sprintf("%s/%s", req.URL.Path, newJob.Id))
}

func handleStop(w http.ResponseWriter, r *http.Request) {
	killChan <- struct{}{}
	select {
	case err := <-killSuccess:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error occured while stopping current job!")
			return
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Job successfully stopped!")
	case <-time.After(10 * time.Second):
		<-killChan
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "There is currently no job running!")
		return
	}
}

func startSlaves(slaves []string) []string {
	var successfulStarts []string
	for _, slave := range slaves {
		resp, err := http.Get(fmt.Sprintf("http://%s/start", slave))
		if err != nil || resp.StatusCode != http.StatusOK {
			logger.WithField("func", "runJob").Warn(fmt.Sprintf("Slave registered under %s could not be started, ignoring for this job...", slave))
			continue
		}
		successfulStarts = append(successfulStarts, slave)
	}
	time.Sleep(5000)
	return successfulStarts
}

func stopSlaves(slaves []string) {
	for _, slave := range slaves {
		resp, err := http.Get(fmt.Sprintf("http://%s/stop", slave))
		if err != nil || resp.StatusCode != http.StatusOK {
			logger.WithField("func", "runJob").Warn(fmt.Sprintf("Error occured while stopping slave registered under %s ...", slave))
		}
	}
}

func runJob(jobToStart job) error {
	f, err := os.Create(fmt.Sprintf("%s/results%s.csv", shared.GetExecDir(), time.Now().Format("01-02-2006_03:04")))
	if err != nil {
		logger.WithField("func", "runJob").Error("Could not create result file!")
		return err
	}
	defer f.Close()
	cmd := generateCommand(jobToStart)
	successfulStarts := startSlaves(jobToStart.Slaves)
	defer stopSlaves(successfulStarts)
	done := make(chan error, 1)
	go func() {
		done <- cmd.Run()
	}()
	select {
	case err = <-done:
		if err != nil {
			logger.WithField("func", "runJob").WithError(err).Error("Error occured during job run!")
			return err
		}
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

func jobConsumer() {
	for nextJob := range jobQueue {
		logger.WithField("func", "jobConsumer").Info(fmt.Sprintf("Starting job with ID %s", nextJob.Id))
		err := runJob(nextJob)
		if err != nil {
			logger.WithField("func", "jobConsumer").Error("Job finished with errors")
		} else {
			logger.WithField("func", "jobConsumer").Info(fmt.Sprintf("Job with ID %s finished", nextJob.Id))
		}
		jobs = jobs[1:]
	}
}

func generateCommand(forJob job) *exec.Cmd {
	execDir := shared.GetExecDir()
	commandArgs := []string{
		"-jar", filepath.Join(execDir, "httploadgenerator.jar"),
		"director",
		"-a", fmt.Sprintf("%s/%s", execDir, forJob.IntensityFile),
		"-l", fmt.Sprintf("%s/%s", execDir, forJob.ScriptName)}
	for _, slave := range forJob.Slaves {
		commandArgs = append(commandArgs, fmt.Sprintf("-s %s", slave))
	}
	result := exec.Command(
		"java",
		commandArgs...,
	)
	result.Stdout = os.Stdout
	result.Stderr = os.Stderr
	return result
}
