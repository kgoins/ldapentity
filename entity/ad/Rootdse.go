package sadAD

import (
	"errors"
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

func rootDSEFromSearchResp(resp *ldap.SearchResult) RootDSE {
	respObj := resp.Entries[0]

	currentDomain := respObj.GetAttributeValue("defaultNamingContext")
	rootNamingContext := respObj.GetAttributeValue("rootDomainNamingContext")
	configNamingContext := respObj.GetAttributeValue("configurationNamingContext")
	dnsHostName := respObj.GetAttributeValue("dnsHostName")

	realm := strings.Split(respObj.GetAttributeValue("ldapServiceName"), "@")[1]
	forestFunctionalLevel, _ := strconv.Atoi(respObj.GetAttributeValue("forestFunctionality"))

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

func (ls *LdapSession) GetRootDSE() (RootDSE, error) {
	controls := []ldap.Control{}
	bindReq := ldap.NewSimpleBindRequest("", "", controls)

	_, err := ls.ldapConn.SimpleBind(bindReq)
	if err != nil {
		return RootDSE{}, errors.New("Error binding to ldap server: " + err.Error())
	}

	searchReq := ldap.NewSearchRequest(
		"",
		ldap.ScopeBaseObject,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(objectClass=*)",
		[]string{},
		nil,
	)

	searchResp, err := ls.ldapConn.Search(searchReq)
	if err != nil {
		return RootDSE{}, errors.New("Error searching server: " + err.Error())
	}

	rootdse := rootDSEFromSearchResp(searchResp)
	return rootdse, nil
}
