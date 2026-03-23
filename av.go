package domain

type MediaPlayer struct {
State         string   `json:"state"` // playing, paused, idle, off
VolumeLevel   float64  `json:"volumeLevel"`
IsVolumeMuted bool     `json:"isVolumeMuted"`
MediaTitle    string   `json:"mediaTitle,omitempty"`
MediaArtist   string   `json:"mediaArtist,omitempty"`
Source        string   `json:"source,omitempty"`
SourceList    []string `json:"sourceList,omitempty"`
}

type Remote struct {
IsOn            bool     `json:"isOn"`
ActivityList    []string `json:"activityList,omitempty"`
CurrentActivity string   `json:"currentActivity,omitempty"`
}

type Camera struct {
IsStreaming     bool   `json:"isStreaming"`
IsRecording     bool   `json:"isRecording"`
MotionDetection bool   `json:"motionDetection"`
StreamSource    string `json:"streamSource,omitempty"`
SnapshotURL     string `json:"snapshotURL,omitempty"`
}
