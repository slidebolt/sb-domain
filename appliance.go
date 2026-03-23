package domain

type Siren struct {
IsOn           bool     `json:"isOn"`
AvailableTones []string `json:"availableTones,omitempty"`
}

type Humidifier struct {
IsOn            bool     `json:"isOn"`
TargetHumidity  int      `json:"targetHumidity"`
CurrentHumidity int      `json:"currentHumidity"`
MinHumidity     int      `json:"minHumidity"`
MaxHumidity     int      `json:"maxHumidity"`
Mode            string   `json:"mode,omitempty"`
AvailableModes  []string `json:"availableModes,omitempty"`
}

type Valve struct {
Position        int    `json:"position"`
ReportsPosition bool   `json:"reportsPosition"`
DeviceClass     string `json:"deviceClass,omitempty"`
}
