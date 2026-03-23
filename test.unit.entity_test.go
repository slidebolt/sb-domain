package domain

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEntityUnmarshalPreservesCommands(t *testing.T) {
	raw := []byte(`{
		"id":"demo_light",
		"plugin":"plugin-clean",
		"deviceID":"demo_device",
		"type":"light",
		"name":"Demo Light",
		"commands":["light_turn_on","light_turn_off","light_set_color_temp"],
		"state":{"power":true,"brightness":128}
	}`)

	var got Entity
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal entity: %v", err)
	}

	want := []string{"light_turn_on", "light_turn_off", "light_set_color_temp"}
	if !reflect.DeepEqual(got.Commands, want) {
		t.Fatalf("commands = %v, want %v", got.Commands, want)
	}
}
