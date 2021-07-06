package ad_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/audibleblink/msldapuac"
	hashset "github.com/kgoins/hashset/pkg"
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
		want := hashset.NewStrHashset(
			ad.GetUACFlagName(msldapuac.Script),
			ad.GetUACFlagName(msldapuac.NormalAccount),
		)

		uac, err := ad.NewUAC("513")
		assert.NoError(t, err)

		flags, err := uac.GetFlagNames()
		assert.NoError(t, err)

		assert.True(t, flags.Equals(want))
	})

	t.Run("returns an error when invalid input is passed", func(t *testing.T) {
		_, err := ad.NewUAC("I AM NOT A NUMBER!")
		assert.NotNil(t, err)
	})
}
