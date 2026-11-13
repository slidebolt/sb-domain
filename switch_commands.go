package domain

import "encoding/json"

// SwitchTurnOn commands a switch to turn on.
type SwitchTurnOn struct {
	// Input carries optional target-specific activation data. Ordinary switches
	// ignore it; recording switches use it to select a channel and output profile.
	Input json.RawMessage `json:"input,omitempty"`
}

func (SwitchTurnOn) ActionName() string { return "switch_turn_on" }

// SwitchTurnOff commands a switch to turn off.
type SwitchTurnOff struct{}

func (SwitchTurnOff) ActionName() string { return "switch_turn_off" }

// SwitchToggle commands a switch to toggle its current state.
type SwitchToggle struct{}

func (SwitchToggle) ActionName() string { return "switch_toggle" }
