package domain

import (
	"encoding/json"
	"time"
)

// MediaItemRef identifies provider-owned media without exposing credentials or
// assuming that the provider's files are locally accessible.
type MediaItemRef struct {
	Provider  string `json:"provider"`
	ServerID  string `json:"serverID"`
	LibraryID string `json:"libraryID,omitempty"`
	ItemID    string `json:"itemID"`
}

type MediaAsset struct {
	Kind string `json:"kind"`
	Ref  string `json:"ref"`
}

type MediaItem struct {
	Ref        MediaItemRef               `json:"ref"`
	Kind       string                     `json:"kind"`
	Title      string                     `json:"title"`
	Playable   bool                       `json:"playable"`
	DurationMS int64                      `json:"durationMS,omitempty"`
	Parent     *MediaItemRef              `json:"parent,omitempty"`
	Metadata   map[string]json.RawMessage `json:"metadata,omitempty"`
	Assets     []MediaAsset               `json:"assets,omitempty"`
}

type MediaActivity struct {
	Item       MediaItemRef `json:"item"`
	PositionMS int64        `json:"positionMS,omitempty"`
	Completed  bool         `json:"completed,omitempty"`
	LastPlayed *time.Time   `json:"lastPlayed,omitempty"`
}

type PlayableSource struct {
	Kind   string        `json:"kind"`
	Item   *MediaItemRef `json:"item,omitempty"`
	Entity string        `json:"entity,omitempty"`
}

type MediaSource struct {
	Kind      string `json:"kind"`
	Available bool   `json:"available"`
	URL       string `json:"url,omitempty"`
	Format    string `json:"format,omitempty"`
	Error     string `json:"error,omitempty"`
}

type Recorder struct {
	Status       string          `json:"status"`
	Source       *PlayableSource `json:"source,omitempty"`
	Output       *MediaItemRef   `json:"output,omitempty"`
	StartedAt    *time.Time      `json:"startedAt,omitempty"`
	BytesWritten int64           `json:"bytesWritten,omitempty"`
	LastError    string          `json:"lastError,omitempty"`
}

type RecordStart struct {
	Source      PlayableSource             `json:"source"`
	Profile     string                     `json:"profile,omitempty"`
	Destination string                     `json:"destination,omitempty"`
	Metadata    map[string]json.RawMessage `json:"metadata,omitempty"`
}

func (RecordStart) ActionName() string { return "record_start" }

type RecordStop struct{}

func (RecordStop) ActionName() string { return "record_stop" }

type RecordCancel struct {
	Reference string `json:"reference,omitempty"`
}

func (RecordCancel) ActionName() string { return "record_cancel" }

type MediaQuery struct {
	Kind       string                     `json:"kind,omitempty"`
	Title      string                     `json:"title,omitempty"`
	Attributes map[string]json.RawMessage `json:"attributes,omitempty"`
}

type RecordingRule struct {
	Enabled        bool       `json:"enabled"`
	Query          MediaQuery `json:"query"`
	Selection      string     `json:"selection"`
	Limit          int        `json:"limit,omitempty"`
	Recorder       string     `json:"recorder"`
	Profile        string     `json:"profile,omitempty"`
	ScheduledCount int        `json:"scheduledCount,omitempty"`
	LastScannedAt  *time.Time `json:"lastScannedAt,omitempty"`
	LastError      string     `json:"lastError,omitempty"`
}

// RecordingRuleCreate asks a recorder to persist and apply a provider-neutral
// catalog query. The recorder plugin owns matching that query to its catalog.
type RecordingRuleCreate struct {
	ID   string        `json:"id,omitempty"`
	Name string        `json:"name,omitempty"`
	Rule RecordingRule `json:"rule"`
}

func (RecordingRuleCreate) ActionName() string { return "recording_rule_create" }

type RecordingRuleDelete struct {
	Reference string `json:"reference"`
}

func (RecordingRuleDelete) ActionName() string { return "recording_rule_delete" }

type MediaServer struct {
	Connected  bool       `json:"connected"`
	Name       string     `json:"name,omitempty"`
	Version    string     `json:"version,omitempty"`
	Owned      bool       `json:"owned,omitempty"`
	LastSyncAt *time.Time `json:"lastSyncAt,omitempty"`
	LastError  string     `json:"lastError,omitempty"`
}

type MediaLibrary struct {
	ServerID  string     `json:"serverID"`
	LibraryID string     `json:"libraryID"`
	Kind      string     `json:"kind"` // movie, show, music, photo, mixed
	ItemCount int        `json:"itemCount,omitempty"`
	CanBrowse bool       `json:"canBrowse"`
	CanSearch bool       `json:"canSearch"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	LastError string     `json:"lastError,omitempty"`
}

type MediaSession struct {
	ServerID   string        `json:"serverID"`
	PlayerID   string        `json:"playerID,omitempty"`
	User       string        `json:"user,omitempty"`
	Item       *MediaItemRef `json:"item,omitempty"`
	Title      string        `json:"title,omitempty"`
	State      string        `json:"state"`
	PositionMS int64         `json:"positionMS,omitempty"`
	DurationMS int64         `json:"durationMS,omitempty"`
	LastError  string        `json:"lastError,omitempty"`
}

type MediaProviderConnection struct {
	Status    string     `json:"status"` // disconnected, pending, connected, error
	ClaimURL  string     `json:"claimURL,omitempty"`
	Code      string     `json:"code,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	LastError string     `json:"lastError,omitempty"`
}
