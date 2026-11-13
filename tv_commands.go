package domain

type TVRecordCurrent struct {
	Recorder string `json:"recorder,omitempty"`
}

func (TVRecordCurrent) ActionName() string { return "tv_record_current" }

type TVRecordSeries struct {
	NewOnly        bool `json:"newOnly,omitempty"`
	IncludeRepeats bool `json:"includeRepeats,omitempty"`
}

func (TVRecordSeries) ActionName() string { return "tv_record_series" }

type TVRecordCancel struct{}

func (TVRecordCancel) ActionName() string { return "tv_record_cancel" }
