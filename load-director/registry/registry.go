package registry

import (
	"encoding/json"
	"load-director/logger"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var ()

const ()

type location string

type registryEntry struct {
	location   location
	lastUpdate time.Time
}

type Registry struct {
	log    *logrus.Entry
	slaves map[location]registryEntry
}

func New() *Registry {
	return &Registry{
		log:    logger.NewLogger().WithField("component", "registry"),
		slaves: make(map[location]registryEntry),
	}
}

func (r *Registry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(r.getCurrentSlaves())
		return
	}
	if req.Method == http.MethodPut {
		r.Put(location(req.RemoteAddr))
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func (r *Registry) getCurrentSlaves() []location {
	result := make([]location, len(r.slaves))
	i := 0
	for k, _ := range r.slaves {
		result[i] = k
		i++
	}
	return result
}

func (r *Registry) Put(slave location) {
	slaveIP := location(strings.Split(string(slave), ":")[0])
	_, ok := r.slaves[slaveIP]
	if ok {
		r.log.Infof("Slave reregistered under %s", slaveIP)
	} else {
		r.log.Infof("New slave registered under %s", slaveIP)
	}
	r.slaves[slaveIP] = registryEntry{
		location:   slaveIP,
		lastUpdate: time.Now(),
	}
}

func (r *Registry) StartCleanUpRoutine() {
	for _ = range time.Tick(1 * time.Minute) {
		r.cleanUp()
	}
}

func (r *Registry) cleanUp() {
	now := time.Now()
	var toDelete []location
	for key, entry := range r.slaves {
		if now.After(entry.lastUpdate.Add(5 * time.Minute)) {
			r.log.Infof("Removing slave from registry located under %s due to failing heartbeat", key)
			toDelete = append(toDelete, key)
		}
	}
	for _, v := range toDelete {
		delete(r.slaves, v)
	}
}
