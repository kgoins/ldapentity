package entity_test

import (
	"testing"

	"github.com/kgoins/ldapentity/entity"
	"github.com/stretchr/testify/require"
)

func TestEntity_Equals_ShouldEqual(t *testing.T) {
	a1 := entity.NewEntityAttribute("a1", "v1")
	a2 := entity.NewEntityAttribute("a2", "v2")

	e1 := entity.NewEntity()
	e1.AddAttribute(a1)
	e1.AddAttribute(a2)

	e2 := entity.NewEntity()
	e2.AddAttribute(a1)
	e2.AddAttribute(a2)

	r := require.New(t)
	r.True(e1.Equals(e2))
}

func TestEntity_Equals_ShouldNotEqual(t *testing.T) {
	a1 := entity.NewEntityAttribute("a1", "v1")
	a2 := entity.NewEntityAttribute("a2", "v2")

	e1 := entity.NewEntity()
	e1.AddAttribute(a1)
	e1.AddAttribute(a2)

	e2 := entity.NewEntity()
	e2.AddAttribute(a1)

	r := require.New(t)
	r.False(e1.Equals(e2))
}
