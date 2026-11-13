package domain

import "time"

// MediaItemRef identifies provider-owned media without exposing credentials or
// assuming that the provider's files are locally accessible.
type MediaItemRef struct {
	Provider  string `json:"provider"`
	ServerID  string `json:"serverID"`
	LibraryID string `json:"libraryID,omitempty"`
	ItemID    string `json:"itemID"`
}

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
