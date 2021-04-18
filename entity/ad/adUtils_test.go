package sadAD_test

import (
	"testing"

	sadAD "github.com/kgoins/sadAD/pkg"
)

func TestSidConversion(t *testing.T) {
	b64TestSid := "AQQAAAAAAAUVAAAA/6wpyB/p2/kT/vkD"
	testSid := "S-1-5-21-3358174463-4191938847-66715155"

	sid, err := sadAD.SidFromBase64(b64TestSid)
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

	if !sadAD.IsValidLdapQuery(goodLdapQuery) {
		t.Errorf("Failed to identify correct LDAP query")
	}

	if sadAD.IsValidLdapQuery(invalidLdapQuery) {
		t.Errorf("Failed to identify invalid query with logical error")
	}
}

/*
func TestTimeFromADGeneralizedTime(t *testing.T) {
	origTime, _ := time.Parse(time.UnixDate, "Wed Nov 28 11:43:42 CST 2018")
	adTime := "20181128174342.0Z"

	convTime, err := sadAD.TimeFromADGeneralizedTime(adTime)
	if err != nil {
		t.Errorf("Failed to convert AD generalized time")
	}

	origTimeStr := origTime.Format(time.ANSIC)
	convTimeStr := convTime.Format(time.ANSIC)

	if convTimeStr != origTimeStr {
		t.Errorf("Failed to convert AD generalized time")
	}
}

func TestTimeFromADTimestamp(t *testing.T) {
	origTime, _ := time.Parse(time.UnixDate, "Wed Nov 28 11:43:42 CST 2018")
	adTime := "131880920666196310"

	convTime := sadAD.TimeFromADTimestamp(adTime)

	origTimeStr := origTime.Format(time.ANSIC)
	convTimeStr := convTime.Format(time.ANSIC)

	if convTimeStr != origTimeStr {
		t.Errorf("Failed to convert AD timestamp")
	}
}
*/

func TestADIntervalToDays(t *testing.T) {
	interval := "-36288000000000"
	days := 42

	parsedDays := sadAD.ADIntervalToDays(interval)

	if parsedDays != days {
		t.Errorf("Failed to convert from interval to days\n")
	}
}

func TestGetFlagsFromUAC(t *testing.T) {
	uac := int64(262672)
	flags := sadAD.GetFlagsFromUAC(uac)

	if !flags.Enabled || !flags.LockedOut || !flags.SmartcardRequired {
		t.Errorf("Failed to decode UAC attributes correctly")
	}
}

func TestIsGuid(t *testing.T) {
	validGuid1 := "{3A8EA7C0-1E7D-4626-BC23-42DA1A951FDB}"
	validGuid2 := "434bb40d-dbc9-4fe7-81d4-d57229f7b080"

	invalidGuid1 := "51b16c80-901c-4270-93a7-120a8c%b42ab"
	invalidGuid2 := "80112dcf-bec8-428d-8f0-b4fd0ca506d5"

	if !sadAD.IsGuid(validGuid1) || !sadAD.IsGuid(validGuid2) {
		t.Errorf("Failed to identify valid guids\n")
	}

	if sadAD.IsGuid(invalidGuid1) || sadAD.IsGuid(invalidGuid2) {
		t.Errorf("Failed to identify invalid guids\n")
	}
}
