package domain

type AlarmArmHome struct{ Code *string }

func (AlarmArmHome) ActionName() string { return "alarm_arm_home" }

type AlarmArmAway struct{ Code *string }

func (AlarmArmAway) ActionName() string { return "alarm_arm_away" }

type AlarmArmNight struct{ Code *string }

func (AlarmArmNight) ActionName() string { return "alarm_arm_night" }

type AlarmDisarm struct{ Code *string }

func (AlarmDisarm) ActionName() string { return "alarm_disarm" }
