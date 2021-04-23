package ad_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strings"
	"testing"

	"github.com/kgoins/ldapentity/entity"
	"github.com/kgoins/ldapentity/entity/ad"
	"github.com/stretchr/testify/assert"
)

func NewTestEntities() []entity.Entity {
	entityLines := strings.Split(entity.EntityStr, "\n")[1:]
	entity := entity.Buildentity.Entity(entityLines)
	return []entity.Entity{entity}
}

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

		want := []string{"NORMAL_ACCOUNT", "SCRIPT"}
		got, err := ad.UACParse("513")
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	})

	t.Run("returns an error when invalid input is passed", func(t *testing.T) {
		got, err := ad.UACParse("I AM NOT A NUMBER!")
		assert.Nil(t, got)
		assert.NotNil(t, err)
	})
}

func TestUACSearch(t *testing.T) {
	entities := NewTestEntities()

	t.Run("returns a correct entity when there is a match preset", func(t *testing.T) {
		got := ad.UACSearch(&entities, 512)
		want := entities
		assert.ElementsMatch(t, want, got)
	})

	t.Run("returns no entity when no match is preset", func(t *testing.T) {
		got := ad.UACSearch(&entities, 513)
		want := make([]entity.Entity, 0)
		assert.ElementsMatch(t, want, got)
	})
}
