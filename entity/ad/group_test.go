package sadAD_test

import (
	"strings"
	"testing"
)

func TestGetGroup(t *testing.T) {
	samaccountname := "Domain Admins"

	groups, err := Session.GetGroup(samaccountname, true)
	if err != nil {
		t.Errorf("Error getting group: %s\n", err.Error())
	}

	if groups == nil || len(groups) < 1 {
		t.Fatalf("Failed to retrieve any groups")
	}

	if !strings.EqualFold(groups[0].SamAccountName, samaccountname) {
		t.Errorf("Failed to retrieve correct group")
	}
}

func TestGetAllGroups(t *testing.T) {
	groups, err := Session.GetAllGroups()
	if err != nil {
		t.Errorf("Error getting all groups: %s", err.Error())
		return
	}

	if len(groups) < 4 {
		t.Errorf("Failed to get builtin groups")
	}
}
