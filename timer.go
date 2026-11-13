package domain

import (
	"encoding/json"
	"time"
)

// TimerTrigger is one edge of a timer. Kind is "at" or "cron".
type TimerTrigger struct {
	Kind     string     `json:"kind"`
	At       *time.Time `json:"at,omitempty"`
	Cron     string     `json:"cron,omitempty"`
	TimeZone string     `json:"timeZone,omitempty"`
}

// Timer sets a target switch on and off. Multiple timers deliberately have
// level-triggered semantics: they do not own or reserve the target switch.
type Timer struct {
	Enabled bool            `json:"enabled"`
	Target  string          `json:"target"`
	Input   json.RawMessage `json:"input,omitempty"`
	On      TimerTrigger    `json:"on"`
	Off     TimerTrigger    `json:"off"`

	NextOnAt  *time.Time `json:"nextOnAt,omitempty"`
	NextOffAt *time.Time `json:"nextOffAt,omitempty"`
	LastEdge  string     `json:"lastEdge,omitempty"`
	LastRunAt *time.Time `json:"lastRunAt,omitempty"`
	LastError string     `json:"lastError,omitempty"`
}
