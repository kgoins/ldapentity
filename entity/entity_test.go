package entity_test

import (
	"testing"

	"github.com/kgoins/ldapentity/entity"
	"github.com/stretchr/testify/require"
)

func TestEntity_Equals_ShouldEqual(t *testing.T) {
	a1 := entity.NewEntityAttribute("a1", "v1")
	a2 := entity.NewEntityAttribute("a2", "v2")

	e1 := entity.NewEntity("cn=testentity")
	e1.AddAttribute(a1)
	e1.AddAttribute(a2)

	e2 := entity.NewEntity("cn=testentity")
	e2.AddAttribute(a1)
	e2.AddAttribute(a2)

	r := require.New(t)
	r.True(e1.Equals(e2))
}

func TestEntity_Equals_ShouldNotEqual(t *testing.T) {
	a1 := entity.NewEntityAttribute("a1", "v1")
	a2 := entity.NewEntityAttribute("a2", "v2")

	e1 := entity.NewEntity("cn=testentity")
	e1.AddAttribute(a1)
	e1.AddAttribute(a2)

	e2 := entity.NewEntity("cn=testentity")
	e2.AddAttribute(a1)

	r := require.New(t)
	r.False(e1.Equals(e2))
}

func TestEntity_GetID(t *testing.T) {
	r := require.New(t)

	e0 := entity.NewEntity("")
	_, err := e0.GetID()
	r.Error(err)

	e1 := entity.NewEntity("cn=testentity")
	knownHash := "57ab144f546fa1e3"

	id, err := e1.GetID()
	r.NoError(err)

	r.Equal(knownHash, id)
}
