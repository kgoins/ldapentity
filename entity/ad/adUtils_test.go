package ad_test

import (
	"testing"

	"github.com/kgoins/ldapentity/entity/ad"
)

func TestSidConversion(t *testing.T) {
	b64TestSid := "AQQAAAAAAAUVAAAA/6wpyB/p2/kT/vkD"
	testSid := "S-1-5-21-3358174463-4191938847-66715155"

	sid, err := ad.SidFromBase64(b64TestSid)
	if err != nil {
		t.Errorf("Error decoding SID: %s", err.Error())
	}

	if sid != testSid {
		t.Errorf("Failed to convert SID")
	}
}

func TestLdapRegex(t *testing.T) {
	goodLdapQuery := "(&(objectClass=user)(mail=*))"
	invalidLdapQuery := "((objectClass=user)(mail=*))"
	//invalidLdapQuery2 := "(&(objectClass=user(mail=*))"

	if !ad.IsValidLdapQuery(goodLdapQuery) {
		t.Errorf("Failed to identify correct LDAP query")
	}

	if ad.IsValidLdapQuery(invalidLdapQuery) {
		t.Errorf("Failed to identify invalid query with logical error")
	}
}

func TestIsGuid(t *testing.T) {
	validGuid1 := "{3A8EA7C0-1E7D-4626-BC23-42DA1A951FDB}"
	validGuid2 := "434bb40d-dbc9-4fe7-81d4-d57229f7b080"

	invalidGuid1 := "51b16c80-901c-4270-93a7-120a8c%b42ab"
	invalidGuid2 := "80112dcf-bec8-428d-8f0-b4fd0ca506d5"

	if !ad.IsGuid(validGuid1) || !ad.IsGuid(validGuid2) {
		t.Errorf("Failed to identify valid guids\n")
	}

	if ad.IsGuid(invalidGuid1) || ad.IsGuid(invalidGuid2) {
		t.Errorf("Failed to identify invalid guids\n")
	}
}
