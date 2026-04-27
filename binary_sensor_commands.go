package domain

// BinarySensorTurnOn sets a binary sensor to on.
type BinarySensorTurnOn struct{}

func (BinarySensorTurnOn) ActionName() string { return "binary_sensor_turn_on" }

// BinarySensorTurnOff sets a binary sensor to off.
type BinarySensorTurnOff struct{}

func (BinarySensorTurnOff) ActionName() string { return "binary_sensor_turn_off" }
