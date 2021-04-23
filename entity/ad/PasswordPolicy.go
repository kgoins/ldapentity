package ad

import (
	"strconv"

	ldap "gopkg.in/ldap.v2"
)

type PasswordPolicy struct {
	MinLength        int
	HistoryLength    int
	LockoutThreshold int
	LockoutDuration  int
	MinAge           int
	MaxAge           int
}

var passwordPolicyAttributes = []string{
	ATTR_minPwdLen,
	ATTR_historyLen,
	ATTR_lockoutThreshold,
	ATTR_lockoutDuration,
	ATTR_minPwdAge,
	ATTR_maxPwdAge,
}

func (p PasswordPolicy) Empty() bool {
	checksum := p.HistoryLength + p.LockoutDuration + p.LockoutThreshold
	checksum += p.MaxAge + p.MinAge + p.MinLength

	return (checksum == 0)
}

func NewPasswordPolicyFromEntry(policy *ldap.Entry) PasswordPolicy {
	minLen, _ := strconv.Atoi(policy.GetAttributeValue(ATTR_minPwdLen))
	histLen, _ := strconv.Atoi(policy.GetAttributeValue(ATTR_historyLen))

	lockoutThreshold, _ := strconv.Atoi(policy.GetAttributeValue(ATTR_lockoutThreshold))
	lockoutDuration := ADIntervalToMins(policy.GetAttributeValue(ATTR_lockoutDuration))

	minAge := ADIntervalToDays(policy.GetAttributeValue(ATTR_minPwdAge))
	maxAge := ADIntervalToDays(policy.GetAttributeValue(ATTR_maxPwdAge))

	return PasswordPolicy{
		MinLength:        minLen,
		HistoryLength:    histLen,
		LockoutThreshold: lockoutThreshold,
		LockoutDuration:  lockoutDuration,
		MinAge:           minAge,
		MaxAge:           maxAge,
	}
}
