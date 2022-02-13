package models

import "time"

// Health struct for health-check response
type Health struct {
	App string `json:"app"`
	Code int64 `json:"code"`
	Time time.Time `json:"time"`
}
