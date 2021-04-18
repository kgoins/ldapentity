package sadAD

import (
	"errors"
	"fmt"
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

// GetGroup finds a group in the current domain by its samaccountname
func (ls LdapSession) GetGroup(samaccountname string, recurse bool) (groups []Group, err error) {

	query := fmt.Sprintf(QUERY_GetGroup, samaccountname)

	respEntries, err := ls.BasicSearch(query, groupAttributes, AD_pageMax)
	if err != nil {
		return
	}

	if len(respEntries) == 0 {
		err = errors.New("Group not found")
		return
	}

	for _, resp := range respEntries {
		group := NewGroupFromEntry(resp)
		if recurse {
			query = fmt.Sprintf(QUERY_GetGroupRe, group.DN)
			memberEntries, sErr := ls.BasicSearch(query, groupAttributes, AD_pageMax)
			if sErr != nil {
				err = sErr
				return
			}

			for _, memberEntry := range memberEntries {
				member := NewUserFromEntry(memberEntry)
				group.Members = append(group.Members, member)
			}

		}
		groups = append(groups, group)
	}

	return
}

// GetAllGroups returns all group objects in the current domain
func (ls LdapSession) GetAllGroups() ([]Group, error) {
	groupEntries, err := ls.BasicSearch(QUERY_GetAllGroups, groupAttributes, AD_pageMax)
	if err != nil {
		return []Group{}, err
	}

	groups := make([]Group, 0, len(groupEntries))
	for _, groupEntry := range groupEntries {
		groups = append(groups, NewGroupFromEntry(groupEntry))
	}

	return groups, nil
}
