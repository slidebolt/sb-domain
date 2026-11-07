package domain

// PBX is the state of a phone system.
type PBX struct {
	Connected   bool   `json:"connected"`
	Version     string `json:"version,omitempty"`
	Uptime      int64  `json:"uptime,omitempty"`
	ActiveCalls int    `json:"active_calls"`
	Error       string `json:"error,omitempty"`
}

// SIPTrunk is a line to the outside world.
type SIPTrunk struct {
	Registered bool   `json:"registered"`
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	Latency    int    `json:"latency,omitempty"`
}

// SIPEndpoint is a phone, intercom or other extension.
type SIPEndpoint struct {
	Registered bool   `json:"registered"`
	InCall     bool   `json:"in_call"`
	IP         string `json:"ip,omitempty"`
	Agent      string `json:"agent,omitempty"`
	Status     string `json:"status,omitempty"` // online, offline
	Role       string `json:"role,omitempty"`   // e.g. intercom, service (a line answered by SlideBolt)
	Call       string `json:"call,omitempty"`
}

// SIPCall is a live conversation between participants.
type SIPCall struct {
	State     string `json:"state"`
	Caller    string `json:"caller"`
	Callee    string `json:"callee"`
	Duration  int    `json:"duration"`
	Direction string `json:"direction,omitempty"`
	Status    string `json:"status,omitempty"` // active, ended
	Reason    string `json:"reason,omitempty"`
}

// SIPParticipant is one party in a call.
type SIPParticipant struct {
	Status   string `json:"status"` // ringing, joined, no_answer, busy, failed, left
	Target   string `json:"target"`
	Muted    bool   `json:"muted"`
	Held     bool   `json:"held"`
	LastDTMF string `json:"last_dtmf,omitempty"`
	Reason   string `json:"reason,omitempty"`
}

// PBXReload reloads the phone system's configuration.
type PBXReload struct{}

func (PBXReload) ActionName() string { return "pbx_reload" }

// SIPDial adds the addressed endpoint, or Target, to the named call. The call is
// created if needed, and every party in it can hear each other.
type SIPDial struct {
	Call        string `json:"call"`
	Target      string `json:"target,omitempty"`
	RingSeconds int    `json:"ring_seconds,omitempty"`
}

func (SIPDial) ActionName() string { return "sip_call" }

// SIPHangup ends a call, or removes an endpoint from its current call.
type SIPHangup struct{}

func (SIPHangup) ActionName() string { return "sip_hangup" }

// SIPPlay plays a sound into a call.
type SIPPlay struct {
	Sound      string `json:"sound"`
	ThenHangup bool   `json:"then_hangup,omitempty"`
}

func (SIPPlay) ActionName() string { return "sip_play" }

type SIPMute struct{}

func (SIPMute) ActionName() string { return "sip_mute" }

type SIPUnmute struct{}

func (SIPUnmute) ActionName() string { return "sip_unmute" }

type SIPHold struct{}

func (SIPHold) ActionName() string { return "sip_hold" }

type SIPUnhold struct{}

func (SIPUnhold) ActionName() string { return "sip_unhold" }

// SIPDTMF sends keypad tones to a participant.
type SIPDTMF struct {
	Digits string `json:"digits"`
}

func (SIPDTMF) ActionName() string { return "sip_dtmf" }

// SIPRemove removes a participant from a call.
type SIPRemove struct{}

func (SIPRemove) ActionName() string { return "sip_remove" }

// SIPMessages is the inbox of a number answered by SlideBolt: the latest text
// message received. Seq rises with every message.
type SIPMessages struct {
	Seq  int    `json:"seq"`
	From string `json:"from,omitempty"` // the sending extension
	Text string `json:"text,omitempty"`
	At   string `json:"at,omitempty"` // RFC 3339
}

// SIPMessageSend sends a text message from the addressed number to an extension.
type SIPMessageSend struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

func (SIPMessageSend) ActionName() string { return "sip_message_send" }
