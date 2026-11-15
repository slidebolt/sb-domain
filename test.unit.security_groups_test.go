package domain

import (
	"encoding/json"
	"testing"
)

func TestSecurityGroupDefaults(t *testing.T) {
	tests := map[string]string{
		"lock":          SecurityGroupSecurity,
		"camera":        SecurityGroupPrivate,
		"light":         SecurityGroupHome,
		"media_library": SecurityGroupMedia,
		"agent":         SecurityGroupSystem,
		"plugin_custom": SecurityGroupSystem,
	}
	for entityType, want := range tests {
		if got := DefaultSecurityGroup(entityType); got != want {
			t.Errorf("DefaultSecurityGroup(%q) = %q, want %q", entityType, got, want)
		}
	}
}

func TestSecurityGroupDefinitionsAreValid(t *testing.T) {
	for _, definition := range SecurityGroups() {
		if definition.ID == "" || definition.Name == "" || definition.Description == "" {
			t.Fatalf("incomplete definition: %+v", definition)
		}
		if !ValidSecurityGroup(definition.ID) {
			t.Fatalf("definition %q is not valid", definition.ID)
		}
	}
	if !ValidSecurityGroup("*") || ValidSecurityGroup("unknown") {
		t.Fatal("wildcard or unknown group validation is incorrect")
	}
}

func TestEntitySecurityGroupRoundTrip(t *testing.T) {
	original := Entity{ID: "lamp", Plugin: "plugin-test", DeviceID: "device", Type: "light", Name: "Lamp", SecurityGroup: SecurityGroupHome, State: Light{}}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	var decoded Entity
	if err := decoded.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}
	if decoded.SecurityGroup != SecurityGroupHome {
		t.Fatalf("securityGroup = %q", decoded.SecurityGroup)
	}
}
