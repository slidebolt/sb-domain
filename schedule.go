package domain

import (
	"encoding/json"
	"time"
)

type Action struct {
	Target  string          `json:"target"`
	Command string          `json:"command"`
	Input   json.RawMessage `json:"input,omitempty"`
}

type Trigger struct {
	Kind     string          `json:"kind"`
	At       *time.Time      `json:"at,omitempty"`
	Cron     string          `json:"cron,omitempty"`
	TimeZone string          `json:"timeZone,omitempty"`
	Source   string          `json:"source,omitempty"`
	Event    string          `json:"event,omitempty"`
	Match    json.RawMessage `json:"match,omitempty"`
}

type ScheduleStep struct {
	ID        string     `json:"id"`
	Trigger   Trigger    `json:"trigger"`
	Actions   []Action   `json:"actions"`
	NextRunAt *time.Time `json:"nextRunAt,omitempty"`
	LastRunAt *time.Time `json:"lastRunAt,omitempty"`
	LastError string     `json:"lastError,omitempty"`
}

type Schedule struct {
	Enabled bool           `json:"enabled"`
	Steps   []ScheduleStep `json:"steps"`
}

type ScheduleEnable struct{}

func (ScheduleEnable) ActionName() string { return "schedule_enable" }

type ScheduleDisable struct{}

func (ScheduleDisable) ActionName() string { return "schedule_disable" }
