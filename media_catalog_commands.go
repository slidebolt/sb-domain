package domain

type MediaProviderConnect struct{}

func (MediaProviderConnect) ActionName() string { return "media_provider_connect" }

type MediaProviderDisconnect struct{}

func (MediaProviderDisconnect) ActionName() string { return "media_provider_disconnect" }

type MediaSeek struct {
	PositionMS int64 `json:"positionMS"`
}

func (MediaSeek) ActionName() string { return "media_seek" }
