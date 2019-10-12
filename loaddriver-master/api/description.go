package api

import (
	"encoding/json"
	"net/http"
)

func PrintDescription(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	resp, err := json.MarshalIndent(map[string]string{
		"GET /scripts":                   "Lists the available script file locations",
		"POST /scripts":                  "Upload a script file to the server",
		"GET /scripts/{fileName}":        "Returns the script file at the specified location",
		"GET /intensities":               "Lists the available intensity file locations",
		"POST /intensities":              "Upload an intensity file to the server",
		"GET /intensities/{fileName}":    "Returns the intensity file at the specified location",
		"POST /jobs":                     "Queue an experiment with all registered slaves",
		"POST /jobs/default":             "Queue the default experiment with all registered slaves",
		"GET /jobs":                      "Lists all jobs",
		"GET /jobs/done":                 "Lists all done jobs",
		"GET /jobs/{jobId}":              "Get the current status of the specified job",
		"GET /jobs/{jobId}/result":       "Returns the result file of the job",
		"GET /jobs/{jobId}/log":          "Retunrs the log file of the job",
		"WEBSOCKET /jobs/current/output": "Opens a socket showing the output of the currently running job",
		"DELETE /jobs/current":           "Force stop the currently running load generation",
		"GET /registry":                  "List all currently registered slaves",
	}, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
