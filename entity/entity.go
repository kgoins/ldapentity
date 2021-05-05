package entity

import (
	"strings"
)

type Entity struct {
	attributes map[string]EntityAttribute
}

func NewEntity() Entity {
	return Entity{
		attributes: make(map[string]EntityAttribute),
	}
}

func (e *Entity) AddAttribute(attr EntityAttribute) {
	attrName := strings.ToLower(attr.Name)

	existing, found := e.GetAttribute(attrName)
	if !found {
		e.attributes[attrName] = attr
		return
	}

	existing.Value.Add(attr.Value.Values()...)
}

func (e Entity) IsEmpty() bool {
	return len(e.attributes) == 0
}

func (e Entity) GetDN() (string, bool) {
	dn, found := e.GetSingleValuedAttribute("dn")
	if !found {
		dn, found = e.GetSingleValuedAttribute("distinguishedName")
	}

	return dn, found
}

func (e Entity) GetAllAttributeNames() []string {
	names := make([]string, e.Size())

	i := 0
	for key := range e.attributes {
		names[i] = key
		i++
	}

	return names
}

func (e Entity) GetAllAttributes() []EntityAttribute {
	attrs := make([]EntityAttribute, e.Size())

	i := 0
	for _, val := range e.attributes {
		attrs[i] = val
	}

	return attrs
}

func (e *Entity) SetAttribute(attr EntityAttribute) {
	attrName := strings.ToLower(attr.Name)
	e.attributes[attrName] = attr
}

func (e Entity) GetAttribute(name string) (EntityAttribute, bool) {
	val, found := e.attributes[strings.ToLower(name)]
	return val, found
}

func (e Entity) GetSingleValuedAttribute(name string) (string, bool) {
	val, found := e.attributes[strings.ToLower(name)]
	if !found {
		return "", false
	}

	if len(val.Value.Values()) != 1 {
		return "", false
	}

	return val.Value.Values()[0], true
}

func (e Entity) Size() int {
	return len(e.attributes)
}
