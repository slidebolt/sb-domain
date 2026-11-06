package domain

import (
	"encoding/json"
	"testing"
)

func TestMediaStream_HydratesFromEntityJSON(t *testing.T) {
	data := []byte(`{"id":"audio","plugin":"p","deviceID":"d","type":"media_stream","state":{"status":"live","kind":"audio","url":"http://x/y","format":"mpegts","viewers":2}}`)
	var e Entity
	if err := json.Unmarshal(data, &e); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	s, ok := e.State.(MediaStream)
	if !ok {
		t.Fatalf("state type = %T, want MediaStream", e.State)
	}
	if s.Status != "live" || s.Kind != "audio" || s.URL != "http://x/y" || s.Format != "mpegts" || s.Viewers != 2 {
		t.Fatalf("state = %+v", s)
	}
}

func TestMediaStream_CommandsRegistered(t *testing.T) {
	for _, action := range []string{"media_stream_start", "media_stream_stop"} {
		if _, ok := LookupCommand(action); !ok {
			t.Errorf("%s not registered", action)
		}
	}
	if (MediaStreamStart{}).ActionName() != "media_stream_start" || (MediaStreamStop{}).ActionName() != "media_stream_stop" {
		t.Error("action names do not match registry")
	}
}
