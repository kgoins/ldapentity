package sadAD_test

import (
	"strings"
	"testing"
)

func TestGetAllZones(t *testing.T) {
	zones, err := Session.GetAllZones()
	if err != nil {
		t.Errorf("Error getting zones: %s\n", err.Error())
		return
	}

	if len(zones) == 0 {
		t.Errorf("Failed to retrieve zones")
		return
	}

	for _, zone := range zones {
		if strings.Contains(zone.DN, "RootDNSServers") {
			t.Errorf("Failed to exclude root dns zone")
		}
	}
}

func TestGetAllARecords(t *testing.T) {
	records, err := Session.GetAllARecords()
	if err != nil {
		t.Errorf("Error getting A records: %s\n", err.Error())
		return
	}

	if len(records) == 0 {
		t.Errorf("Failed to retrieve A records")
		return
	}

	found := false
	dcName := strings.ToLower(strings.Split(Session.RootDSE.DCHostname, ".")[0])

	for _, record := range records {
		if record.Hostname == dcName {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Records incomplete - DC not found")
	}
}
