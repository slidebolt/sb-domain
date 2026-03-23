package domain

// MediaPlayer commands
type MediaPlay struct{}

func (MediaPlay) ActionName() string { return "media_play" }

type MediaPause struct{}

func (MediaPause) ActionName() string { return "media_pause" }

type MediaStop struct{}

func (MediaStop) ActionName() string { return "media_stop" }

type MediaNextTrack struct{}

func (MediaNextTrack) ActionName() string { return "media_next_track" }

type MediaPreviousTrack struct{}

func (MediaPreviousTrack) ActionName() string { return "media_previous_track" }

type MediaSetVolume struct{ VolumeLevel float64 }

func (MediaSetVolume) ActionName() string { return "media_set_volume" }

type MediaMute struct{ Mute bool }

func (MediaMute) ActionName() string { return "media_mute" }

type MediaSelectSource struct{ Source string }

func (MediaSelectSource) ActionName() string { return "media_select_source" }

// Remote commands
type RemoteTurnOn struct{}

func (RemoteTurnOn) ActionName() string { return "remote_turn_on" }

type RemoteTurnOff struct{}

func (RemoteTurnOff) ActionName() string { return "remote_turn_off" }

type RemoteSetActivity struct{ Activity string }

func (RemoteSetActivity) ActionName() string { return "remote_set_activity" }

type RemoteSendCommand struct{ Command string }

func (RemoteSendCommand) ActionName() string { return "remote_send_command" }

// Camera commands
type CameraRecordStart struct{}

func (CameraRecordStart) ActionName() string { return "camera_record_start" }

type CameraRecordStop struct{}

func (CameraRecordStop) ActionName() string { return "camera_record_stop" }

type CameraEnableMotion struct{}

func (CameraEnableMotion) ActionName() string { return "camera_enable_motion" }

type CameraDisableMotion struct{}

func (CameraDisableMotion) ActionName() string { return "camera_disable_motion" }
