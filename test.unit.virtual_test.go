package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/slidebolt/sb-domain"
)

func TestEntity_TargetRoundTrip(t *testing.T) {
	target := json.RawMessage(`{"pattern":"*.*.*","where":[{"field":"type","op":"eq","value":"light"}]}`)

	e := domain.Entity{
		ID:       "group1",
		Plugin:   "virtual",
		DeviceID: "living-room",
		Type:     "light",
		Name:     "All Lights",
		Target:   target,
		State:    domain.Light{Power: true, Brightness: 200},
	}

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got domain.Entity
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if got.ID != e.ID {
		t.Errorf("ID = %q, want %q", got.ID, e.ID)
	}
	if got.Target == nil {
		t.Fatal("Target is nil after round-trip")
	}

	// Verify Target is preserved as-is.
	var gotQuery map[string]any
	if err := json.Unmarshal(got.Target, &gotQuery); err != nil {
		t.Fatalf("unmarshal target: %v", err)
	}
	if gotQuery["pattern"] != "*.*.*" {
		t.Errorf("target pattern = %v, want *.*.*", gotQuery["pattern"])
	}
}

func TestEntity_NilTargetOmitted(t *testing.T) {
	e := domain.Entity{
		ID:       "bulb1",
		Plugin:   "wiz",
		DeviceID: "device1",
		Type:     "light",
		Name:     "Bulb",
		State:    domain.Light{Power: false},
	}

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshal raw: %v", err)
	}

	if _, ok := raw["target"]; ok {
		t.Error("target should be omitted when nil")
	}
}

func TestEntity_TargetFromJSON(t *testing.T) {
	input := `{
		"id": "group1",
		"plugin": "virtual",
		"deviceID": "room",
		"type": "light",
		"name": "Group",
		"target": {"pattern": "wiz.*.*", "where": [{"field": "type", "op": "eq", "value": "light"}]},
		"state": {"power": true, "brightness": 100}
	}`

	var e domain.Entity
	if err := json.Unmarshal([]byte(input), &e); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if e.Target == nil {
		t.Fatal("Target should not be nil")
	}
	if e.ID != "group1" {
		t.Errorf("ID = %q, want group1", e.ID)
	}

	// State should still hydrate correctly.
	light, ok := e.State.(domain.Light)
	if !ok {
		t.Fatalf("State type = %T, want domain.Light", e.State)
	}
	if !light.Power || light.Brightness != 100 {
		t.Errorf("State = %+v, want power=true brightness=100", light)
	}
}

func TestLightStrip_TargetsRoundTrip(t *testing.T) {
	strip := domain.LightStrip{
		Power:      true,
		Brightness: 200,
		RGB:        []int{255, 0, 0},
		ColorMode:  "rgb",
		Targets:    []string{"wiz.room.bulb1", "wiz.room.bulb2", "wiz.room.bulb3"},
		Segments: []domain.Segment{
			{ID: 1, RGB: []int{255, 0, 0}, Brightness: 200},
			{ID: 2, RGB: []int{0, 255, 0}, Brightness: 200},
			{ID: 3, RGB: []int{0, 0, 255}, Brightness: 200},
		},
	}

	data, err := json.Marshal(strip)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got domain.LightStrip
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if len(got.Targets) != 3 {
		t.Fatalf("Targets len = %d, want 3", len(got.Targets))
	}
	for i, want := range strip.Targets {
		if got.Targets[i] != want {
			t.Errorf("Targets[%d] = %q, want %q", i, got.Targets[i], want)
		}
	}
	if len(got.Segments) != 3 {
		t.Fatalf("Segments len = %d, want 3", len(got.Segments))
	}
}

func TestLightStrip_TargetsOmittedWhenEmpty(t *testing.T) {
	strip := domain.LightStrip{Power: true, Brightness: 100}

	data, err := json.Marshal(strip)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshal raw: %v", err)
	}

	if _, ok := raw["targets"]; ok {
		t.Error("targets should be omitted when empty")
	}
}

func TestLightStrip_EntityRoundTrip(t *testing.T) {
	e := domain.Entity{
		ID:       "strip1",
		Plugin:   "virtual",
		DeviceID: "living-room",
		Type:     "light_strip",
		Name:     "Virtual Strip",
		State: domain.LightStrip{
			Power:      true,
			Brightness: 180,
			Targets:    []string{"wiz.room.bulb1", "wiz.room.bulb2"},
			Segments: []domain.Segment{
				{ID: 1, RGB: []int{255, 0, 0}, Brightness: 180},
				{ID: 2, RGB: []int{0, 255, 0}, Brightness: 180},
			},
		},
	}

	data, err := json.Marshal(e)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var got domain.Entity
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	strip, ok := got.State.(domain.LightStrip)
	if !ok {
		t.Fatalf("State type = %T, want domain.LightStrip", got.State)
	}
	if len(strip.Targets) != 2 {
		t.Errorf("Targets len = %d, want 2", len(strip.Targets))
	}
	if strip.Targets[0] != "wiz.room.bulb1" {
		t.Errorf("Targets[0] = %q, want wiz.room.bulb1", strip.Targets[0])
	}
}

func TestLightstripSetSegments_Validate(t *testing.T) {
	valid := domain.LightstripSetSegments{
		Power:     true,
		ColorMode: "rgb",
		Segments: []domain.Segment{
			{ID: 1, RGB: []int{255, 0, 0}, Brightness: 180},
		},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	empty := domain.LightstripSetSegments{Power: true}
	if err := empty.Validate(); err == nil {
		t.Error("expected error for empty segments")
	}
}

func TestLightstripSetSegments_ActionName(t *testing.T) {
	cmd := domain.LightstripSetSegments{}
	if cmd.ActionName() != "lightstrip_set_segments" {
		t.Errorf("ActionName() = %q, want lightstrip_set_segments", cmd.ActionName())
	}
}

func TestLightstripSetSegments_Registered(t *testing.T) {
	// Verify the command can be looked up from the registry.
	typ, ok := domain.LookupCommand("lightstrip_set_segments")
	if !ok {
		t.Fatal("lightstrip_set_segments not registered")
	}
	if typ.Name() != "LightstripSetSegments" {
		t.Errorf("registered type = %q, want LightstripSetSegments", typ.Name())
	}
}
