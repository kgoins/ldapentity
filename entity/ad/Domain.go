package sadAD

import (
	"errors"
	"fmt"
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

// NewDomainFromEntry builds a Domain from it's ldap entry.
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

// GetDomain queries the directory for the domain specified by domainDN and returns a Domain object.
func (ls LdapSession) GetDomain(domainDN string) (Domain, error) {
	attrs := append(domainAttributes, passwordPolicyAttributes...)

	params, err := BuildSearchParams(QUERY_All, attrs, domainDN, ScopeBase, AD_pageMax)
	if err != nil {
		return Domain{}, err
	}

	searchResp, err := ls.Search(params)
	if err != nil {
		return Domain{}, err
	}

	if len(searchResp) != 1 {
		msg := fmt.Sprintf("Invalid number of responses returned: %d", len(searchResp))
		return Domain{}, errors.New(msg)
	}

	domain := NewDomainFromEntry(searchResp[0])
	domain.SupportedEtypes, _ = ls.GetSupportedEtypes()

	return domain, nil
}

// GetDomainsInForest finds all domains in the current forest.
func (ls LdapSession) GetDomainsInForest() ([]Domain, error) {
	configDN := ls.RootDSE.ConfigNamingContext
	attrs := []string{"nCName"}

	params, err := BuildSearchParams("(NETBIOSName=*)", attrs, configDN, ScopeSubtree, AD_pageMax)
	if err != nil {
		return nil, err
	}

	domainEntries, err := ls.Search(params)
	if err != nil {
		return nil, err
	}

	domains := make([]Domain, 0, len(domainEntries))

	ncname := ""
	for _, entry := range domainEntries {
		ncname = entry.GetAttributeValue("nCName")

		domain := Domain{
			DN: ncname,
		}

		domains = append(domains, domain)
	}

	return domains, nil
}

// GetDomainStructure will return all containers / OUs in the current domain, excluding certain built-ins.
func (ls LdapSession) GetDomainStructure(domainDN string) ([]string, error) {
	query := "(|(objectClass=organizationalUnit)(objectClass=container))"
	params, err := BuildSearchParams(query, nil, domainDN, ScopeSubtree, AD_pageMax)
	if err != nil {
		return nil, err
	}

	resp, err := ls.Search(params)
	if err != nil {
		return nil, err
	}

	containers := make([]string, 0, len(resp))
	for _, entry := range resp {
		if !IsGuid(entry.DN) {
			containers = append(containers, entry.DN)
		}
	}

	return containers, nil
}
