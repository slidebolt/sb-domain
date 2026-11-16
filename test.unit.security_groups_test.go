package domain

import (
	"encoding/json"
	"reflect"
	"testing"

	contract "github.com/slidebolt/sb-contract"
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

func TestDeviceSecurityGroupsUnionAndSort(t *testing.T) {
	device := Device{SecurityGroups: []string{SecurityGroupSystem}, Entities: []Entity{{Type: "light"}}}
	groups, err := DeviceSecurityGroups(device, Entity{Type: "media_player"}, Entity{Type: "light"})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{SecurityGroupHome, SecurityGroupMedia, SecurityGroupSystem}
	if !reflect.DeepEqual(groups, want) {
		t.Fatalf("groups = %v, want %v", groups, want)
	}
}

func TestDeviceSecurityGroupsDefaultsEmptyDeviceToSystem(t *testing.T) {
	groups, err := DeviceSecurityGroups(Device{})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(groups, []string{SecurityGroupSystem}) {
		t.Fatalf("groups = %v", groups)
	}
}

func TestNormalizeSecurityGroupsRejectsResourceWildcardAndUnknown(t *testing.T) {
	for _, group := range []string{"*", "unknown", ""} {
		if _, err := NormalizeSecurityGroups(group); err == nil {
			t.Errorf("NormalizeSecurityGroups(%q) succeeded", group)
		}
	}
}

func TestDeviceSecurityGroupsRoundTrip(t *testing.T) {
	original := Device{ID: "mixed", Plugin: "plugin-test", Name: "Mixed", SecurityGroups: []string{SecurityGroupHome, SecurityGroupSystem}, Entities: []Entity{}}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}
	var decoded Device
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(decoded.SecurityGroups, original.SecurityGroups) {
		t.Fatalf("securityGroups = %v", decoded.SecurityGroups)
	}
}

func TestValidatePluginSecurityGroups(t *testing.T) {
	definitions := contract.DescribePlugin("Media", "Media plugin.", "media_player")
	definitions.DeviceClasses = []contract.DeviceClassDefinition{{Class: "media_player", SecurityGroups: []string{SecurityGroupMedia}}}
	if err := ValidatePluginSecurityGroups(definitions); err != nil {
		t.Fatal(err)
	}
	definitions.DeviceClasses[0].SecurityGroups = []string{"unknown"}
	if err := ValidatePluginSecurityGroups(definitions); err == nil {
		t.Fatal("unknown device security group was accepted")
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
