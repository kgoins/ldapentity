package ad

import "github.com/kgoins/ldapentity/entity"

type SecurableEntity struct {
	SID            SID
	SamAccountName string
}

func NewSecurableEntity(entity entity.Entity) (se SecurableEntity, err error) {
	sid, err := NewSIDFromEntity(entity)
	if err != nil {
		return
	}

	samaccountname, _ := entity.GetSingleValuedAttribute(ATTR_sAMAccountName)

	return SecurableEntity{
		SID:            sid,
		SamAccountName: samaccountname,
	}, nil
}
