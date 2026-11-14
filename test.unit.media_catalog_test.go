package domain

import (
	"encoding/json"
	"testing"
)

func TestMediaCatalogEntitiesDecode(t *testing.T) {
	for _, raw := range []string{
		`{"id":"library-1","plugin":"plugin-plex","deviceID":"server-1","type":"media_library","name":"TV","state":{"serverID":"server-1","kind":"show","canBrowse":true,"canSearch":true}}`,
		`{"id":"session-1","plugin":"plugin-plex","deviceID":"server-1","type":"media_session","name":"Now Playing","state":{"serverID":"server-1","state":"playing"}}`,
		`{"id":"recorder","plugin":"plugin-hdhomerun","deviceID":"tuner-1","type":"recorder","name":"Recorder","state":{"status":"idle"}}`,
	} {
		var e Entity
		if err := json.Unmarshal([]byte(raw), &e); err != nil {
			t.Fatal(err)
		}
	}
}

func TestMediaCatalogCommandsRegistered(t *testing.T) {
	for _, action := range []string{"media_provider_connect", "media_provider_disconnect", "media_play", "media_seek", "record_start", "record_stop", "recording_rule_create", "recording_rule_delete", "schedule_enable"} {
		if _, ok := LookupCommand(action); !ok {
			t.Fatalf("command %s not registered", action)
		}
	}
}
