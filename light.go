package domain

// Light represents the state of a light entity.
// Brightness range: 0–254 (matches Zigbee/Z2M native scale; no conversion needed).
// Temperature is in mireds (153–500 typical range).
type Light struct {
	Power       bool      `json:"power"`
	Brightness  int       `json:"brightness"`
	ColorMode   string    `json:"colorMode,omitempty"`
	RGB         []int     `json:"rgb,omitempty"`
	RGBW        []int     `json:"rgbw,omitempty"`
	RGBWW       []int     `json:"rgbww,omitempty"`
	HS          []float64 `json:"hs,omitempty"`
	XY          []float64 `json:"xy,omitempty"`
	Temperature int       `json:"temperature,omitempty"`
	White       int       `json:"white,omitempty"`
	Effect      string    `json:"effect,omitempty"`
}

type LightPanel struct {
	Power       bool    `json:"power"`
	Brightness  int     `json:"brightness"`
	RGB         []int   `json:"rgb,omitempty"`
	Effect      string  `json:"effect,omitempty"`
	EffectSpeed int     `json:"effectSpeed,omitempty"`
	ColorMode   string  `json:"colorMode,omitempty"`
	Panels      []Panel `json:"panels,omitempty"`
}

type Panel struct {
	ID         int   `json:"id"`
	RGB        []int `json:"rgb,omitempty"`
	Brightness int   `json:"brightness,omitempty"`
}

type LightStrip struct {
	Power       bool      `json:"power"`
	Brightness  int       `json:"brightness"`
	RGB         []int     `json:"rgb,omitempty"`
	Temperature int       `json:"temperature,omitempty"`
	Effect      string    `json:"effect,omitempty"`
	EffectSpeed int       `json:"effectSpeed,omitempty"`
	ColorMode   string    `json:"colorMode,omitempty"`
	Segments    []Segment `json:"segments,omitempty"`
	Targets     []string  `json:"targets,omitempty"` // ordered entity keys — Targets[i] = segment i+1
}

type Segment struct {
	ID         int   `json:"id"`
	RGB        []int `json:"rgb,omitempty"`
	Brightness int   `json:"brightness,omitempty"`
}
