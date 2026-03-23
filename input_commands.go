package domain

import "fmt"

// NumberSetValue sets a number entity's value.
type NumberSetValue struct {
	Value float64 `json:"value"`
}

func (NumberSetValue) ActionName() string { return "number_set_value" }

// SelectOption selects an option on a select entity.
type SelectOption struct {
	Option string `json:"option"`
}

func (SelectOption) ActionName() string { return "select_option" }

func (c SelectOption) Validate() error {
	if c.Option == "" {
		return fmt.Errorf("option must not be empty")
	}
	return nil
}

// TextSetValue sets a text entity's value.
type TextSetValue struct {
	Value string `json:"value"`
}

func (TextSetValue) ActionName() string { return "text_set_value" }

func (c TextSetValue) Validate() error {
	if c.Value == "" {
		return fmt.Errorf("value must not be empty")
	}
	return nil
}
