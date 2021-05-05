package ad

import (
	"errors"
	"time"

	"github.com/kgoins/ldapentity/entity"
)

type ADEntity struct {
	DN          string
	CN          string
	Name        string
	DisplayName string
	Description string

	WhenCreated time.Time
	WhenChanged time.Time
}

func NewADEntity(entity entity.Entity) (ADEntity, error) {
	dn, found := entity.GetDN()
	if !found {
		return ADEntity{}, errors.New("Unable to parse DN")
	}

	cn, _ := entity.GetSingleValuedAttribute(ATTR_CN)
	name, _ := entity.GetSingleValuedAttribute(ATTR_name)
	displayName, _ := entity.GetSingleValuedAttribute(ATTR_displayName)
	desc, _ := entity.GetSingleValuedAttribute(ATTR_description)

	adCreatedTime, _ := entity.GetSingleValuedAttribute("whenCreated")
	created, err := TimeFromADGeneralizedTime(adCreatedTime)
	if err != nil {
		return ADEntity{}, err
	}

	adChangedTime, _ := entity.GetSingleValuedAttribute("whenChanged")
	changed, err := TimeFromADGeneralizedTime(adChangedTime)
	if err != nil {
		return ADEntity{}, err
	}

	return ADEntity{
		DN:          dn,
		CN:          cn,
		Name:        name,
		DisplayName: displayName,
		Description: desc,
		WhenCreated: created,
		WhenChanged: changed,
	}, nil
}
