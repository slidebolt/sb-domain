package domain

import "fmt"

// CoverOpen commands a cover to fully open.
type CoverOpen struct{}

func (CoverOpen) ActionName() string { return "cover_open" }

// CoverClose commands a cover to fully close.
type CoverClose struct{}

func (CoverClose) ActionName() string { return "cover_close" }

// CoverSetPosition sets a cover to a specific position (0=closed, 100=open).
type CoverSetPosition struct {
	Position int `json:"position"`
}

func (CoverSetPosition) ActionName() string { return "cover_set_position" }

func (c CoverSetPosition) Validate() error {
	if c.Position < 0 || c.Position > 100 {
		return fmt.Errorf("position %d out of range [0,100]", c.Position)
	}
	return nil
}
