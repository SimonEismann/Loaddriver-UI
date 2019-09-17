package main

import (
	"io"
	"load-generator/shared"
	"net/http"
	"os/exec"
	"path/filepath"
)

func handleStop(w http.ResponseWriter, r *http.Request) {
	if sem.TryAcquire(1) {
		sem.Release(1)
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "Load Generator is currently not running!")
		return
	}
	stopChan <- struct{}{}
	err := <-success
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error occured while killing process!")
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Load Generator successfully stopped!")
}

func startLoadGenerator() {
	defer sem.Release(1)
	execDir := shared.GetExecDir()
	cmd := exec.Command(
		"java",
		"-jar", filepath.Join(execDir, "httploadgenerator.jar"),
		"loadgenerator")
	err := cmd.Start()
	shared.Must(err)
	for {
		select {
		case <-stopChan:
			err := cmd.Process.Kill()
			success <- err
			if err == nil {
				logger.Println("Load Generator successfully stopped!")
				return
			}
			logger.WithError(err).Println("Error occured while killing process!")
			return
		}
	}
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	if !sem.TryAcquire(1) {
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "Load generator is already running!")
		return
	}
	go startLoadGenerator()
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Load generator started successfully!")
}
