package ad

import (
	"fmt"
	"strconv"
	"strings"

	ldap "gopkg.in/ldap.v2"
)

type RootDSE struct {
	CurrentDomain         string
	RootNamingContext     string
	ConfigNamingContext   string
	RealmName             string
	DCHostname            string
	ForestFunctionalLevel int
}

func (rootdse RootDSE) String() string {
	str := fmt.Sprintf("Current domain: %s\n", rootdse.CurrentDomain)
	str += fmt.Sprintf("Root naming context: %s\n", rootdse.RootNamingContext)
	str += fmt.Sprintf("Domain config container: %s\n", rootdse.ConfigNamingContext)
	str += fmt.Sprintf("Kerberos realm: %s\n", rootdse.RealmName)
	str += fmt.Sprintf("Current domain controller: %s\n", rootdse.DCHostname)
	str += fmt.Sprintf("Forest functional level: %d\n", rootdse.ForestFunctionalLevel)

	return str
}

func NewRootDSEFromEntity(entity *ldap.Entry) RootDSE {
	currentDomain := entity.GetAttributeValue("defaultNamingContext")
	rootNamingContext := entity.GetAttributeValue("rootDomainNamingContext")
	configNamingContext := entity.GetAttributeValue("configurationNamingContext")
	dnsHostName := entity.GetAttributeValue("dnsHostName")

	realm := strings.Split(entity.GetAttributeValue("ldapServiceName"), "@")[1]
	forestFunctionalLevel, _ := strconv.Atoi(entity.GetAttributeValue("forestFunctionality"))

	rootdse := RootDSE{
		currentDomain,
		rootNamingContext,
		configNamingContext,
		realm,
		dnsHostName,
		forestFunctionalLevel,
	}

	return rootdse
}
