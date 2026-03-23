package domain

// ButtonPress commands a button press.
type ButtonPress struct{}

func (ButtonPress) ActionName() string { return "button_press" }
