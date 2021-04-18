package sadAD

import (
	"fmt"
	"strings"
)

// GetDomainService returns a ServicePrincipal for the queried SPNs.
func (ls LdapSession) GetDomainService(serviceName, domainDN string) (services []ServicePrincipal, err error) {
	query := fmt.Sprintf("(servicePrincipalName=%s*)", serviceName)
	attrs := []string{ATTR_spn, ATTR_sAMAccountName}

	params, err := BuildSearchParams(query, attrs, domainDN, ScopeSubtree, AD_pageMax)
	if err != nil {
		return
	}

	resp, err := ls.Search(params)
	if err != nil {
		return
	}

	for _, entity := range resp {
		var filtered []string
		for _, spn := range entity.GetAttributeValues(ATTR_spn) {
			if strings.Contains(strings.ToLower(spn), strings.ToLower(serviceName)) {
				filtered = append(filtered, spn)
			}
		}

		newSvcs := NewServicePrincipals(
			filtered,
			entity.GetAttributeValue(ATTR_sAMAccountName),
		)
		services = append(services, newSvcs...)
	}
	return
}

// GetDomainServices returns ServicePrincipals for all registered SPNs in the domain.
func (ls LdapSession) GetDomainServices(domainDN string) ([]ServicePrincipal, error) {
	return ls.GetDomainService("", domainDN)
}
