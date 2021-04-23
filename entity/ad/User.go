package ad

import (
	"strconv"
	"time"

	ldap "gopkg.in/ldap.v2"
)

type UACFlags struct {
	Enabled              bool
	LockedOut            bool
	MustChangePassword   bool
	SmartcardRequired    bool
	PasswordNeverExpires bool
}

type User struct {
	DN             string
	CN             string
	UPN            string
	SamAccountName string
	DisplayName    string

	AccountFlags UserAccountControl

	SID   string
	Email string

	WhenCreated time.Time
	WhenChanged time.Time

	LastLogon       time.Time
	LogonCount      int
	PasswordLastSet time.Time

	IsAdmin bool
	Groups  []Group

	Title       string
	Description string

	Services []ServicePrincipal
}

var userAttributes = []string{
	"distinguishedName",
	"cn",
	"samaccountname",
	"userPrincipalName",
	"displayName",

	"userAccountControl",
	"objectSid",
	"mail",

	"lastLogon",
	"lastLogonTimestamp",
	"logonCount",
	"pwdLastSet",

	"whenCreated",
	"whenChanged",

	"adminCount",
	"memberOf",

	"title",
	"description",

	"servicePrincipalName",
}

func (user User) IsEmpty() bool {
	return user.SamAccountName == ""
}

func newUserStubsFromDNs(dnList []string) []User {
	users := make([]User, 0, len(dnList))
	for _, dn := range dnList {
		users = append(users, User{DN: dn})
	}

	return users
}

func NewUserFromEntry(userEntry *ldap.Entry) (usr User, err error) {
	uacStr := userEntry.GetAttributeValue("userAccountControl")
	acctFlags, err := NewUAC(uacStr)

	sidBytes := userEntry.GetRawAttributeValue("objectSid")
	sid := SidFromBytes(sidBytes)

	adCreatedTime := userEntry.GetAttributeValue("whenCreated")
	created, err := TimeFromADGeneralizedTime(adCreatedTime)
	if err != nil {
		return
	}

	adChangedTime := userEntry.GetAttributeValue("whenChanged")
	changed, _ := TimeFromADGeneralizedTime(adChangedTime)

	adPwdSetTime := userEntry.GetAttributeValue("pwdLastSet")
	passLastSet := TimeFromADTimestamp(adPwdSetTime)

	lastLogonStr := userEntry.GetAttributeValue("lastLogon")
	lastLogonTime := TimeFromADTimestamp(lastLogonStr)

	lastLogonTimestampStr := userEntry.GetAttributeValue("lastLogonTimestamp")
	lastLogonTimestamp := TimeFromADTimestamp(lastLogonTimestampStr)

	var lastLogon time.Time
	if lastLogonTimestamp.Before(lastLogonTime) {
		lastLogon = lastLogonTime
	} else {
		lastLogon = lastLogonTimestamp
	}

	logonCount, err := strconv.Atoi(userEntry.GetAttributeValue("logonCount"))
	if err != nil {
		return
	}

	adminCount, err := strconv.Atoi(userEntry.GetAttributeValue("adminCount"))
	if err != nil {
		return
	}
	isAdmin := (adminCount == 1)

	groupDNList := userEntry.GetAttributeValues("memberOf")
	groups := newGroupStubsFromDNs(groupDNList)

	samaccountname := userEntry.GetAttributeValue("sAMAccountName")

	serviceStrs := userEntry.GetAttributeValues("servicePrincipalName")
	services := NewServicePrincipals(serviceStrs, samaccountname)

	usr = User{
		DN:             userEntry.GetAttributeValue("distinguishedName"),
		CN:             userEntry.GetAttributeValue("cn"),
		UPN:            userEntry.GetAttributeValue("userPrincipalName"),
		SamAccountName: samaccountname,
		DisplayName:    userEntry.GetAttributeValue("displayName"),

		AccountFlags: acctFlags,

		SID:   sid,
		Email: userEntry.GetAttributeValue("mail"),

		WhenCreated: created,
		WhenChanged: changed,

		LastLogon:       lastLogon,
		LogonCount:      logonCount,
		PasswordLastSet: passLastSet,

		IsAdmin: isAdmin,

		Title:       userEntry.GetAttributeValue("title"),
		Description: userEntry.GetAttributeValue("description"),

		Groups:   groups,
		Services: services,
	}

	return
}
