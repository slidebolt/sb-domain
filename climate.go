package domain

type Climate struct {
HVACMode           string   `json:"hvacMode"`
HVACModes          []string `json:"hvacModes,omitempty"`
Temperature        float64  `json:"temperature"`
CurrentTemperature float64  `json:"currentTemperature,omitempty"`
TemperatureUnit    string   `json:"temperatureUnit,omitempty"`
MinTemp            float64  `json:"minTemp,omitempty"`
MaxTemp            float64  `json:"maxTemp,omitempty"`
TargetTempStep     float64  `json:"targetTempStep,omitempty"`
}
