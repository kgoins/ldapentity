package ad

import (
	"github.com/kgoins/ldapentity/entity"
)

// Domain represents an AD domain object from the directory.
type Domain struct {
	ADEntity
	SecurableEntity

	PwdPolicy       PasswordPolicy
	SupportedEtypes []KrbEtype
}

// NewDomainFromEntity builds a Domain from its ldap entry.
func NewDomainFromEntity(domainEntity entity.Entity) (d Domain, err error) {
	d.ADEntity, err = NewADEntity(domainEntity)
	if err != nil {
		return
	}

	d.SecurableEntity, err = NewSecurableEntity(domainEntity)
	if err != nil {
		return
	}

	d.PwdPolicy, err = NewPasswordPolicyFromEntity(domainEntity)
	return
}
