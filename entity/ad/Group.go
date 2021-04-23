package ad

import (
	"strconv"
	"time"

	ldap "gopkg.in/ldap.v2"
)

// Group represents a generic AD group object
type Group struct {
	DN             string
	CN             string
	SamAccountName string
	DisplayName    string
	Description    string

	SID string

	WhenCreated time.Time
	WhenChanged time.Time

	GroupType int32

	Members  []User
	MemberOf []Group
}

var groupAttributes = []string{
	"distinguishedName",
	"cn",
	"samaccountname",
	"displayName",
	"description",

	"objectSid",

	"whenCreated",
	"whenChanged",

	"groupType",

	"member",
	"memberOf",
}

func newGroupStubsFromDNs(dnList []string) []Group {
	groups := make([]Group, 0, len(dnList))
	for _, dn := range dnList {
		groups = append(groups, Group{DN: dn})
	}

	return groups
}

// NewGroupFromEntry creates a Group object from its ldap entry
func NewGroupFromEntry(groupEntry *ldap.Entry) Group {
	sidBytes := groupEntry.GetRawAttributeValue("objectSid")
	sid := SidFromBytes(sidBytes)

	adCreatedTime := groupEntry.GetAttributeValue("whenCreated")
	created, _ := TimeFromADGeneralizedTime(adCreatedTime)

	adChangedTime := groupEntry.GetAttributeValue("whenChanged")
	changed, _ := TimeFromADGeneralizedTime(adChangedTime)

	groupDNList := groupEntry.GetAttributeValues("memberOf")
	groups := newGroupStubsFromDNs(groupDNList)

	memberDNList := groupEntry.GetAttributeValues("member")
	members := newUserStubsFromDNs(memberDNList)

	groupTypeStr := groupEntry.GetAttributeValue("groupType")
	groupType, _ := strconv.Atoi(groupTypeStr)

	group := Group{
		DN:             groupEntry.GetAttributeValue("distinguishedName"),
		CN:             groupEntry.GetAttributeValue("cn"),
		SamAccountName: groupEntry.GetAttributeValue("sAMAccountName"),
		DisplayName:    groupEntry.GetAttributeValue("displayName"),
		Description:    groupEntry.GetAttributeValue("description"),

		SID: sid,

		WhenCreated: created,
		WhenChanged: changed,

		GroupType: int32(groupType),

		Members:  members,
		MemberOf: groups,
	}

	return group
}
