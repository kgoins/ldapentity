package sadAD_test

import (
	"strings"
	"testing"
)

func TestGetDomain(t *testing.T) {
	domain, err := Session.GetDomain(Session.RootDSE.CurrentDomain)
	if err != nil {
		t.Errorf("Error getting current domain: %s", err.Error())
	}

	if domain.DN != Session.RootDSE.CurrentDomain {
		t.Errorf("Domain result not equal to current domain")
	}

	if domain.PwdPolicy.Empty() {
		t.Errorf("Failed to get domain password policy")
	}
}

func TestGetDomainsInForest(t *testing.T) {
	domains, err := Session.GetDomainsInForest()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if len(domains) == 0 {
		t.Errorf("Failed to return any domains")
		return
	}

	found := false
	for _, domain := range domains {
		if strings.EqualFold(domain.DN, Session.BaseDN) {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Failed to return current domain")
	}
}

func TestGetDomainStructure(t *testing.T) {
	containers, err := Session.GetDomainStructure(Session.RootDSE.CurrentDomain)
	if err != nil {
		t.Errorf("Error getting domain structure: %s\n", err.Error())
	}

	if len(containers) == 0 {
		t.Errorf("Failed to retrieve domain structure")
	}
}

func TestGetDomainServices(t *testing.T) {
	services, err := Session.GetDomainServices(Session.RootDSE.CurrentDomain)
	if err != nil {
		t.Errorf("Error finding services: %s\n", err.Error())
	}

	if len(services) == 0 {
		t.Errorf("Failed to retrieve any domain services")
	}

	foundKadmin := false
	for _, svc := range services {
		if svc.Service == "kadmin" {
			foundKadmin = true
			break
		}
	}

	if !foundKadmin {
		t.Errorf("Didn't get all services: kadmin not found")
	}
}
