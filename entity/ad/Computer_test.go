package sadAD_test

import "testing"

func TestGetDomainControllers(t *testing.T) {
	domainControllers := Session.GetDomainControllers()

	if len(domainControllers) == 0 {
		t.Errorf("Failed to retrieve domain controllers")
	}

	for _, dc := range domainControllers {
		if !dc.HasService("ldap") {
			t.Errorf("Domain controller entry corrupted - does not have ldap service")
		}
	}
}

func TestGetComputer(t *testing.T) {
	computer, err := Session.GetComputer(Session.RootDSE.DCHostname)
	if err != nil {
		t.Errorf("Get computer test failed: %s\n", err.Error())
	}

	if computer.Hostname != Session.RootDSE.DCHostname {
		t.Errorf("Computer hostname does not match\n")
	}
}
