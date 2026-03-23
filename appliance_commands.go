package domain

// Siren commands
type SirenTurnOn struct{}

func (SirenTurnOn) ActionName() string { return "siren_turn_on" }

type SirenTurnOff struct{}

func (SirenTurnOff) ActionName() string { return "siren_turn_off" }

type SirenSetTone struct{ Tone string }

func (SirenSetTone) ActionName() string { return "siren_set_tone" }

// Humidifier commands
type HumidifierTurnOn struct{}

func (HumidifierTurnOn) ActionName() string { return "humidifier_turn_on" }

type HumidifierTurnOff struct{}

func (HumidifierTurnOff) ActionName() string { return "humidifier_turn_off" }

type HumidifierSetHumidity struct{ Humidity int }

func (HumidifierSetHumidity) ActionName() string { return "humidifier_set_humidity" }

type HumidifierSetMode struct{ Mode string }

func (HumidifierSetMode) ActionName() string { return "humidifier_set_mode" }

// Valve commands
type ValveOpen struct{}

func (ValveOpen) ActionName() string { return "valve_open" }

type ValveClose struct{}

func (ValveClose) ActionName() string { return "valve_close" }

type ValveSetPosition struct{ Position int }

func (ValveSetPosition) ActionName() string { return "valve_set_position" }
