package entity

import (
	"strconv"
	"errors"
	"fmt"

	"github.com/cespare/xxhash/v2"
)

type Entity struct {
	attributes map[string]Attribute
}

func NewEntity(dn string) Entity {
	e := Entity{
		attributes: make(map[string]Attribute),
	}
	e.AddAttribute(NewEntityAttribute("dn", dn))

	return e
}

func (e *Entity) AddAttribute(attr Attribute) {
	existing, found := e.GetAttribute(attr.Name)
	if !found {
		e.attributes[attr.Name] = attr
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

	if dn == "" {
		return dn, false
	}

	return dn, found
}

func (e Entity) GetID() (string, error) {
	dn, found := e.GetDN()
	if !found {
		return "", errors.New("unable to generate ID without DN")
	}

	dnHash := xxhash.Sum64String(dn)
	return fmt.Sprintf("%x", dnHash), nil
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
	e.attributes[attr.Name] = attr
}

func (e Entity) GetAttribute(name string) (Attribute, bool) {
	val, found := e.attributes[name]
	return val, found
}

func (e Entity) GetSingleValuedAttribute(name string) (string, bool) {
	val, found := e.attributes[name]
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
