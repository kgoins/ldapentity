package ad

import (
	"strconv"
	"time"

	"github.com/kgoins/ldapentity/entity"
)

type User struct {
	ADEntity

	UPN            string
	SamAccountName string

	AccountFlags UserAccountControl

	SID   string
	Email string

	LastLogon       time.Time
	LogonCount      int
	PasswordLastSet time.Time

	IsAdmin bool
	Groups  []Group

	Title string

	Services []ServicePrincipal
}

func (user User) IsEmpty() bool {
	return user.SamAccountName == ""
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
	uacStr, _ := entity.GetSingleValuedAttribute(ATTR_uac)
	acctFlags, err := NewUAC(uacStr)

	sidB64, _ := entity.GetSingleValuedAttribute(ATTR_sid)
	sid, err := SidFromBase64(sidB64)
	if err != nil {
		return
	}

	adPwdSetTime, _ := entity.GetSingleValuedAttribute(ATTR_pwdLastSet)
	passLastSet := TimeFromADTimestamp(adPwdSetTime)

	lastLogonStr, _ := entity.GetSingleValuedAttribute(ATTR_lastLogon)
	lastLogonTime := TimeFromADTimestamp(lastLogonStr)

	lastLogonTimestampStr, _ := entity.GetSingleValuedAttribute(ATTR_lastLogonTimestamp)
	lastLogonTimestamp := TimeFromADTimestamp(lastLogonTimestampStr)

	var lastLogon time.Time
	if lastLogonTimestamp.Before(lastLogonTime) {
		lastLogon = lastLogonTime
	} else {
		lastLogon = lastLogonTimestamp
	}

	logonCountStr, _ := entity.GetSingleValuedAttribute(ATTR_logonCount)
	logonCount, err := strconv.Atoi(logonCountStr)
	if err != nil {
		return
	}

	adminCountStr, _ := entity.GetSingleValuedAttribute(ATTR_adminCount)
	adminCount, err := strconv.Atoi(adminCountStr)
	if err != nil {
		return
	}
	isAdmin := (adminCount == 1)

	groupDNList, _ := entity.GetAttribute(ATTR_memberOf)
	groups := newGroupStubsFromDNs(groupDNList.Value.Values())

	samaccountname, _ := entity.GetSingleValuedAttribute(ATTR_sAMAccountName)

	serviceStrs, _ := entity.GetAttribute(ATTR_spn)
	services := NewServicePrincipals(serviceStrs.Value.Values(), samaccountname)

	upn, _ := entity.GetSingleValuedAttribute(ATTR_userPrincipalName)
	mail, _ := entity.GetSingleValuedAttribute(ATTR_mail)
	title, _ := entity.GetSingleValuedAttribute(ATTR_title)

	adEntity, err := NewADEntity(entity)
	if err != nil {
		return
	}

	return User{
		ADEntity: adEntity,

		UPN:            upn,
		SamAccountName: samaccountname,

		AccountFlags: acctFlags,

		SID:   sid,
		Email: mail,

		LastLogon:       lastLogon,
		LogonCount:      logonCount,
		PasswordLastSet: passLastSet,

		IsAdmin: isAdmin,

		Title: title,

		Groups:   groups,
		Services: services,
	}, nil
}
