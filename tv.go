package domain

import "time"

type TVProgram struct {
	AiringID     string        `json:"airingID"`
	SeriesID     string        `json:"seriesID,omitempty"`
	Title        string        `json:"title"`
	EpisodeTitle string        `json:"episodeTitle,omitempty"`
	Season       int           `json:"season,omitempty"`
	Episode      int           `json:"episode,omitempty"`
	Synopsis     string        `json:"synopsis,omitempty"`
	StartsAt     time.Time     `json:"startsAt"`
	EndsAt       time.Time     `json:"endsAt"`
	IsNew        bool          `json:"isNew,omitempty"`
	ProviderRef  *MediaItemRef `json:"providerRef,omitempty"`
}

type TVChannel struct {
	Number         string     `json:"number"`
	Name           string     `json:"name"`
	Available      bool       `json:"available"`
	DRM            bool       `json:"drm,omitempty"`
	HD             bool       `json:"hd,omitempty"`
	VideoCodec     string     `json:"videoCodec,omitempty"`
	AudioCodec     string     `json:"audioCodec,omitempty"`
	SignalStrength int        `json:"signalStrength,omitempty"`
	SignalQuality  int        `json:"signalQuality,omitempty"`
	CurrentProgram *TVProgram `json:"currentProgram,omitempty"`
	NextProgram    *TVProgram `json:"nextProgram,omitempty"`
	GuideUpdatedAt *time.Time `json:"guideUpdatedAt,omitempty"`
	GuideError     string     `json:"guideError,omitempty"`
}

type TVTuner struct {
	Available bool   `json:"available"`
	DeviceID  string `json:"deviceID"`
	Model     string `json:"model,omitempty"`
	Tuner     int    `json:"tuner"`
	LastError string `json:"lastError,omitempty"`
}

type RecordingStartInput struct {
	ChannelKey string     `json:"channel"`
	Program    *TVProgram `json:"program,omitempty"`
	Profile    string     `json:"profile,omitempty"`
}

type RecordingSwitch struct {
	Power        bool       `json:"power"`
	ChannelKey   string     `json:"channel,omitempty"`
	Program      *TVProgram `json:"program,omitempty"`
	FileDeviceID string     `json:"fileDeviceID,omitempty"`
	StartedAt    *time.Time `json:"startedAt,omitempty"`
	BytesWritten int64      `json:"bytesWritten,omitempty"`
	LastError    string     `json:"lastError,omitempty"`
}

type RecordingStatus string

const (
	RecordingStatusRecording RecordingStatus = "recording"
	RecordingStatusEncoding  RecordingStatus = "encoding"
	RecordingStatusCompleted RecordingStatus = "completed"
	RecordingStatusFailed    RecordingStatus = "failed"
)

type RecordingFile struct {
	Status             RecordingStatus `json:"status"`
	RelativePath       string          `json:"relativePath"`
	PartialPath        string          `json:"partialPath,omitempty"`
	Format             string          `json:"format"`
	Profile            string          `json:"profile"`
	Channel            TVChannel       `json:"channel"`
	Program            *TVProgram      `json:"program,omitempty"`
	StartedAt          time.Time       `json:"startedAt"`
	FinishedAt         *time.Time      `json:"finishedAt,omitempty"`
	EncodingStartedAt  *time.Time      `json:"encodingStartedAt,omitempty"`
	EncodingFinishedAt *time.Time      `json:"encodingFinishedAt,omitempty"`
	SourceBytes        int64           `json:"sourceBytes,omitempty"`
	BytesWritten       int64           `json:"bytesWritten"`
	Error              string          `json:"error,omitempty"`
}

type TVRecordingRule struct {
	Enabled        bool       `json:"enabled"`
	SeriesID       string     `json:"seriesID"`
	Title          string     `json:"title"`
	NewOnly        bool       `json:"newOnly"`
	ChannelNumbers []string   `json:"channelNumbers,omitempty"`
	Profile        string     `json:"profile,omitempty"`
	LastScannedAt  *time.Time `json:"lastScannedAt,omitempty"`
	LastError      string     `json:"lastError,omitempty"`
	Provider       string     `json:"provider,omitempty"`
	ExternalID     string     `json:"externalID,omitempty"`
	ScheduledCount int        `json:"scheduledCount,omitempty"`
}

type TVRecording struct {
	Status      string        `json:"status"` // scheduled, recording, processing, completed, failed, canceled
	Provider    string        `json:"provider"`
	ExternalID  string        `json:"externalID,omitempty"`
	Channel     TVChannel     `json:"channel"`
	Program     TVProgram     `json:"program"`
	RuleID      string        `json:"ruleID,omitempty"`
	StartsAt    time.Time     `json:"startsAt"`
	EndsAt      time.Time     `json:"endsAt"`
	LibraryItem *MediaItemRef `json:"libraryItem,omitempty"`
	LastError   string        `json:"lastError,omitempty"`
}
