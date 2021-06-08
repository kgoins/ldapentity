package ad

import (
	"errors"

	"github.com/kgoins/ldapentity/entity"
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

func NewTrustFromEntity(trustEntity entity.Entity) (t Trust, err error) {
	tdInt, found, err := trustEntity.GetAsInt(ATTR_trustDirection)
	if err != nil {
		return
	}

	if !found {
		err = errors.New("Trust direction not found")
		return
	}

	t.Direction = DecodeTrustDirection(tdInt)
	t.Name, _ = trustEntity.GetSingleValuedAttribute(ATTR_name)

	return
}
