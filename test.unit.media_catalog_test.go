package domain

import (
	"encoding/json"
	"testing"
)

func TestMediaCatalogEntitiesDecode(t *testing.T) {
	for _, raw := range []string{
		`{"id":"library-1","plugin":"plugin-plex","deviceID":"server-1","type":"media_library","name":"TV","state":{"serverID":"server-1","kind":"show","canBrowse":true,"canSearch":true}}`,
		`{"id":"session-1","plugin":"plugin-plex","deviceID":"server-1","type":"media_session","name":"Now Playing","state":{"serverID":"server-1","state":"playing"}}`,
		`{"id":"airing-1","plugin":"plugin-plex","deviceID":"server-1","type":"tv_recording","name":"News","state":{"status":"scheduled","provider":"plex","channel":{"number":"4.2","name":"Start TV","available":true},"program":{"airingID":"a1","title":"News","startsAt":"2026-11-13T01:00:00Z","endsAt":"2026-11-13T02:00:00Z"},"startsAt":"2026-11-13T01:00:00Z","endsAt":"2026-11-13T02:00:00Z"}}`,
	} {
		var e Entity
		if err := json.Unmarshal([]byte(raw), &e); err != nil {
			t.Fatal(err)
		}
	}
}

func TestMediaCatalogCommandsRegistered(t *testing.T) {
	for _, action := range []string{"media_provider_connect", "media_provider_disconnect", "media_play_item", "media_seek", "tv_record_cancel"} {
		if _, ok := LookupCommand(action); !ok {
			t.Fatalf("command %s not registered", action)
		}
	}
}
