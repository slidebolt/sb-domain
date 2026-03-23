package domain_test

import (
	"testing"

	"github.com/slidebolt/sb-domain"
)

func ptr[T any](v T) *T { return &v }

func TestLightSetBrightness_Validate(t *testing.T) {
	cases := []struct {
		name    string
		cmd     domain.LightSetBrightness
		wantErr bool
	}{
		{"valid 0", domain.LightSetBrightness{Brightness: 0}, false},
		{"valid 254", domain.LightSetBrightness{Brightness: 254}, false},
		{"valid 127", domain.LightSetBrightness{Brightness: 127}, false},
		{"255 rejected", domain.LightSetBrightness{Brightness: 255}, true},
		{"-1 rejected", domain.LightSetBrightness{Brightness: -1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestLightSetColorTemp_Validate(t *testing.T) {
	cases := []struct {
		name    string
		cmd     domain.LightSetColorTemp
		wantErr bool
	}{
		{"valid 153 (min)", domain.LightSetColorTemp{Mireds: 153}, false},
		{"valid 500 (max)", domain.LightSetColorTemp{Mireds: 500}, false},
		{"valid 370", domain.LightSetColorTemp{Mireds: 370}, false},
		{"152 rejected (below min)", domain.LightSetColorTemp{Mireds: 152}, true},
		{"501 rejected", domain.LightSetColorTemp{Mireds: 501}, true},
		{"with brightness 128", domain.LightSetColorTemp{Mireds: 370, Brightness: 128}, false},
		{"with brightness 254", domain.LightSetColorTemp{Mireds: 370, Brightness: 254}, false},
		{"brightness 255 rejected", domain.LightSetColorTemp{Mireds: 370, Brightness: 255}, true},
		{"brightness negative rejected", domain.LightSetColorTemp{Mireds: 370, Brightness: -1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestLightSetRGB_Validate(t *testing.T) {
	cases := []struct {
		name    string
		cmd     domain.LightSetRGB
		wantErr bool
	}{
		{"valid all 0", domain.LightSetRGB{R: 0, G: 0, B: 0}, false},
		{"valid all 255", domain.LightSetRGB{R: 255, G: 255, B: 255}, false},
		{"valid mixed", domain.LightSetRGB{R: 255, G: 128, B: 0}, false},
		{"R out of range", domain.LightSetRGB{R: 256, G: 0, B: 0}, true},
		{"G negative", domain.LightSetRGB{R: 0, G: -1, B: 0}, true},
		{"with brightness 100", domain.LightSetRGB{R: 255, G: 128, B: 0, Brightness: 100}, false},
		{"brightness 255 rejected", domain.LightSetRGB{R: 255, G: 128, B: 0, Brightness: 255}, true},
		{"brightness negative rejected", domain.LightSetRGB{R: 255, G: 128, B: 0, Brightness: -1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestLightSetHS_Validate(t *testing.T) {
	cases := []struct {
		name    string
		cmd     domain.LightSetHS
		wantErr bool
	}{
		{"valid 0/0", domain.LightSetHS{Hue: 0, Saturation: 0}, false},
		{"valid 360/100", domain.LightSetHS{Hue: 360, Saturation: 100}, false},
		{"hue 361 rejected", domain.LightSetHS{Hue: 361, Saturation: 50}, true},
		{"saturation 101 rejected", domain.LightSetHS{Hue: 180, Saturation: 101}, true},
		{"hue negative rejected", domain.LightSetHS{Hue: -1, Saturation: 50}, true},
		{"with brightness 200", domain.LightSetHS{Hue: 180, Saturation: 50, Brightness: 200}, false},
		{"brightness 255 rejected", domain.LightSetHS{Hue: 180, Saturation: 50, Brightness: 255}, true},
		{"brightness negative rejected", domain.LightSetHS{Hue: 180, Saturation: 50, Brightness: -1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestLightSetXY_Validate(t *testing.T) {
	cases := []struct {
		name    string
		cmd     domain.LightSetXY
		wantErr bool
	}{
		{"valid 0/0", domain.LightSetXY{X: 0, Y: 0}, false},
		{"valid 1/1", domain.LightSetXY{X: 1, Y: 1}, false},
		{"valid 0.3/0.3", domain.LightSetXY{X: 0.3, Y: 0.3}, false},
		{"X > 1 rejected", domain.LightSetXY{X: 1.1, Y: 0.5}, true},
		{"Y negative rejected", domain.LightSetXY{X: 0.5, Y: -0.1}, true},
		{"with brightness 150", domain.LightSetXY{X: 0.3, Y: 0.3, Brightness: 150}, false},
		{"brightness 255 rejected", domain.LightSetXY{X: 0.3, Y: 0.3, Brightness: 255}, true},
		{"brightness negative rejected", domain.LightSetXY{X: 0.3, Y: 0.3, Brightness: -1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestLightSetWhite_Validate(t *testing.T) {
	cases := []struct {
		name    string
		cmd     domain.LightSetWhite
		wantErr bool
	}{
		{"valid 0", domain.LightSetWhite{White: 0}, false},
		{"valid 254", domain.LightSetWhite{White: 254}, false},
		{"255 rejected", domain.LightSetWhite{White: 255}, true},
		{"-1 rejected", domain.LightSetWhite{White: -1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestLightSetEffect_Validate(t *testing.T) {
	if err := (domain.LightSetEffect{Effect: "rainbow"}).Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := (domain.LightSetEffect{Effect: ""}).Validate(); err == nil {
		t.Error("expected error for empty effect")
	}
}

func TestFanSetSpeed_Validate(t *testing.T) {
	cases := []struct {
		name    string
		cmd     domain.FanSetSpeed
		wantErr bool
	}{
		{"valid 0", domain.FanSetSpeed{Percentage: 0}, false},
		{"valid 100", domain.FanSetSpeed{Percentage: 100}, false},
		{"valid 50", domain.FanSetSpeed{Percentage: 50}, false},
		{"101 rejected", domain.FanSetSpeed{Percentage: 101}, true},
		{"-1 rejected", domain.FanSetSpeed{Percentage: -1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestCoverSetPosition_Validate(t *testing.T) {
	cases := []struct {
		name    string
		cmd     domain.CoverSetPosition
		wantErr bool
	}{
		{"valid 0", domain.CoverSetPosition{Position: 0}, false},
		{"valid 100", domain.CoverSetPosition{Position: 100}, false},
		{"valid 50", domain.CoverSetPosition{Position: 50}, false},
		{"101 rejected", domain.CoverSetPosition{Position: 101}, true},
		{"-1 rejected", domain.CoverSetPosition{Position: -1}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestClimateSetMode_Validate(t *testing.T) {
	valid := []string{"off", "heat", "cool", "auto", "dry", "fan_only", "heat_cool"}
	for _, mode := range valid {
		t.Run("valid "+mode, func(t *testing.T) {
			if err := (domain.ClimateSetMode{HVACMode: mode}).Validate(); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
	invalid := []string{"", "warm", "HEAT", "Heat", "turbo"}
	for _, mode := range invalid {
		t.Run("invalid "+mode, func(t *testing.T) {
			if err := (domain.ClimateSetMode{HVACMode: mode}).Validate(); err == nil {
				t.Errorf("expected error for mode %q", mode)
			}
		})
	}
}

func TestClimateSetTemperature_Validate(t *testing.T) {
	cases := []struct {
		name    string
		temp    float64
		wantErr bool
	}{
		{"20°C", 20, false},
		{"0°C", 0, false},
		{"-50°C boundary", -50, false},
		{"100°C boundary", 100, false},
		{"-51°C rejected", -51, true},
		{"101°C rejected", 101, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := (domain.ClimateSetTemperature{Temperature: tc.temp}).Validate()
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestSelectOption_Validate(t *testing.T) {
	if err := (domain.SelectOption{Option: "mode1"}).Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := (domain.SelectOption{Option: ""}).Validate(); err == nil {
		t.Error("expected error for empty option")
	}
}

func TestTextSetValue_Validate(t *testing.T) {
	if err := (domain.TextSetValue{Value: "hello"}).Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if err := (domain.TextSetValue{Value: ""}).Validate(); err == nil {
		t.Error("expected error for empty value")
	}
}
