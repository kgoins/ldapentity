package entity

import (
	"strconv"
	"strings"
)

type Entity struct {
	attributes map[string]Attribute
}

func NewEntity() Entity {
	return Entity{
		attributes: make(map[string]Attribute),
	}
}

func (e *Entity) AddAttribute(attr Attribute) {
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

func (e Entity) GetAllAttributes() []Attribute {
	attrs := make([]Attribute, e.Size())

	i := 0
	for _, val := range e.attributes {
		attrs[i] = val
	}

	return attrs
}

func (e *Entity) SetAttribute(attr Attribute) {
	attrName := strings.ToLower(attr.Name)
	e.attributes[attrName] = attr
}

func (e Entity) GetAttribute(name string) (Attribute, bool) {
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

func (e Entity) GetAsInt(name string) (i int, found bool, err error) {
	val, found := e.GetSingleValuedAttribute(name)
	if !found {
		return
	}

	i, err = strconv.Atoi(val)
	return
}

func (e Entity) Size() int {
	return len(e.attributes)
}

func (e Entity) Equals(e2 Entity) bool {
	if e.Size() != e2.Size() {
		return false
	}

	for _, attrName := range e.GetAllAttributeNames() {
		a1, f1 := e.GetAttribute(attrName)
		a2, f2 := e2.GetAttribute(attrName)

		if !f1 || !f2 {
			return false
		}

		if !a1.Equals(a2) {
			return false
		}
	}

	return true
}
