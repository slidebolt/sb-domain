package domain

import "time"

// Event is a read-only entity representing HA event entities (doorbell, button presses, etc.)
type Event struct {
	EventTypes  []string       `json:"eventTypes"`
	DeviceClass string         `json:"deviceClass,omitempty"`
	LastEventID string         `json:"lastEventID,omitempty"`
	LastEventAt *time.Time     `json:"lastEventAt,omitempty"`
	Active      bool           `json:"active"`
	Attributes  map[string]any `json:"attributes,omitempty"`
}
