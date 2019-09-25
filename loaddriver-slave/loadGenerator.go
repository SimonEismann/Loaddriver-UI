package main

import (
	"loaddriver-slave/shared"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func handleStop(w http.ResponseWriter, r *http.Request) {
	if sem.TryAcquire(1) {
		sem.Release(1)
		w.WriteHeader(http.StatusConflict)
		logger.WithField("func", "handleStop").Info("Load Generator is currently not running!")
		return
	}
	stopChan <- struct{}{}
	err := <-success
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.WithField("func", "handleStop").Error("Error occured while killing process!")
		return
	}
	w.WriteHeader(http.StatusOK)
	logger.WithField("func", "handleStop").Info("Load Generator successfully stopped!")
}

func startLoadGenerator() {
	defer sem.Release(1)
	execDir := shared.GetExecDir()
	cmd := exec.Command(
		"java",
		"-jar", filepath.Join(execDir, "httploadgenerator.jar"),
		"loadgenerator")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	shared.Must(err)
	<-stopChan
	err = cmd.Process.Kill()
	success <- err
	if err != nil {
		logger.WithError(err).Println("Error occured while killing process!")
		return
	}
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	if !sem.TryAcquire(1) {
		w.WriteHeader(http.StatusConflict)
		logger.WithField("func", "handleStart").Info("Load generator is already running!")
		return
	}
	go startLoadGenerator()
	w.WriteHeader(http.StatusOK)
	logger.WithField("func", "handleStart").Info("Load generator started successfully!")
}
