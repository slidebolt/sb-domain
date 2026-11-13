package domain

import (
	"encoding/json"
	"testing"
	"time"
)

func TestEntityResourceStatusRoundTrip(t *testing.T) {
	observed := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	in := Entity{ID: "channel", Plugin: "plugin-tv", DeviceID: "tuner", Type: "tv_channel", Name: "Channel", Status: &ResourceStatus{Health: "healthy", ObservedAt: observed, StaleAfterSeconds: 120}, State: map[string]any{}}
	b, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	var out Entity
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	if out.Status == nil || out.Status.Health != "healthy" || !out.Status.ObservedAt.Equal(observed) {
		t.Fatalf("status did not round trip: %#v", out.Status)
	}
}
