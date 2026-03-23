package domain

type Sensor struct {
	Value       any    `json:"value"`
	Unit        string `json:"unit,omitempty"`
	DeviceClass string `json:"deviceClass,omitempty"`
}

type BinarySensor struct {
	On          bool   `json:"on"`
	DeviceClass string `json:"deviceClass,omitempty"`
}

type Number struct {
	Value       float64 `json:"value"`
	Min         float64 `json:"min"`
	Max         float64 `json:"max"`
	Step        float64 `json:"step"`
	Unit        string  `json:"unit,omitempty"`
	DeviceClass string  `json:"deviceClass,omitempty"`
}

type Select struct {
	Option  string   `json:"option"`
	Options []string `json:"options"`
}

type Text struct {
	Value   string `json:"value"`
	Min     int    `json:"min,omitempty"`
	Max     int    `json:"max,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	Mode    string `json:"mode,omitempty"`
}
