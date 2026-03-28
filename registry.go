package domain

import (
	"reflect"
	"sync"
)

var (
	registryMu sync.RWMutex
	registry   = map[string]reflect.Type{}

	commandRegistryMu sync.RWMutex
	commandRegistry   = map[string]reflect.Type{}
)

// Register maps a type name (matching Entity.Type) to a concrete Go type.
// Built-in types are registered automatically. Plugins call this for custom types.
func Register(typeName string, zero any) {
	registryMu.Lock()
	registry[typeName] = reflect.TypeOf(zero)
	registryMu.Unlock()
}

// RegisterCommand maps an action name to a concrete command type.
// Built-in commands are registered automatically. Plugins call this for custom commands.
func RegisterCommand(actionName string, zero any) {
	commandRegistryMu.Lock()
	commandRegistry[actionName] = reflect.TypeOf(zero)
	commandRegistryMu.Unlock()
}

// LookupCommand returns the registered type for an action name.
func LookupCommand(actionName string) (reflect.Type, bool) {
	commandRegistryMu.RLock()
	t, ok := commandRegistry[actionName]
	commandRegistryMu.RUnlock()
	return t, ok
}

func lookupType(typeName string) (reflect.Type, bool) {
	registryMu.RLock()
	t, ok := registry[typeName]
	registryMu.RUnlock()
	return t, ok
}

func init() {
	Register("light", Light{})
	Register("light_panel", LightPanel{})
	Register("light_strip", LightStrip{})
	Register("switch", Switch{})
	Register("fan", Fan{})
	Register("cover", Cover{})
	Register("lock", Lock{})
	Register("button", Button{})
	Register("sensor", Sensor{})
	Register("binary_sensor", BinarySensor{})
	Register("number", Number{})
	Register("select", Select{})
	Register("text", Text{})
	Register("climate", Climate{})

	RegisterCommand("light_turn_on", LightTurnOn{})
	RegisterCommand("light_turn_off", LightTurnOff{})
	RegisterCommand("light_set_brightness", LightSetBrightness{})
	RegisterCommand("light_set_color_temp", LightSetColorTemp{})
	RegisterCommand("light_set_rgb", LightSetRGB{})
	RegisterCommand("light_set_rgbw", LightSetRGBW{})
	RegisterCommand("light_set_rgbww", LightSetRGBWW{})
	RegisterCommand("light_set_hs", LightSetHS{})
	RegisterCommand("light_set_xy", LightSetXY{})
	RegisterCommand("light_set_white", LightSetWhite{})
	RegisterCommand("light_set_effect", LightSetEffect{})
	RegisterCommand("lightstrip_set_segments", LightstripSetSegments{})

	RegisterCommand("switch_turn_on", SwitchTurnOn{})
	RegisterCommand("switch_turn_off", SwitchTurnOff{})
	RegisterCommand("switch_toggle", SwitchToggle{})

	RegisterCommand("fan_turn_on", FanTurnOn{})
	RegisterCommand("fan_turn_off", FanTurnOff{})
	RegisterCommand("fan_set_speed", FanSetSpeed{})

	RegisterCommand("cover_open", CoverOpen{})
	RegisterCommand("cover_close", CoverClose{})
	RegisterCommand("cover_set_position", CoverSetPosition{})

	RegisterCommand("lock_lock", LockLock{})
	RegisterCommand("lock_unlock", LockUnlock{})

	RegisterCommand("button_press", ButtonPress{})

	RegisterCommand("number_set_value", NumberSetValue{})
	RegisterCommand("select_option", SelectOption{})
	RegisterCommand("text_set_value", TextSetValue{})

	RegisterCommand("climate_set_mode", ClimateSetMode{})
	RegisterCommand("climate_set_temperature", ClimateSetTemperature{})
}

func init() {
	Register("alarm", Alarm{})
	Register("camera", Camera{})
	Register("valve", Valve{})
	Register("siren", Siren{})
	Register("humidifier", Humidifier{})
	Register("media_player", MediaPlayer{})
	Register("remote", Remote{})
	Register("event", Event{})
	Register("phone", Phone{})

	// Alarm
	RegisterCommand("alarm_arm_home", AlarmArmHome{})
	RegisterCommand("alarm_arm_away", AlarmArmAway{})
	RegisterCommand("alarm_arm_night", AlarmArmNight{})
	RegisterCommand("alarm_disarm", AlarmDisarm{})

	// Camera
	RegisterCommand("camera_record_start", CameraRecordStart{})
	RegisterCommand("camera_record_stop", CameraRecordStop{})
	RegisterCommand("camera_enable_motion", CameraEnableMotion{})
	RegisterCommand("camera_disable_motion", CameraDisableMotion{})

	// Valve
	RegisterCommand("valve_open", ValveOpen{})
	RegisterCommand("valve_close", ValveClose{})
	RegisterCommand("valve_set_position", ValveSetPosition{})

	// Siren
	RegisterCommand("siren_turn_on", SirenTurnOn{})
	RegisterCommand("siren_turn_off", SirenTurnOff{})
	RegisterCommand("siren_set_tone", SirenSetTone{})

	// Humidifier
	RegisterCommand("humidifier_turn_on", HumidifierTurnOn{})
	RegisterCommand("humidifier_turn_off", HumidifierTurnOff{})
	RegisterCommand("humidifier_set_humidity", HumidifierSetHumidity{})
	RegisterCommand("humidifier_set_mode", HumidifierSetMode{})

	// MediaPlayer
	RegisterCommand("media_play", MediaPlay{})
	RegisterCommand("media_pause", MediaPause{})
	RegisterCommand("media_stop", MediaStop{})
	RegisterCommand("media_next_track", MediaNextTrack{})
	RegisterCommand("media_previous_track", MediaPreviousTrack{})
	RegisterCommand("media_set_volume", MediaSetVolume{})
	RegisterCommand("media_mute", MediaMute{})
	RegisterCommand("media_select_source", MediaSelectSource{})

	// Remote
	RegisterCommand("remote_turn_on", RemoteTurnOn{})
	RegisterCommand("remote_turn_off", RemoteTurnOff{})
	RegisterCommand("remote_set_activity", RemoteSetActivity{})
	RegisterCommand("remote_send_command", RemoteSendCommand{})
	RegisterCommand("phone_register_push_token", PhoneRegisterPushToken{})
	RegisterCommand("phone_send_notification", PhoneSendNotification{})
	RegisterCommand("phone_send_data_message", PhoneSendDataMessage{})
}
