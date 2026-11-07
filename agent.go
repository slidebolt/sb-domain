package domain

// Agent is an AI agent that takes prompts and reports results.
type Agent struct {
	Status  string `json:"status"` // idle, thinking, error
	Session string `json:"session,omitempty"`
	Prompt  string `json:"prompt,omitempty"`
	Result  string `json:"result,omitempty"`
	Error   string `json:"error,omitempty"`
	Seq     int    `json:"seq"`
}

// AgentPrompt sends text to an agent. Calls with the same session share a conversation.
type AgentPrompt struct {
	Text    string `json:"text"`
	Session string `json:"session,omitempty"`
}

func (AgentPrompt) ActionName() string { return "agent_prompt" }

// AgentCancel stops what an agent is doing.
type AgentCancel struct{}

func (AgentCancel) ActionName() string { return "agent_cancel" }

// Transcript is live speech-to-text. Seq rises with every utterance, so a repeated
// phrase is still a state change.
type Transcript struct {
	Status string `json:"status"` // idle, starting, listening, error
	Seq    int    `json:"seq"`
	Text   string `json:"text,omitempty"`
	Error  string `json:"error,omitempty"`
}

// SIPTranscribeStart starts speech-to-text on a call.
type SIPTranscribeStart struct{}

func (SIPTranscribeStart) ActionName() string { return "sip_transcribe_start" }

// SIPTranscribeStop stops speech-to-text. The call carries on.
type SIPTranscribeStop struct{}

func (SIPTranscribeStop) ActionName() string { return "sip_transcribe_stop" }

// Speech is text being spoken into a call. Seq rises with every utterance spoken.
type Speech struct {
	Status string `json:"status"` // idle, speaking, error
	Seq    int    `json:"seq"`
	Text   string `json:"text,omitempty"`
	Error  string `json:"error,omitempty"`
}

// SIPSay speaks text into a call, through its speech-to-text session. Saying
// something new interrupts anything still being spoken.
type SIPSay struct {
	Text string `json:"text"`
}

func (SIPSay) ActionName() string { return "sip_say" }
