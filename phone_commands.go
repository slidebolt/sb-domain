package domain

import "fmt"

type PhoneRegisterPushToken struct {
	Token                  string  `json:"token"`
	Platform               string  `json:"platform,omitempty"`
	LastSeen               string  `json:"lastSeen,omitempty"`
	BatteryLevel           float64 `json:"batteryLevel,omitempty"`
	IsCharging             bool    `json:"isCharging"`
	NotificationPermission string  `json:"notificationPermission,omitempty"`
	DeviceModel            string  `json:"deviceModel,omitempty"`
	AppVersion             string  `json:"appVersion,omitempty"`
}

func (PhoneRegisterPushToken) ActionName() string { return "phone_register_push_token" }

func (c PhoneRegisterPushToken) Validate() error {
	if c.Token == "" {
		return fmt.Errorf("phone token must not be empty")
	}
	if c.BatteryLevel < 0 || c.BatteryLevel > 100 {
		return fmt.Errorf("batteryLevel %.2f out of range [0,100]", c.BatteryLevel)
	}
	return nil
}

type PhoneSendNotification struct {
	Title    string            `json:"title,omitempty"`
	Body     string            `json:"body,omitempty"`
	Data     map[string]string `json:"data,omitempty"`
	ImageURL string            `json:"imageURL,omitempty"`
}

func (PhoneSendNotification) ActionName() string { return "phone_send_notification" }

func (c PhoneSendNotification) Validate() error {
	if c.Title == "" && c.Body == "" && len(c.Data) == 0 {
		return fmt.Errorf("notification must include title, body, or data")
	}
	return nil
}

type PhoneSendDataMessage struct {
	Data map[string]string `json:"data"`
}

func (PhoneSendDataMessage) ActionName() string { return "phone_send_data_message" }

func (c PhoneSendDataMessage) Validate() error {
	if len(c.Data) == 0 {
		return fmt.Errorf("data message must not be empty")
	}
	return nil
}
