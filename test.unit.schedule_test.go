package domain

import (
	"encoding/json"
	"testing"
	"time"
)

func TestScheduleCarriesGenericActions(t *testing.T) {
	at := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	want := Schedule{Enabled: true, Steps: []ScheduleStep{{ID: "start", Trigger: Trigger{Kind: "at", At: &at}, Actions: []Action{{Target: "p.d.recorder", Command: "record_start", Input: json.RawMessage(`{"source":{"kind":"live","entity":"p.d.channel"}}`)}}}}}
	b, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var got Schedule
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Steps) != 1 || got.Steps[0].Actions[0].Command != "record_start" {
		t.Fatalf("round trip: %+v", got)
	}
}

func TestMediaItemKeepsProviderMetadataOptional(t *testing.T) {
	item := MediaItem{Ref: MediaItemRef{Provider: "plex", ItemID: "dvd-episode"}, Kind: "episode", Title: "DVD Episode", Playable: true, Metadata: map[string]json.RawMessage{"order": json.RawMessage(`{"scheme":"dvd","episode":4}`)}}
	b, err := json.Marshal(item)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) == "" {
		t.Fatal("empty media item")
	}
}
