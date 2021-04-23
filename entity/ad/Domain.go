package ad

import (
	"time"

	ldap "gopkg.in/ldap.v2"
)

// Domain represents an AD domain object from the directory.
type Domain struct {
	DN              string
	Name            string
	SID             string
	WhenCreated     time.Time
	PwdPolicy       PasswordPolicy
	SupportedEtypes []KrbEtype
}

var domainAttributes = []string{
	ATTR_DN,
	ATTR_name,
	ATTR_sid,
	ATTR_created,
}

// NewDomainFromEntry builds a Domain from its ldap entry.
func NewDomainFromEntry(domainEntry *ldap.Entry) Domain {
	sidBytes := domainEntry.GetRawAttributeValue(ATTR_sid)
	sid := SidFromBytes(sidBytes)

	adCreatedTime := domainEntry.GetAttributeValue(ATTR_created)
	created, _ := TimeFromADGeneralizedTime(adCreatedTime)

	passwordPolicy := NewPasswordPolicyFromEntry(domainEntry)

	return Domain{
		domainEntry.GetAttributeValue(ATTR_DN),
		domainEntry.GetAttributeValue(ATTR_name),
		sid,
		created,
		passwordPolicy,
		[]KrbEtype{},
	}
}
