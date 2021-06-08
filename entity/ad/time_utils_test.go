package ad_test

import (
	"testing"
	"time"

	"github.com/kgoins/ldapentity/entity/ad"
)

func TestTimeFromADGeneralizedTime(t *testing.T) {
	origTime, _ := time.Parse(time.UnixDate, "Wed Nov 28 11:43:42 CST 2018")
	adTime := "20181128174342.0Z"

	convTime, err := ad.TimeFromADGeneralizedTime(adTime)
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

	convTime := ad.TimeFromADTimestamp(adTime)

	origTimeStr := origTime.Format(time.ANSIC)
	convTimeStr := convTime.Format(time.ANSIC)

	if convTimeStr != origTimeStr {
		t.Errorf("Failed to convert AD timestamp")
	}
}

func TestADIntervalToDays(t *testing.T) {
	interval := "-36288000000000"
	days := 42

	parsedDays := ad.ADIntervalToDays(interval)

	if parsedDays != days {
		t.Errorf("Failed to convert from interval to days\n")
	}
}
