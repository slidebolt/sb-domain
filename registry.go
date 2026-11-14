package domain

import (
	"reflect"
	"sort"
	"strings"
	"sync"
)

type RegisteredDefinition struct {
	Type        string
	Name        string
	Description string
	GoType      reflect.Type
}

func displayName(value string) string {
	parts := strings.Split(value, "_")
	for i := range parts {
		if parts[i] != "" {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, " ")
}

var entityDescriptions = map[string]string{
	"alarm": "Security alarm state and arming controls.", "agent": "A software agent that accepts prompts and reports results.",
	"binary_sensor": "A boolean observation such as occupancy, contact, or connectivity.", "button": "A stateless control that performs one action when pressed.",
	"camera": "Camera availability, detection, snapshot, streaming, and recording state.", "climate": "Heating and cooling mode, temperature, and set-point state.",
	"cover": "A movable cover such as a door, shade, or blind.", "event": "The latest occurrence of one or more event types.",
	"group": "A dynamic collection of entities with aggregate state and controls.",
	"fan":   "Fan power and speed state.", "humidifier": "Humidifier power, target humidity, and operating mode.",
	"light": "A controllable light with power, brightness, color, and effect capabilities.", "light_panel": "A light composed of independently addressable panel elements.",
	"light_strip": "A light composed of independently addressable strip segments.", "lock": "Lock state and lock or unlock controls.",
	"media_library": "A browsable collection of media supplied by a media server.", "media_player": "Playback, volume, source, and current-media state.",
	"media_provider_connection": "Connection and authentication state for a media provider.", "media_server": "A discoverable server that exposes media libraries.",
	"media_session": "A current or recent media playback session.", "media_stream": "An addressable audio, video, or image stream.",
	"number": "A bounded numeric setting.", "pbx": "Telephone exchange state and administration controls.",
	"phone": "A telephone endpoint and its current call state.", "recording_file": "A completed recording and its output media file.",
	"recording_switch": "The current recording control and capture state.", "remote": "Remote-control power and activity state.",
	"select": "A setting constrained to one of a declared list of options.", "sensor": "A scalar or structured observation.",
	"sip_call": "A SIP call and its lifecycle state.", "sip_endpoint": "A SIP endpoint registration and reachability state.",
	"sip_messages": "Message history and sending controls for a SIP endpoint.", "sip_participant": "A participant in a SIP call.",
	"sip_trunk": "A SIP trunk registration and connectivity state.", "siren": "Siren power and tone state.",
	"speech": "Speech synthesis or recognition state.", "switch": "A generic controllable on or off component.",
	"text": "A constrained text setting.", "timer": "A scheduled on and off control definition.",
	"time": "A local calendar and clock observation.", "location": "A geographic position and time-zone observation.",
	"transcript": "Incremental or completed recognized speech text.", "tv_channel": "A television channel and its current or upcoming programme.",
	"tv_recording": "A scheduled or active television recording.", "tv_recording_rule": "A reusable rule that schedules matching television episodes.",
	"tv_tuner": "A television tuner and its selected channel or recording state.", "valve": "Valve open, closed, and position state.",
}

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

func EntityDefinitions() []RegisteredDefinition {
	registryMu.RLock()
	defer registryMu.RUnlock()
	out := make([]RegisteredDefinition, 0, len(registry))
	for name, typ := range registry {
		description := entityDescriptions[name]
		if description == "" {
			description = "Shared domain state contract for " + displayName(name) + "."
		}
		out = append(out, RegisteredDefinition{Type: name, Name: displayName(name), Description: description, GoType: typ})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Type < out[j].Type })
	return out
}

func CommandDefinitions() []RegisteredDefinition {
	commandRegistryMu.RLock()
	defer commandRegistryMu.RUnlock()
	out := make([]RegisteredDefinition, 0, len(commandRegistry))
	for name, typ := range commandRegistry {
		display := displayName(name)
		out = append(out, RegisteredDefinition{Type: name, Name: display, Description: "Executes the " + strings.ToLower(display) + " action.", GoType: typ})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Type < out[j].Type })
	return out
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
	Register("timer", Timer{})
	Register("tv_channel", TVChannel{})
	Register("tv_tuner", TVTuner{})
	Register("recording_switch", RecordingSwitch{})
	Register("recording_file", RecordingFile{})
	Register("tv_recording_rule", TVRecordingRule{})
	Register("tv_recording", TVRecording{})
	Register("media_server", MediaServer{})
	Register("media_library", MediaLibrary{})
	Register("media_session", MediaSession{})
	Register("media_provider_connection", MediaProviderConnection{})

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
	RegisterCommand("binary_sensor_turn_on", BinarySensorTurnOn{})
	RegisterCommand("binary_sensor_turn_off", BinarySensorTurnOff{})

	RegisterCommand("number_set_value", NumberSetValue{})
	RegisterCommand("select_option", SelectOption{})
	RegisterCommand("text_set_value", TextSetValue{})

	RegisterCommand("climate_set_mode", ClimateSetMode{})
	RegisterCommand("climate_set_temperature", ClimateSetTemperature{})
	RegisterCommand("timer_enable", TimerEnable{})
	RegisterCommand("timer_disable", TimerDisable{})
	RegisterCommand("tv_record_current", TVRecordCurrent{})
	RegisterCommand("tv_record_series", TVRecordSeries{})
	RegisterCommand("tv_record_airing", TVRecordAiring{})
	RegisterCommand("tv_record_cancel", TVRecordCancel{})
	RegisterCommand("media_provider_connect", MediaProviderConnect{})
	RegisterCommand("media_provider_disconnect", MediaProviderDisconnect{})
	RegisterCommand("media_play_item", MediaPlayItem{})
	RegisterCommand("media_seek", MediaSeek{})
}

func init() {
	Register("alarm", Alarm{})
	Register("camera", Camera{})
	Register("valve", Valve{})
	Register("siren", Siren{})
	Register("humidifier", Humidifier{})
	Register("media_player", MediaPlayer{})
	Register("media_stream", MediaStream{})
	Register("pbx", PBX{})
	Register("sip_trunk", SIPTrunk{})
	Register("sip_endpoint", SIPEndpoint{})
	Register("sip_call", SIPCall{})
	Register("sip_participant", SIPParticipant{})
	Register("agent", Agent{})
	Register("transcript", Transcript{})
	Register("speech", Speech{})
	Register("sip_messages", SIPMessages{})
	Register("remote", Remote{})
	Register("event", Event{})
	Register("phone", Phone{})
	Register("group", Group{})
	Register("time", Time{})
	Register("location", Location{})

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

	// MediaStream
	RegisterCommand("media_stream_start", MediaStreamStart{})
	RegisterCommand("media_stream_stop", MediaStreamStop{})

	// SIP
	RegisterCommand("pbx_reload", PBXReload{})
	RegisterCommand("sip_call", SIPDial{})
	RegisterCommand("sip_hangup", SIPHangup{})
	RegisterCommand("sip_play", SIPPlay{})
	RegisterCommand("sip_mute", SIPMute{})
	RegisterCommand("sip_unmute", SIPUnmute{})
	RegisterCommand("sip_hold", SIPHold{})
	RegisterCommand("sip_unhold", SIPUnhold{})
	RegisterCommand("sip_dtmf", SIPDTMF{})
	RegisterCommand("sip_remove", SIPRemove{})
	RegisterCommand("sip_transcribe_start", SIPTranscribeStart{})
	RegisterCommand("sip_transcribe_stop", SIPTranscribeStop{})
	RegisterCommand("sip_say", SIPSay{})
	RegisterCommand("sip_message_send", SIPMessageSend{})

	// Agent
	RegisterCommand("agent_prompt", AgentPrompt{})
	RegisterCommand("agent_cancel", AgentCancel{})

	// Remote
	RegisterCommand("remote_turn_on", RemoteTurnOn{})
	RegisterCommand("remote_turn_off", RemoteTurnOff{})
	RegisterCommand("remote_set_activity", RemoteSetActivity{})
	RegisterCommand("remote_send_command", RemoteSendCommand{})
	RegisterCommand("phone_register_push_token", PhoneRegisterPushToken{})
	RegisterCommand("phone_send_notification", PhoneSendNotification{})
	RegisterCommand("phone_send_data_message", PhoneSendDataMessage{})
}
