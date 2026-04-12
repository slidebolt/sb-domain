package domain

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestDevice_PrivateRoundtrip(t *testing.T) {
	priv := json.RawMessage(`{"host":"192.168.88.121","username":"admin","password":"secret"}`)
	d := Device{
		ID:      "cam1",
		Plugin:  "plugin-amcrest",
		Name:    "Cam",
		Private: priv,
	}
	raw, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !bytes.Contains(raw, []byte(`"private":`)) {
		t.Fatalf("expected private key in json, got %s", raw)
	}
	var got Device
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !bytes.Equal(got.Private, priv) {
		t.Fatalf("private mismatch: got %s, want %s", got.Private, priv)
	}
}

func TestDevice_PrivateOmittedWhenNil(t *testing.T) {
	d := Device{ID: "cam1", Plugin: "plugin-amcrest", Name: "Cam"}
	raw, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if strings.Contains(string(raw), `"private"`) {
		t.Fatalf("expected private key omitted, got %s", raw)
	}
}

func TestEntity_PrivateRoundtrip(t *testing.T) {
	priv := json.RawMessage(`{"api_key":"xyz"}`)
	e := Entity{
		ID:       "default",
		Plugin:   "plugin-amcrest",
		DeviceID: "default",
		Type:     "camera",
		Name:     "Cam",
		State:    Camera{},
		Private:  priv,
	}
	raw, err := json.Marshal(e)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !bytes.Contains(raw, []byte(`"private":`)) {
		t.Fatalf("expected private key in json, got %s", raw)
	}
	var got Entity
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !bytes.Equal(got.Private, priv) {
		t.Fatalf("private mismatch: got %s, want %s", got.Private, priv)
	}
}

func TestEntity_PrivateOmittedWhenNil(t *testing.T) {
	e := Entity{
		ID:       "default",
		Plugin:   "plugin-amcrest",
		DeviceID: "default",
		Type:     "camera",
		State:    Camera{},
	}
	raw, err := json.Marshal(e)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if strings.Contains(string(raw), `"private"`) {
		t.Fatalf("expected private key omitted, got %s", raw)
	}
}

func TestEntity_PrivateCoexistsWithStateAndMeta(t *testing.T) {
	priv := json.RawMessage(`{"token":"abc"}`)
	raw := []byte(`{
		"id":"demo",
		"plugin":"plugin-clean",
		"deviceID":"demo",
		"type":"light",
		"name":"Demo",
		"meta":{"room":"kitchen"},
		"private":{"token":"abc"},
		"state":{"power":true,"brightness":128}
	}`)
	var got Entity
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !bytes.Equal(got.Private, priv) {
		t.Fatalf("private mismatch: got %s, want %s", got.Private, priv)
	}
	if len(got.Meta) == 0 || string(got.Meta["room"]) != `"kitchen"` {
		t.Fatalf("meta not preserved: %+v", got.Meta)
	}
	if got.State == nil {
		t.Fatalf("state not hydrated")
	}
}
