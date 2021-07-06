package ad_test

import (
	"testing"
	"time"

	"github.com/kgoins/ldapentity/entity/ad"
	"github.com/stretchr/testify/require"
)

func TestTimeFromADGeneralizedTime(t *testing.T) {
	r := require.New(t)

	origTime, err := time.Parse(time.UnixDate, "Wed Nov 28 17:43:42 +0000 2018")
	r.NoError(err)

	adTime := "20181128174342.0Z"
	convTime, err := ad.TimeFromADGeneralizedTime(adTime)
	r.NoError(err)

	origTimeStr := origTime.Format(time.UnixDate)
	convTimeStr := convTime.Format(time.UnixDate)

	r.Equal(origTimeStr, convTimeStr)
}

func TestTimeFromADTimestamp(t *testing.T) {
	r := require.New(t)

	origTime, err := time.Parse(time.UnixDate, "Fri Nov 30 17:54:26 CST 2018")
	r.NoError(err)

	adTime := "131880920666196310"
	convTime := ad.TimeFromADTimestamp(adTime)

	origTimeStr := origTime.Format(time.ANSIC)
	convTimeStr := convTime.Format(time.ANSIC)

	r.Equal(origTimeStr, convTimeStr)
}

func TestADIntervalToDays(t *testing.T) {
	interval := "-36288000000000"
	days := 42

	parsedDays := ad.ADIntervalToDays(interval)

	if parsedDays != days {
		t.Errorf("Failed to convert from interval to days\n")
	}
}
