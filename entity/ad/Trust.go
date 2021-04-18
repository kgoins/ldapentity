package sadAD

import (
	"strconv"

	ldap "gopkg.in/ldap.v2"
)

type TrustDirection int

const (
	TRUST_Disabled      TrustDirection = 0
	TRUST_Inbound       TrustDirection = 1
	TRUST_Outbound      TrustDirection = 2
	TRUST_Bidirectional TrustDirection = 3
)

var trustDirectionMap = map[TrustDirection]string{
	TRUST_Disabled:      "Disabled",
	TRUST_Inbound:       "Inbound",
	TRUST_Outbound:      "Outbound",
	TRUST_Bidirectional: "Bidirectional",
}

func (td TrustDirection) String() string {
	return trustDirectionMap[td]
}

func DecodeTrustDirection(tdInt int) TrustDirection {
	for td, _ := range trustDirectionMap {
		if int(td)&tdInt != 0 {
			return td
		}
	}

	return TrustDirection(0)
}

type Trust struct {
	Name      string
	Direction TrustDirection
}

var trustAttributes = []string{
	ATTR_name,
	ATTR_trustDirection,
}

func NewTrustFromEntity(trustEntity *ldap.Entry) Trust {
	tdInt, _ := strconv.Atoi(trustEntity.GetAttributeValue(ATTR_trustDirection))
	trustDirection := DecodeTrustDirection(tdInt)

	return Trust{
		Name:      trustEntity.GetAttributeValue(ATTR_name),
		Direction: trustDirection,
	}
}

func (ls LdapSession) GetExternalTrusts() ([]Trust, error) {
	trustEntities, err := ls.BasicSearch("(objectClass=trustedDomain)", trustAttributes, AD_pageMax)
	trusts := make([]Trust, 0, len(trustEntities))

	if err != nil {
		return trusts, err
	}

	if len(trustEntities) == 0 {
		return trusts, err
	}

	for _, trustEntity := range trustEntities {
		trusts = append(trusts, NewTrustFromEntity(trustEntity))
	}

	return trusts, nil
}
