package model

import "time"

type Build struct {
	JobName    string    `json:"job_name"`
	Number     int       `json:"number"`
	Status     string    `json:"status"`
	StartedAt  time.Time `json:"started_at"`
	DurationMs int64     `json:"duration_ms"`
	URL        string    `json:"url"`
}
