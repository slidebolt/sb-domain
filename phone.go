package domain

type Phone struct {
	Platform               string  `json:"platform,omitempty"`
	Online                 bool    `json:"online"`
	PushTokenConfigured    bool    `json:"pushTokenConfigured"`
	LastSeen               string  `json:"lastSeen,omitempty"`
	BatteryLevel           float64 `json:"batteryLevel,omitempty"`
	IsCharging             bool    `json:"isCharging"`
	NotificationPermission string  `json:"notificationPermission,omitempty"`
	DeviceModel            string  `json:"deviceModel,omitempty"`
	AppVersion             string  `json:"appVersion,omitempty"`
	LastNotificationAt     string  `json:"lastNotificationAt,omitempty"`
	LastNotificationError  string  `json:"lastNotificationError,omitempty"`
}
