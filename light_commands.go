package domain

import "fmt"

// ValidColorModes lists the accepted color mode strings for lights.
var ValidColorModes = []string{"color_temp", "rgb", "rgbw", "rgbww", "hs", "xy", "white", "brightness"}

// LightTurnOn powers a light on at its current state.
type LightTurnOn struct{}

func (LightTurnOn) ActionName() string { return "light_turn_on" }

// LightTurnOff powers a light off.
type LightTurnOff struct {
	Transition *int `json:"transition,omitempty"`
}

func (LightTurnOff) ActionName() string { return "light_turn_off" }

// LightSetBrightness sets brightness (0–254).
type LightSetBrightness struct {
	Brightness int  `json:"brightness"`
	Transition *int `json:"transition,omitempty"`
}

func (LightSetBrightness) ActionName() string { return "light_set_brightness" }

func (c LightSetBrightness) Validate() error {
	if c.Brightness < 0 || c.Brightness > 254 {
		return fmt.Errorf("brightness %d out of range [0,254]", c.Brightness)
	}
	return nil
}

// LightSetColorTemp sets color temperature in mireds (153–500).
// Brightness is optional (0 means no change); when provided it must be in [0,254].
type LightSetColorTemp struct {
	Mireds     int  `json:"mireds"`
	Brightness int  `json:"brightness,omitempty"`
	Transition *int `json:"transition,omitempty"`
}

func (LightSetColorTemp) ActionName() string { return "light_set_color_temp" }

func (c LightSetColorTemp) Validate() error {
	if c.Mireds < 153 || c.Mireds > 500 {
		return fmt.Errorf("mireds %d out of range [153,500]", c.Mireds)
	}
	if c.Brightness < 0 || c.Brightness > 254 {
		return fmt.Errorf("brightness %d out of range [0,254]", c.Brightness)
	}
	return nil
}

// LightSetRGB sets RGB color. Each component is 0–255.
// Brightness is optional (0 means no change); when provided it must be in [0,254].
type LightSetRGB struct {
	R          int  `json:"r"`
	G          int  `json:"g"`
	B          int  `json:"b"`
	Brightness int  `json:"brightness,omitempty"`
	Transition *int `json:"transition,omitempty"`
}

func (LightSetRGB) ActionName() string { return "light_set_rgb" }

func (c LightSetRGB) Validate() error {
	for name, v := range map[string]int{"r": c.R, "g": c.G, "b": c.B} {
		if v < 0 || v > 255 {
			return fmt.Errorf("%s value %d out of range [0,255]", name, v)
		}
	}
	if c.Brightness < 0 || c.Brightness > 254 {
		return fmt.Errorf("brightness %d out of range [0,254]", c.Brightness)
	}
	return nil
}

// LightSetRGBW sets RGB + white channel. Each component is 0–255.
// Brightness is optional (0 means no change); when provided it must be in [0,254].
type LightSetRGBW struct {
	R          int  `json:"r"`
	G          int  `json:"g"`
	B          int  `json:"b"`
	W          int  `json:"w"`
	Brightness int  `json:"brightness,omitempty"`
	Transition *int `json:"transition,omitempty"`
}

func (LightSetRGBW) ActionName() string { return "light_set_rgbw" }

func (c LightSetRGBW) Validate() error {
	for name, v := range map[string]int{"r": c.R, "g": c.G, "b": c.B, "w": c.W} {
		if v < 0 || v > 255 {
			return fmt.Errorf("%s value %d out of range [0,255]", name, v)
		}
	}
	if c.Brightness < 0 || c.Brightness > 254 {
		return fmt.Errorf("brightness %d out of range [0,254]", c.Brightness)
	}
	return nil
}

// LightSetRGBWW sets RGB + cool white + warm white. Each component is 0–255.
// Brightness is optional (0 means no change); when provided it must be in [0,254].
type LightSetRGBWW struct {
	R          int  `json:"r"`
	G          int  `json:"g"`
	B          int  `json:"b"`
	CW         int  `json:"cw"`
	WW         int  `json:"ww"`
	Brightness int  `json:"brightness,omitempty"`
	Transition *int `json:"transition,omitempty"`
}

func (LightSetRGBWW) ActionName() string { return "light_set_rgbww" }

func (c LightSetRGBWW) Validate() error {
	for name, v := range map[string]int{"r": c.R, "g": c.G, "b": c.B, "cw": c.CW, "ww": c.WW} {
		if v < 0 || v > 255 {
			return fmt.Errorf("%s value %d out of range [0,255]", name, v)
		}
	}
	if c.Brightness < 0 || c.Brightness > 254 {
		return fmt.Errorf("brightness %d out of range [0,254]", c.Brightness)
	}
	return nil
}

// LightSetHS sets hue (0–360) and saturation (0–100).
// Brightness is optional (0 means no change); when provided it must be in [0,254].
type LightSetHS struct {
	Hue        float64 `json:"hue"`
	Saturation float64 `json:"saturation"`
	Brightness int     `json:"brightness,omitempty"`
	Transition *int    `json:"transition,omitempty"`
}

func (LightSetHS) ActionName() string { return "light_set_hs" }

func (c LightSetHS) Validate() error {
	if c.Hue < 0 || c.Hue > 360 {
		return fmt.Errorf("hue %.2f out of range [0,360]", c.Hue)
	}
	if c.Saturation < 0 || c.Saturation > 100 {
		return fmt.Errorf("saturation %.2f out of range [0,100]", c.Saturation)
	}
	if c.Brightness < 0 || c.Brightness > 254 {
		return fmt.Errorf("brightness %d out of range [0,254]", c.Brightness)
	}
	return nil
}

// LightSetXY sets CIE xy chromaticity coordinates (each 0.0–1.0).
// Brightness is optional (0 means no change); when provided it must be in [0,254].
type LightSetXY struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	Brightness int     `json:"brightness,omitempty"`
	Transition *int    `json:"transition,omitempty"`
}

func (LightSetXY) ActionName() string { return "light_set_xy" }

func (c LightSetXY) Validate() error {
	if c.X < 0 || c.X > 1 {
		return fmt.Errorf("x %.4f out of range [0,1]", c.X)
	}
	if c.Y < 0 || c.Y > 1 {
		return fmt.Errorf("y %.4f out of range [0,1]", c.Y)
	}
	if c.Brightness < 0 || c.Brightness > 254 {
		return fmt.Errorf("brightness %d out of range [0,254]", c.Brightness)
	}
	return nil
}

// LightSetWhite sets the dedicated white channel (0–254).
type LightSetWhite struct {
	White      int  `json:"white"`
	Transition *int `json:"transition,omitempty"`
}

func (LightSetWhite) ActionName() string { return "light_set_white" }

func (c LightSetWhite) Validate() error {
	if c.White < 0 || c.White > 254 {
		return fmt.Errorf("white %d out of range [0,254]", c.White)
	}
	return nil
}

// LightSetEffect activates a named light effect.
type LightSetEffect struct {
	Effect string `json:"effect"`
}

func (LightSetEffect) ActionName() string { return "light_set_effect" }

func (c LightSetEffect) Validate() error {
	if c.Effect == "" {
		return fmt.Errorf("effect must not be empty")
	}
	return nil
}

// LightstripSetSegments sets per-segment state on a light strip.
type LightstripSetSegments struct {
	Power       bool      `json:"power"`
	ColorMode   string    `json:"colorMode,omitempty"`
	Effect      string    `json:"effect,omitempty"`
	EffectSpeed int       `json:"effectSpeed,omitempty"`
	Segments    []Segment `json:"segments"`
}

func (LightstripSetSegments) ActionName() string { return "lightstrip_set_segments" }

func (c LightstripSetSegments) Validate() error {
	if len(c.Segments) == 0 {
		return fmt.Errorf("segments must not be empty")
	}
	return nil
}

