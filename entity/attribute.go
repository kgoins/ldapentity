package entity

import (
	"fmt"

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

func (attr Attribute) Stringify() []string {
	vals := make([]string, 0, attr.Value.Size())

	for _, value := range attr.Value.Values() {
		vals = append(vals, fmt.Sprintf("%s: %s", attr.Name, value))
	}

	return vals
}
