// Package domain defines the core data objects for the SlideBolt system.
// These are plain structs with JSON serialization. No commands, no events,
// no storage, no routing. Just the shape of things.
package domain

import (
	"encoding/json"
	"reflect"
)

// decodeLabels unmarshals a raw label map into canonical array format:
// "room": ["kitchen"]. Accepts both string and []string values in JSON.
func decodeLabels(raw map[string]json.RawMessage) map[string][]string {
	if len(raw) == 0 {
		return nil
	}
	out := make(map[string][]string, len(raw))
	for k, v := range raw {
		var ss []string
		if json.Unmarshal(v, &ss) == nil {
			out[k] = ss
			continue
		}
		var s string
		if json.Unmarshal(v, &s) == nil {
			out[k] = []string{s}
		}
	}
	return out
}

// PluginKey identifies a plugin.
type PluginKey struct {
	ID string
}

func (k PluginKey) Key() string { return k.ID }

// DeviceKey identifies a device within a plugin.
type DeviceKey struct {
	Plugin string
	ID     string
}

func (k DeviceKey) Key() string { return k.Plugin + "." + k.ID }

// EntityKey identifies an entity within a device.
type EntityKey struct {
	Plugin   string
	DeviceID string
	ID       string
}

func (k EntityKey) Key() string { return k.Plugin + "." + k.DeviceID + "." + k.ID }

// DeviceProfile holds user-owned properties that survive rediscovery.
// Plugin-driven discovery must never overwrite these fields.
type DeviceProfile struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

// Device represents a physical or virtual device discovered by a plugin.
type Device struct {
	ID       string                     `json:"id"`
	Plugin   string                     `json:"plugin"`
	Name     string                     `json:"name"`
	Labels   map[string][]string        `json:"labels,omitempty"`
	Profile  *DeviceProfile             `json:"profile,omitempty"`
	Meta     map[string]json.RawMessage `json:"meta,omitempty"`
	Entities []Entity                   `json:"entities"`
}

func (d Device) Key() string { return d.Plugin + "." + d.ID }

// EntityProfile holds user-owned properties that survive rediscovery.
// Plugin-driven discovery must never overwrite these fields.
type EntityProfile struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

// Entity represents a single controllable or observable unit within a device.
// Commands lists the action names this entity accepts (e.g. "turn_on", "set_brightness").
// Set by the plugin during discovery. An empty list means no command filtering is applied.
//
// Target is an optional opaque query that, when set, makes this a virtual entity.
// The core resolves Target into member entities and fans commands out to them.
// The domain does not interpret Target — it is passed through to the storage layer.
type Entity struct {
	ID       string                     `json:"id"`
	Plugin   string                     `json:"plugin"`
	DeviceID string                     `json:"deviceID"`
	Type     string                     `json:"type"`
	Name     string                     `json:"name"`
	Labels   map[string][]string        `json:"labels,omitempty"`
	Profile  *EntityProfile             `json:"profile,omitempty"`
	Meta     map[string]json.RawMessage `json:"meta,omitempty"`
	Commands []string                   `json:"commands,omitempty"`
	State    any                        `json:"state"`
	Target   json.RawMessage            `json:"target,omitempty"`
}

func (e Entity) Key() string { return e.Plugin + "." + e.DeviceID + "." + e.ID }

// UnmarshalJSON hydrates State into the concrete registered type based on Type.
func (e *Entity) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID       string                     `json:"id"`
		Plugin   string                     `json:"plugin"`
		DeviceID string                     `json:"deviceID"`
		Type     string                     `json:"type"`
		Name     string                     `json:"name"`
		Commands []string                   `json:"commands,omitempty"`
		Labels   map[string]json.RawMessage `json:"labels,omitempty"`
		Profile  *EntityProfile             `json:"profile,omitempty"`
		Meta     map[string]json.RawMessage `json:"meta,omitempty"`
		State    json.RawMessage            `json:"state,omitempty"`
		Target   json.RawMessage            `json:"target,omitempty"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	e.ID = aux.ID
	e.Plugin = aux.Plugin
	e.DeviceID = aux.DeviceID
	e.Type = aux.Type
	e.Name = aux.Name
	e.Commands = aux.Commands
	e.Labels = decodeLabels(aux.Labels)
	e.Profile = aux.Profile
	e.Meta = aux.Meta
	e.Target = aux.Target

	if len(aux.State) == 0 || string(aux.State) == "null" {
		return nil
	}

	if typ, ok := lookupType(aux.Type); ok {
		v := reflect.New(typ).Interface()
		if err := json.Unmarshal(aux.State, v); err != nil {
			return err
		}
		e.State = reflect.ValueOf(v).Elem().Interface()
	} else {
		// Unknown type — keep as raw map so it's still accessible.
		var m any
		json.Unmarshal(aux.State, &m)
		e.State = m
	}

	return nil
}
