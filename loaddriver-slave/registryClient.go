package main

import (
	"fmt"
	"net/http"
	"time"
)

func startRegistryHeartbeat() {
	log := logger.WithField("comp", "registry-heartbeat")
	for c := time.Tick(10 * time.Second); ; <-c {
		client := http.DefaultClient
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s:%d/registry", registryHost, registryPort), nil)
		if err != nil {
			log.Panic("Malformed request, check parameters!", err)
		}
		res, err := client.Do(req)
		if err != nil {
			log.Warnf("Registry under http://%s:%d is not responding, trying again in 10 seconds", registryHost, registryPort)
			continue
		}
		if res.StatusCode != http.StatusOK {
			log.Warnf("Registry under http://%s:%d rejected registration, trying again in 10 seconds", registryHost, registryPort)
			continue
		}
	}
}
