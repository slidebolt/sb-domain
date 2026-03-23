package domain

// SwitchTurnOn commands a switch to turn on.
type SwitchTurnOn struct{}

func (SwitchTurnOn) ActionName() string { return "switch_turn_on" }

// SwitchTurnOff commands a switch to turn off.
type SwitchTurnOff struct{}

func (SwitchTurnOff) ActionName() string { return "switch_turn_off" }

// SwitchToggle commands a switch to toggle its current state.
type SwitchToggle struct{}

func (SwitchToggle) ActionName() string { return "switch_toggle" }
