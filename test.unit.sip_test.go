package domain

import (
	"encoding/json"
	"testing"
)

func TestSIP_EntityTypesHydrate(t *testing.T) {
	cases := map[string]struct {
		state string
		check func(any) bool
	}{
		"pbx":          {`{"connected":true,"active_calls":2}`, func(v any) bool { s, ok := v.(PBX); return ok && s.Connected && s.ActiveCalls == 2 }},
		"sip_trunk":    {`{"registered":true}`, func(v any) bool { s, ok := v.(SIPTrunk); return ok && s.Registered }},
		"sip_endpoint": {`{"status":"online","call":"x"}`, func(v any) bool { s, ok := v.(SIPEndpoint); return ok && s.Status == "online" && s.Call == "x" }},
		"sip_call":     {`{"status":"ended","reason":"r"}`, func(v any) bool { s, ok := v.(SIPCall); return ok && s.Status == "ended" && s.Reason == "r" }},
		"agent":        {`{"status":"idle","result":"r","seq":2}`, func(v any) bool { s, ok := v.(Agent); return ok && s.Result == "r" && s.Seq == 2 }},
		"transcript":   {`{"status":"listening","text":"t","seq":3}`, func(v any) bool { s, ok := v.(Transcript); return ok && s.Text == "t" && s.Seq == 3 }},
		"speech": {`{"status":"speaking","text":"hi","seq":2}`, func(v any) bool {
			s, ok := v.(Speech)
			return ok && s.Text == "hi" && s.Seq == 2 && s.Status == "speaking"
		}},
		"sip_participant": {`{"status":"joined","muted":true}`, func(v any) bool { s, ok := v.(SIPParticipant); return ok && s.Status == "joined" && s.Muted }},
	}
	for typ, tc := range cases {
		var e Entity
		data := []byte(`{"id":"a","plugin":"p","deviceID":"d","type":"` + typ + `","state":` + tc.state + `}`)
		if err := json.Unmarshal(data, &e); err != nil {
			t.Fatalf("%s: %v", typ, err)
		}
		if !tc.check(e.State) {
			t.Errorf("%s: state %T %+v", typ, e.State, e.State)
		}
	}
}

// sb-script and sb-api only accept commands registered here, and scripts send
// these by name.
func TestSIP_CommandsAreRegisteredWithMatchingNames(t *testing.T) {
	cmds := []interface{ ActionName() string }{
		PBXReload{}, SIPDial{}, SIPHangup{}, SIPPlay{}, SIPMute{}, SIPUnmute{},
		SIPHold{}, SIPUnhold{}, SIPDTMF{}, SIPRemove{},
		SIPTranscribeStart{}, SIPTranscribeStop{}, SIPSay{}, AgentPrompt{}, AgentCancel{},
	}
	for _, c := range cmds {
		if _, ok := LookupCommand(c.ActionName()); !ok {
			t.Errorf("%s is not registered", c.ActionName())
		}
	}
}
