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
