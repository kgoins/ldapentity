package sadAD

import (
	"errors"
	"fmt"
	"strings"
	"time"

	ldap "gopkg.in/ldap.v2"
)

type Computer struct {
	Hostname        string
	SamAccountName  string
	OperatingSystem string
	OSVersion       string
	WhenCreated     time.Time
	WhenChanged     time.Time
	SID             string
	Services        []ServicePrincipal
}

var compAttrs = []string{
	ATTR_dnsHostname,
	ATTR_sAMAccountName,
	ATTR_os,
	ATTR_osVersion,
	ATTR_created,
	ATTR_changed,
	ATTR_spn,
	ATTR_sid,
}

func (c Computer) Empty() bool {
	return c.Hostname == ""
}

func (c Computer) HasService(serviceName string) bool {
	for _, service := range c.Services {
		if service.Service == serviceName {
			return true
		}
	}

	return false
}

func NewComputerFromEntry(compEntry *ldap.Entry) Computer {
	samAccountName := compEntry.GetAttributeValue(ATTR_sAMAccountName)

	spns := compEntry.GetAttributeValues(ATTR_spn)
	services := NewServicePrincipals(spns, samAccountName)

	sidBytes := compEntry.GetRawAttributeValue(ATTR_sid)
	sid := SidFromBytes(sidBytes)

	adCreatedTime := compEntry.GetAttributeValue(ATTR_created)
	created, _ := TimeFromADGeneralizedTime(adCreatedTime)

	adChangedTime := compEntry.GetAttributeValue(ATTR_changed)
	changed, _ := TimeFromADGeneralizedTime(adChangedTime)

	computer := Computer{
		Hostname:        compEntry.GetAttributeValue(ATTR_dnsHostname),
		SamAccountName:  samAccountName,
		OperatingSystem: compEntry.GetAttributeValue(ATTR_os),
		OSVersion:       compEntry.GetAttributeValue(ATTR_osVersion),
		WhenCreated:     created,
		WhenChanged:     changed,
		SID:             sid,
		Services:        services,
	}

	return computer
}

func NewComputersFromEntries(compEntries []*ldap.Entry) []Computer {
	computers := make([]Computer, 0, len(compEntries))
	for _, compEntry := range compEntries {
		computer := NewComputerFromEntry(compEntry)
		computers = append(computers, computer)
	}

	return computers
}

func (ls LdapSession) GetComputer(hostname string) (Computer, error) {
	filter := fmt.Sprintf(QUERY_ComputerByHostname, hostname)

	compEntries, err := ls.BasicSearch(filter, compAttrs, AD_pageMax)
	if err != nil {
		return Computer{}, err
	}

	if len(compEntries) == 0 {
		return Computer{}, nil
	}

	computers := NewComputersFromEntries(compEntries)
	if len(computers) > 1 {
		errStrBuilder := strings.Builder{}
		errStrBuilder.WriteString("Multiple results found:\n")

		for _, comp := range computers {
			errStrBuilder.WriteString(comp.Hostname + "\n")
		}

		return Computer{}, errors.New(errStrBuilder.String())
	}

	return computers[0], nil
}

func (ls LdapSession) GetAllComputers() ([]Computer, error) {
	compEntries, err := ls.BasicSearch(QUERY_AllComputers, compAttrs, AD_pageMax)
	if err != nil {
		return []Computer{}, err
	}

	if len(compEntries) == 0 {
		return []Computer{}, nil
	}

	computers := NewComputersFromEntries(compEntries)
	return computers, nil
}

func (ls LdapSession) GetDomainControllers() []Computer {
	dcEntries, err := ls.BasicSearch(QUERY_DomainControllers, compAttrs, AD_pageMax)
	if err != nil {
		return []Computer{}
	}

	domainControllers := NewComputersFromEntries(dcEntries)
	return domainControllers
}
