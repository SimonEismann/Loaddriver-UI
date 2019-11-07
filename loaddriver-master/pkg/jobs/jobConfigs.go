package jobs

import "time"

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
