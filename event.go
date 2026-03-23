package domain

// Event is a read-only entity representing HA event entities (doorbell, button presses, etc.)
type Event struct {
EventTypes  []string `json:"eventTypes"`
DeviceClass string   `json:"deviceClass,omitempty"`
}
