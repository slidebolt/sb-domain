package domain

import "fmt"

// ValidHVACModes lists the accepted HVAC mode strings.
var ValidHVACModes = []string{"off", "heat", "cool", "auto", "dry", "fan_only", "heat_cool"}

// ClimateSetMode sets the HVAC mode on a climate entity.
type ClimateSetMode struct {
	HVACMode string `json:"hvacMode"`
}

func (ClimateSetMode) ActionName() string { return "climate_set_mode" }

func (c ClimateSetMode) Validate() error {
	if c.HVACMode == "" {
		return fmt.Errorf("hvacMode must not be empty")
	}
	for _, m := range ValidHVACModes {
		if c.HVACMode == m {
			return nil
		}
	}
	return fmt.Errorf("hvacMode %q not one of %v", c.HVACMode, ValidHVACModes)
}

// ClimateSetTemperature sets the target temperature.
type ClimateSetTemperature struct {
	Temperature float64 `json:"temperature"`
}

func (ClimateSetTemperature) ActionName() string { return "climate_set_temperature" }

func (c ClimateSetTemperature) Validate() error {
	if c.Temperature < -50 || c.Temperature > 100 {
		return fmt.Errorf("temperature %.1f°C out of range [-50,100]", c.Temperature)
	}
	return nil
}
