package ad_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/audibleblink/msldapuac"
	"github.com/kgoins/ldapentity/entity/ad"
	"github.com/stretchr/testify/assert"
)

func TestUACPrint(t *testing.T) {
	t.Run("prints correct sorted output", func(t *testing.T) {
		want := "b5106b8639687bb965a84af85e69113a"
		buffer := &bytes.Buffer{}
		ad.UACPrint(buffer)

		got := fmt.Sprintf("%x", md5.Sum(buffer.Bytes()))
		assert.Equal(t, want, got)
	})
}

func TestUACParse(t *testing.T) {
	t.Run("parses and calculates the correct values", func(t *testing.T) {
		want := []string{
			ad.GetUACFlagName(msldapuac.Script),
			ad.GetUACFlagName(msldapuac.NormalAccount),
		}

		uac, err := ad.NewUAC("513")
		assert.NoError(t, err)

		flags, err := uac.GetFlagNames()
		assert.NoError(t, err)

		assert.Equal(t, want, flags)
	})

	t.Run("returns an error when invalid input is passed", func(t *testing.T) {
		got, err := ad.NewUAC("I AM NOT A NUMBER!")
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
}
