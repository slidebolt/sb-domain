package domain

type Switch struct {
Power bool `json:"power"`
}

type Fan struct {
Power      bool `json:"power"`
Percentage int  `json:"percentage"`
}

type Cover struct {
Position    int    `json:"position"`
DeviceClass string `json:"deviceClass,omitempty"`
}

type Lock struct {
Locked bool `json:"locked"`
}

type Button struct {
Presses     int    `json:"presses"`
DeviceClass string `json:"deviceClass,omitempty"`
}
