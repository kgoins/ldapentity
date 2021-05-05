package ad

import (
	"github.com/kgoins/ldapentity/entity"
)

type PasswordPolicy struct {
	MinLength     int
	HistoryLength int

	LockoutThreshold int
	LockoutDuration  int

	MinAge int
	MaxAge int
}

func NewPasswordPolicyFromEntity(policyEntity entity.Entity) (p PasswordPolicy, err error) {
	p.MinLength, _, err = policyEntity.GetAsInt(ATTR_minPwdLen)
	if err != nil {
		return
	}

	p.HistoryLength, _, err = policyEntity.GetAsInt(ATTR_historyLen)
	if err != nil {
		return
	}

	p.LockoutThreshold, _, err = policyEntity.GetAsInt(ATTR_lockoutThreshold)
	if err != nil {
		return
	}

	durStr, _ := policyEntity.GetSingleValuedAttribute(ATTR_minPwdAge)
	p.LockoutDuration = ADIntervalToMins(durStr)

	minAgeStr, _ := policyEntity.GetSingleValuedAttribute(ATTR_minPwdAge)
	p.MinAge = ADIntervalToDays(minAgeStr)

	maxAgeStr, _ := policyEntity.GetSingleValuedAttribute(ATTR_maxPwdAge)
	p.MaxAge = ADIntervalToDays(maxAgeStr)
	return
}
