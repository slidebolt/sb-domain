package domain

import "fmt"

// FanTurnOn commands a fan to turn on at its current speed.
type FanTurnOn struct{}

func (FanTurnOn) ActionName() string { return "fan_turn_on" }

// FanTurnOff commands a fan to turn off.
type FanTurnOff struct{}

func (FanTurnOff) ActionName() string { return "fan_turn_off" }

// FanSetSpeed sets a fan's speed percentage (0–100).
type FanSetSpeed struct {
	Percentage int `json:"percentage"`
}

func (FanSetSpeed) ActionName() string { return "fan_set_speed" }

func (c FanSetSpeed) Validate() error {
	if c.Percentage < 0 || c.Percentage > 100 {
		return fmt.Errorf("percentage %d out of range [0,100]", c.Percentage)
	}
	return nil
}
