package entity

import (
	hashset "github.com/kgoins/hashset/pkg"
)

type Attribute struct {
	Name  string
	Value hashset.StrHashset
}

func NewEntityAttribute(name string, value string) Attribute {
	return Attribute{
		Name:  name,
		Value: hashset.NewStrHashset(value),
	}
}

func (attr *Attribute) SetValue(vals ...string) {
	attr.Value.Clear()
	attr.Value.Add(vals...)
}

func (attr Attribute) HasValue(val string) bool {
	return attr.Value.Contains(val)
}

func (attr Attribute) Equals(a2 Attribute) bool {
	if attr.Name != a2.Name {
		return false
	}

	return attr.Value.Equals(a2.Value)
}
