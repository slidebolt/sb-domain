package domain

type TimerEnable struct{}

func (TimerEnable) ActionName() string { return "timer_enable" }

type TimerDisable struct{}

func (TimerDisable) ActionName() string { return "timer_disable" }
