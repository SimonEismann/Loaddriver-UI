package api

import (
	"encoding/json"
	"net/http"
)

func PrintDescription(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	resp, err := json.MarshalIndent(map[string]string{
		"GET /scripts":             "Lists the available script file names",
		"GET /scripts/{fileName}":  "Returns the script with the specified file name",
		"POST /jobs":               "Queue an experiment with all registered slaves",
		"POST /jobs/default":       "Queue the default experiment with all registered slaves",
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
