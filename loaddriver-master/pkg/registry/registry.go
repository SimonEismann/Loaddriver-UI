package registry

import (
	"encoding/json"
	"loaddriver-master/logger"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type IP string

type RegistryEntry struct {
	Location   IP
	LastUpdate time.Time
}

type Registry struct {
	log    *logrus.Entry
	Slaves map[IP]RegistryEntry
}

func New() *Registry {
	return &Registry{
		log:    logger.NewLogger().WithField("comp", "registry"),
		Slaves: make(map[IP]RegistryEntry),
	}
}

func (r *Registry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(r.getCurrentSlaves())
		return
	}
	if req.Method == http.MethodPut {
		r.Put(IP(req.RemoteAddr))
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func (r *Registry) getCurrentSlaves() []IP {
	result := make([]IP, len(r.Slaves))
	i := 0
	for k, _ := range r.Slaves {
		result[i] = k
		i++
	}
	return result
}

func (r *Registry) Put(slave IP) {
	slaveIP := IP(strings.Split(string(slave), ":")[0])
	_, ok := r.Slaves[slaveIP]
	if ok {
		r.log.Infof("Slave reregistered under %s", slaveIP)
	} else {
		r.log.Infof("New slave registered under %s", slaveIP)
	}
	r.Slaves[slaveIP] = RegistryEntry{
		Location:   slaveIP,
		LastUpdate: time.Now(),
	}
}

func (r *Registry) StartCleanUpRoutine() {
	for _ = range time.Tick(1 * time.Minute) {
		r.cleanUp()
	}
}

func (r *Registry) cleanUp() {
	now := time.Now()
	var toDelete []IP
	for key, entry := range r.Slaves {
		if now.After(entry.LastUpdate.Add(5 * time.Minute)) {
			r.log.Infof("Removing slave from registry located under %s due to failing heartbeat", key)
			toDelete = append(toDelete, key)
		}
	}
	for _, v := range toDelete {
		delete(r.Slaves, v)
	}
}
