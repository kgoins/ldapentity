package ad

import (
	"github.com/kgoins/ldapentity/entity"
)

// Group represents a generic AD group object
type Group struct {
	ADEntity
	SecurableEntity

	GroupType int32

	Members  []ADEntity
	MemberOf []Group
}

func newGroupStubsFromDNs(dnList []string) []Group {
	groups := make([]Group, 0, len(dnList))
	for _, dn := range dnList {
		groups = append(groups, Group{ADEntity: ADEntity{DN: dn}})
	}

	return groups
}

// NewGroupFromEntity creates a Group object from its ldap entry
func NewGroupFromEntity(groupEntity entity.Entity) (grp Group, err error) {
	grp.ADEntity, err = NewADEntity(groupEntity)
	if err != nil {
		return
	}

	grp.SecurableEntity, err = NewSecurableEntity(groupEntity)
	if err != nil {
		return
	}

	groupDNList, _ := groupEntity.GetAttribute(ATTR_memberOf)
	grp.MemberOf = newGroupStubsFromDNs(groupDNList.Value.Values())

	memberDNList, _ := groupEntity.GetAttribute(ATTR_member)
	grp.Members = newADEntitiesFromDNs(memberDNList.Value.Values())

	tempInt, _, err := groupEntity.GetAsInt(ATTR_groupType)
	if err != nil {
		return
	}
	grp.GroupType = int32(tempInt)
	return
}
