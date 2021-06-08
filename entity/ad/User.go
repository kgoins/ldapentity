package ad

import (
	"time"

	"github.com/kgoins/ldapentity/entity"
)

type User struct {
	ADEntity
	SecurableEntity

	UPN          string
	AccountFlags UserAccountControl

	Email string
	Title string

	LastLogon       time.Time
	LogonCount      int
	PasswordLastSet time.Time

	IsAdmin  bool
	Groups   []Group
	Services []ServicePrincipal
}

func newUserStubsFromDNs(dnList []string) []User {
	users := make([]User, 0, len(dnList))
	for _, dn := range dnList {
		user := User{ADEntity: ADEntity{DN: dn}}
		users = append(users, user)
	}

	return users
}

func NewUserFromEntry(entity entity.Entity) (usr User, err error) {
	usr.ADEntity, err = NewADEntity(entity)
	if err != nil {
		return
	}

	usr.SecurableEntity, err = NewSecurableEntity(entity)
	if err != nil {
		return
	}

	uacStr, _ := entity.GetSingleValuedAttribute(ATTR_uac)
	usr.AccountFlags, err = NewUAC(uacStr)
	if err != nil {
		return
	}

	adPwdSetTime, _ := entity.GetSingleValuedAttribute(ATTR_pwdLastSet)
	usr.PasswordLastSet = TimeFromADTimestamp(adPwdSetTime)

	lastLogonStr, _ := entity.GetSingleValuedAttribute(ATTR_lastLogon)
	lastLogonTime := TimeFromADTimestamp(lastLogonStr)

	lastLogonTimestampStr, _ := entity.GetSingleValuedAttribute(ATTR_lastLogonTimestamp)
	lastLogonTimestamp := TimeFromADTimestamp(lastLogonTimestampStr)

	if lastLogonTimestamp.Before(lastLogonTime) {
		usr.LastLogon = lastLogonTime
	} else {
		usr.LastLogon = lastLogonTimestamp
	}

	usr.LogonCount, _, err = entity.GetAsInt(ATTR_logonCount)
	if err != nil {
		return
	}

	adminCount, _, err := entity.GetAsInt(ATTR_adminCount)
	if err != nil {
		return
	}
	usr.IsAdmin = (adminCount == 1)

	groupDNList, _ := entity.GetAttribute(ATTR_memberOf)
	usr.Groups = newGroupStubsFromDNs(groupDNList.Value.Values())

	serviceStrs, _ := entity.GetAttribute(ATTR_spn)
	usr.Services = NewServicePrincipals(serviceStrs.Value.Values(), usr.SamAccountName)

	usr.UPN, _ = entity.GetSingleValuedAttribute(ATTR_userPrincipalName)
	usr.Email, _ = entity.GetSingleValuedAttribute(ATTR_mail)
	usr.Title, _ = entity.GetSingleValuedAttribute(ATTR_title)

	return
}
