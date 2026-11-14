package domain

import "time"

type MediaPlayer struct {
	State         string        `json:"state"` // playing, paused, idle, off
	VolumeLevel   float64       `json:"volumeLevel"`
	IsVolumeMuted bool          `json:"isVolumeMuted"`
	MediaTitle    string        `json:"mediaTitle,omitempty"`
	MediaArtist   string        `json:"mediaArtist,omitempty"`
	Source        string        `json:"source,omitempty"`
	SourceList    []string      `json:"sourceList,omitempty"`
	ServerID      string        `json:"serverID,omitempty"`
	Item          *MediaItemRef `json:"item,omitempty"`
	PositionMS    int64         `json:"positionMS,omitempty"`
	DurationMS    int64         `json:"durationMS,omitempty"`
	Available     bool          `json:"available,omitempty"`
	LastError     string        `json:"lastError,omitempty"`
}

type Remote struct {
	IsOn            bool     `json:"isOn"`
	ActivityList    []string `json:"activityList,omitempty"`
	CurrentActivity string   `json:"currentActivity,omitempty"`
}

type Camera struct {
	IsStreaming      bool     `json:"isStreaming"`
	IsRecording      bool     `json:"isRecording"`
	MotionDetection  bool     `json:"motionDetection"`
	Available        bool     `json:"available"`
	Enabled          bool     `json:"enabled"`
	DetectionEnabled bool     `json:"detectionEnabled"`
	SnapshotsEnabled bool     `json:"snapshotsEnabled"`
	Zones            []string `json:"zones,omitempty"`
	LastError        string   `json:"lastError,omitempty"`
	StreamSource     string   `json:"streamSource,omitempty"`
	SnapshotURL      string   `json:"snapshotURL,omitempty"`
}

// MediaStream describes a stream served by sb-media. Kind is audio today and
// will also cover video and image. URL carries an access token, so treat it as
// a secret and keep it out of logs.
type MediaStream struct {
	Status    string         `json:"status"`
	Source    PlayableSource `json:"source"`
	Target    string         `json:"target,omitempty"`
	Kind      string         `json:"kind"`
	URL       string         `json:"url,omitempty"`
	Format    string         `json:"format,omitempty"`
	Viewers   int            `json:"viewers"`
	StartedAt *time.Time     `json:"startedAt,omitempty"`
	Error     string         `json:"error,omitempty"`
}
