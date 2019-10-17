package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"loaddriver-master/api"
	"loaddriver-master/consoleEcho"
	log "loaddriver-master/logger"
	"loaddriver-master/registry"
	"loaddriver-master/shared"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const ()

var (
	consoleChan   = make(chan []byte, shared.MessageBufferSize)
	killChan      = make(chan struct{}, 1)
	killSuccess   = make(chan error, 1)
	jobsMap       = make(map[string]job)
	jobQueue      = make(chan job)
	jobsDone      []job
	slaveRegistry = registry.New()
	logger        = log.NewLogger()
)

type job struct {
	Id             string   `json:"id"`
	Slaves         []string `json:"slaves"`
	ScriptName     string   `json:"scriptName"`
	IntensityFile  string   `json:"intensityFile"`
	WarmupDuration int      `json:"warmupDuration"`
	WarmupPause    int      `json:"warmupPause"`
	WarmupRate     float64  `json:"warmupRate"`
	Threads        int      `json:"threads"`
	Timeout        int      `json:"timeout"`
}

func NewDefaultJob() job {
	var allSlaves []string
	for _, slave := range slaveRegistry.Slaves {
		allSlaves = append(allSlaves, string(slave.Location))
	}
	return job{
		Id:             time.Now().Format("02-01-2006_15:04:05"),
		Slaves:         allSlaves,
		ScriptName:     "teastore_browse.lua",
		IntensityFile:  "defaultIntensity.csv",
		WarmupDuration: 30,
		WarmupPause:    5,
		WarmupRate:     0.0,
		Threads:        128,
		Timeout:        0,
	}
}

func main() {
	hub := consoleEcho.NewHub(logger, consoleChan)
	r := mux.NewRouter()
	r.HandleFunc("", api.PrintDescription)
	r.HandleFunc("/", api.PrintDescription)
	r.HandleFunc("/jobs", handlePostJob).Methods(http.MethodPost)
	r.HandleFunc("/jobs", handleGetJobs).Methods(http.MethodGet)
	r.HandleFunc("/jobs/done", handleGetJobsDone).Methods(http.MethodGet)
	r.HandleFunc("/jobs/default", handlePostJobDefault).Methods(http.MethodPost)
	r.HandleFunc("/jobs/current", handleDeleteJobCurrent).Methods(http.MethodDelete)
	r.HandleFunc("/jobs/{jobId}", handleGetJobByID).Methods(http.MethodGet)
	r.HandleFunc("/jobs/{jobId}/log", handleGetJobLog).Methods(http.MethodGet)
	r.HandleFunc("/jobs/{jobId}/result", handleGetJobResults).Methods(http.MethodGet)
	r.Handle("/jobs/current/output", hub)
	r.Handle("/registry", slaveRegistry)
	go hub.Run()
	go jobConsumer()
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

func handleGetJobs(w http.ResponseWriter, req *http.Request) {
	var jobs []job
	for _, v := range jobsMap {
		jobs = append(jobs, v)
	}
	resp, err := json.MarshalIndent(jobs, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func handleGetJobByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobId := vars["jobId"]
	if _, ok := jobsMap[jobId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, err := json.MarshalIndent(jobsMap[jobId], "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func handleGetJobLog(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobId := vars["jobId"]
	if _, ok := jobsMap[jobId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	logFile, err := ioutil.ReadFile(filepath.Join(shared.GetExecDir(), "logs", jobId))
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

func handleGetJobResults(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobId := vars["jobId"]
	if _, ok := jobsMap[jobId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resultsFile, err := ioutil.ReadFile(filepath.Join(shared.GetExecDir(), "results", fmt.Sprintf("%s.csv", jobId)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	buffer := bytes.NewBuffer(resultsFile)
	w.Header().Set("Content-type", "text/csv")
	if _, err := buffer.WriteTo(w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handleGetJobsDone(w http.ResponseWriter, req *http.Request) {
	resp, err := json.MarshalIndent(jobsDone, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func handlePostJob(w http.ResponseWriter, req *http.Request) {
	postedJob := NewDefaultJob()
	err := json.NewDecoder(req.Body).Decode(&postedJob)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jobsMap[postedJob.Id] = postedJob
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("location", fmt.Sprintf("%s/%s", req.URL.Path, postedJob.Id))
}

func handlePostJobDefault(w http.ResponseWriter, req *http.Request) {
	newJob := NewDefaultJob()
	jobsMap[newJob.Id] = newJob
	jobQueue <- newJob
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("location", fmt.Sprintf("%s/%s", req.URL.Path, newJob.Id))
}

func handleDeleteJobCurrent(w http.ResponseWriter, r *http.Request) {
	killChan <- struct{}{}
	select {
	case err := <-killSuccess:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error occured while stopping current job")
			return
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Job successfully stopped")
	case <-time.After(10 * time.Second):
		<-killChan
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "There is currently no job running")
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
	resultsFileName := fmt.Sprintf("%s.csv", jobToStart.Id)
	f, err := shared.CreateFile(filepath.Join(shared.GetExecDir(), "results"), resultsFileName)
	if err != nil {
		logger.WithField("func", "runJob").WithError(err).Error("Could not create result file")
		return err
	}
	f.Close()
	f, err = shared.CreateFile(filepath.Join(shared.GetExecDir(), "logs"), jobToStart.Id)
	if err != nil {
		logger.WithField("func", "runJob").WithError(err).Error("Could not create log file")
		return err
	}
	defer f.Close()
	cmd := generateCommand(jobToStart, filepath.Join("results", resultsFileName), f)
	successfulStarts := startSlaves(jobToStart.Slaves)
	defer stopSlaves(successfulStarts)
	done := make(chan error, 1)
	go func() {
		done <- cmd.Run()
	}()
	select {
	case err = <-done:
		if err != nil {
			logger.WithField("func", "runJob").WithError(err).Error("Error occured during job run")
			return err
		}
		return nil
	case <-killChan:
		err := cmd.Process.Kill()
		killSuccess <- err
		if err == nil {
			logger.Println("Job successfully stopped")
			return nil
		}
		logger.WithError(err).Println("Error occured while killing process")
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
		jobsDone = append(jobsDone, nextJob)
	}
}

func generateCommand(forJob job, resultsFile string, logFile *os.File) *exec.Cmd {
	logger.Info(resultsFile)
	execDir := shared.GetExecDir()
	commandArgs := []string{
		"-jar", filepath.Join(execDir, "httploadgenerator.jar"),
		"director",
		"-o", fmt.Sprintf("../%s", resultsFile),
		"-a", filepath.Join(execDir, "intensities", forJob.IntensityFile),
		"-l", filepath.Join(execDir, "scripts", forJob.ScriptName),
		"--wd", strconv.Itoa(forJob.WarmupDuration),
		"--wp", strconv.Itoa(forJob.WarmupPause),
		"--wr", fmt.Sprintf("%f", forJob.WarmupRate),
		"-t", strconv.Itoa(forJob.Threads),
		"-u", strconv.Itoa(forJob.Timeout),
	}
	for _, slave := range forJob.Slaves {
		commandArgs = append(commandArgs, fmt.Sprintf("-s %s", slave))
	}
	result := exec.Command(
		"java",
		commandArgs...,
	)
	channelWriter := consoleEcho.NewChannelWriter(consoleChan)
	result.Stdout = io.MultiWriter(os.Stdout, channelWriter, logFile)
	result.Stderr = io.MultiWriter(os.Stderr, channelWriter, logFile)
	return result
}
