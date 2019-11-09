package jobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"loaddriver-master/consoleEcho"
	"loaddriver-master/logger"
	"loaddriver-master/pkg/registry"
	"loaddriver-master/shared"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

func NewDefaultJob(slaves []string) job {
	return job{
		Id:             time.Now().Format("02-01-2006_15:04:05"),
		Slaves:         slaves,
		ScriptName:     "teastore_browse.lua",
		IntensityFile:  "defaultIntensity.csv",
		WarmupDuration: 30,
		WarmupPause:    5,
		WarmupRate:     0.0,
		Threads:        128,
		Timeout:        0,
	}
}

type jobRunner struct {
	Router        *mux.Router
	logger        *logger.StandardLogger
	consoleChan   chan []byte
	killChan      chan struct{}
	killSuccess   chan error
	jobsMap       map[string]job
	jobsDone      []job
	nextJobChan   chan job
	slaveRegistry *registry.Registry
}

func NewJobRunner(r *mux.Router, consoleChan chan []byte, slaveRegistry *registry.Registry) *jobRunner {
	result := &jobRunner{
		Router:        r,
		logger:        logger.NewLogger(),
		consoleChan:   consoleChan,
		killChan:      make(chan struct{}),
		killSuccess:   make(chan error),
		jobsMap:       make(map[string]job),
		nextJobChan:   make(chan job),
		jobsDone:      make([]job, 0),
		slaveRegistry: slaveRegistry,
	}
	r.HandleFunc("", result.handlePostJob).Methods(http.MethodPost)
	r.HandleFunc("", result.handleGetJobs).Methods(http.MethodGet)
	r.HandleFunc("/done", result.handleGetJobsDone).Methods(http.MethodGet)
	r.HandleFunc("/default", result.handlePostJobDefault).Methods(http.MethodPost)
	r.HandleFunc("/current", result.handleDeleteJobCurrent).Methods(http.MethodDelete)
	r.HandleFunc("/{jobId}", result.handleGetJobByID).Methods(http.MethodGet)
	r.HandleFunc("/{jobId}/log", result.handleGetJobLog).Methods(http.MethodGet)
	r.HandleFunc("/{jobId}/result", result.handleGetJobResults).Methods(http.MethodGet)
	return result
}

func (jq *jobRunner) Start() {
	for nextJob := range jq.nextJobChan {
		jq.logger.WithField("func", "jobConsumer").Info(fmt.Sprintf("Starting job with ID %s", nextJob.Id))
		err := jq.runJob(nextJob)
		if err != nil {
			jq.logger.WithField("func", "jobConsumer").Error("Job finished with errors")
		} else {
			jq.logger.WithField("func", "jobConsumer").Info(fmt.Sprintf("Job with ID %s finished", nextJob.Id))
		}
		jq.jobsDone = append(jq.jobsDone, nextJob)
	}
}

func (jq *jobRunner) handleGetJobs(w http.ResponseWriter, req *http.Request) {
	var jobs []job
	for _, v := range jq.jobsMap {
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

func (jq *jobRunner) handleGetJobByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobId := vars["jobId"]
	if _, ok := jq.jobsMap[jobId]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, err := json.MarshalIndent(jq.jobsMap[jobId], "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (jq *jobRunner) handleGetJobLog(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobId := vars["jobId"]
	if _, ok := jq.jobsMap[jobId]; !ok {
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

func (jq *jobRunner) handleGetJobResults(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	jobId := vars["jobId"]
	if _, ok := jq.jobsMap[jobId]; !ok {
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

func (jq *jobRunner) handleGetJobsDone(w http.ResponseWriter, req *http.Request) {
	resp, err := json.MarshalIndent(jq.jobsDone, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (jq *jobRunner) handlePostJob(w http.ResponseWriter, req *http.Request) {
	var allSlaves []string
	for _, slave := range jq.slaveRegistry.Slaves {
		allSlaves = append(allSlaves, string(slave.Location))
	}
	postedJob := NewDefaultJob(allSlaves)
	err := json.NewDecoder(req.Body).Decode(&postedJob)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jq.jobsMap[postedJob.Id] = postedJob
	jq.nextJobChan <- postedJob
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("location", fmt.Sprintf("%s/%s", req.URL.Path, postedJob.Id))
}

func (jq *jobRunner) handlePostJobDefault(w http.ResponseWriter, req *http.Request) {
	var allSlaves []string
	for _, slave := range jq.slaveRegistry.Slaves {
		allSlaves = append(allSlaves, string(slave.Location))
	}
	newJob := NewDefaultJob(allSlaves)
	jq.jobsMap[newJob.Id] = newJob
	jq.nextJobChan <- newJob
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("location", fmt.Sprintf("%s/%s", req.URL.Path, newJob.Id))
}

func (jq *jobRunner) handleDeleteJobCurrent(w http.ResponseWriter, r *http.Request) {
	jq.killChan <- struct{}{}
	select {
	case err := <-jq.killSuccess:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Error occured while stopping current job")
			return
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Job successfully stopped")
	case <-time.After(10 * time.Second):
		<-jq.killChan
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "There is currently no job running")
		return
	}
}

func (jq *jobRunner) startSlaves(slaves []string) []string {
	var successfulStarts []string
	for _, slave := range slaves {
		resp, err := http.Get(fmt.Sprintf("http://%s/start", slave))
		if err != nil || resp.StatusCode != http.StatusOK {
			jq.logger.WithField("func", "runJob").Warn(fmt.Sprintf("Slave registered under %s could not be started, ignoring for this job...", slave))
			continue
		}
		successfulStarts = append(successfulStarts, slave)
	}
	time.Sleep(5000)
	return successfulStarts
}

func (jq *jobRunner) stopSlaves(slaves []string) {
	for _, slave := range slaves {
		resp, err := http.Get(fmt.Sprintf("http://%s/stop", slave))
		if err != nil || resp.StatusCode != http.StatusOK {
			jq.logger.WithField("func", "runJob").Warn(fmt.Sprintf("Error occured while stopping slave registered under %s ...", slave))
		}
	}
}

func (jq *jobRunner) runJob(jobToStart job) error {
	resultsFileName := fmt.Sprintf("%s.csv", jobToStart.Id)
	f, err := shared.CreateFile(filepath.Join(shared.GetExecDir(), "results"), resultsFileName)
	if err != nil {
		jq.logger.WithField("func", "runJob").WithError(err).Error("Could not create result file")
		return err
	}
	f.Close()
	f, err = shared.CreateFile(filepath.Join(shared.GetExecDir(), "logs"), jobToStart.Id)
	if err != nil {
		jq.logger.WithField("func", "runJob").WithError(err).Error("Could not create log file")
		return err
	}
	defer f.Close()
	cmd := jq.generateCommand(jobToStart, filepath.Join("results", resultsFileName), f)
	successfulStarts := jq.startSlaves(jobToStart.Slaves)
	defer jq.stopSlaves(successfulStarts)
	done := make(chan error, 1)
	go func() {
		done <- cmd.Run()
	}()
	select {
	case err = <-done:
		if err != nil {
			jq.logger.WithField("func", "runJob").WithError(err).Error("Error occured during job run")
			return err
		}
		return nil
	case <-jq.killChan:
		err := cmd.Process.Kill()
		jq.killSuccess <- err
		if err == nil {
			jq.logger.Println("Job successfully stopped")
			return nil
		}
		jq.logger.WithError(err).Println("Error occured while killing process")
		return err
	}
}

func (jq *jobRunner) generateCommand(forJob job, resultsFile string, logFile *os.File) *exec.Cmd {
	jq.logger.Info(resultsFile)
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
	channelWriter := consoleEcho.NewChannelWriter(jq.consoleChan)
	result.Stdout = io.MultiWriter(os.Stdout, channelWriter, logFile)
	result.Stderr = io.MultiWriter(os.Stderr, channelWriter, logFile)
	return result
}
