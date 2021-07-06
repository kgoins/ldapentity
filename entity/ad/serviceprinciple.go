package ad

import (
	"fmt"
	"strings"
)

type ServicePrincipal struct {
	Service   string
	Target    string
	Principal string
}

var ignoredServices = []string{
	"Host",
	"RestrictedKrbHost",
}

func isIgnoredService(svc string) bool {
	for _, ignoredSvc := range ignoredServices {
		if strings.EqualFold(svc, ignoredSvc) {
			return true
		}
	}

	return false
}

func NewServicePrincipals(spns []string, principalName string) []ServicePrincipal {
	services := make([]ServicePrincipal, 0, len(spns))
	var spnParts []string

	for _, spn := range spns {
		spnParts = strings.SplitN(spn, "/", 2)

		if isIgnoredService(spnParts[0]) {
			continue
		}

		service := ServicePrincipal{
			Service:   spnParts[0],
			Target:    spnParts[1],
			Principal: principalName,
		}

		services = append(services, service)
	}

	return services
}

func (svc ServicePrincipal) String() string {
	return fmt.Sprintf("%s/%s", svc.Service, svc.Target)
}
