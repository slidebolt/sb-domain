package domain

type Alarm struct {
AlarmState      string `json:"alarmState"` // disarmed, armed_home, armed_away, armed_night, triggered
CodeArmRequired bool   `json:"codeArmRequired"`
}
