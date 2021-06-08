package entity

import (
	"errors"
	"strings"

	hashset "github.com/kgoins/hashset/pkg"
)

type EntityAttribute struct {
	Name  string
	Value hashset.StrHashset
}

func NewEntityAttribute(name string, value string) EntityAttribute {
	return EntityAttribute{
		Name:  name,
		Value: hashset.NewStrHashset(value),
	}
}

func BuildEntityAttribute(name string, initValue string) EntityAttribute {
	return EntityAttribute{
		Name:  strings.TrimRight(name, ":"),
		Value: hashset.NewStrHashset(initValue),
	}
}

func BuildAttributeFromLine(attrLine string) (EntityAttribute, error) {
	lineParts := strings.Split(attrLine, ": ")
	if len(lineParts) != 2 {
		return EntityAttribute{}, errors.New("malformed attribute line")
	}

	return BuildEntityAttribute(lineParts[0], lineParts[1]), nil
}

func (attr *EntityAttribute) SetValue(vals ...string) {
	attr.Value.Clear()
	attr.Value.Add(vals...)
}

func (attr EntityAttribute) HasValue(val string) bool {
	return attr.Value.Contains(val)
}
