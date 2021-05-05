package ad

import (
	"github.com/kgoins/ldapentity/entity"
)

type Computer struct {
	ADEntity
	SecurableEntity

	Hostname        string
	OperatingSystem string
	OSVersion       string
	Services        []ServicePrincipal
}

func (c Computer) HasService(serviceName string) bool {
	for _, service := range c.Services {
		if service.Service == serviceName {
			return true
		}
	}

	return false
}

func NewComputerFromEntry(compEntity entity.Entity) (comp Computer, err error) {
	adEntity, err := NewADEntity(compEntity)
	if err != nil {
		return
	}

	se, err := NewSecurableEntity(compEntity)
	if err != nil {
		return
	}

	spns, _ := compEntity.GetAttribute(ATTR_spn)
	services := NewServicePrincipals(spns.Value.Values(), se.SamAccountName)

	hostname, _ := compEntity.GetSingleValuedAttribute(ATTR_dnsHostname)
	os, _ := compEntity.GetSingleValuedAttribute(ATTR_os)
	version, _ := compEntity.GetSingleValuedAttribute(ATTR_osVersion)

	computer := Computer{
		ADEntity:        adEntity,
		SecurableEntity: se,

		Hostname:        hostname,
		OperatingSystem: os,
		OSVersion:       version,
		Services:        services,
	}

	return computer, nil
}

func NewComputersFromEntries(compEntries []entity.Entity) []Computer {
	computers := make([]Computer, 0, len(compEntries))
	for _, compEntry := range compEntries {
		computer, err := NewComputerFromEntry(compEntry)
		if err != nil {
			continue
		}
		computers = append(computers, computer)
	}

	return computers
}
