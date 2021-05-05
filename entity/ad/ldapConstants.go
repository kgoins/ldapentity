package ad

const (
	ATTR_DN                 = "distinguishedName"
	ATTR_CN                 = "cn"
	ATTR_name               = "name"
	ATTR_displayName        = "displayName"
	ATTR_description        = "description"
	ATTR_sid                = "objectSid"
	ATTR_created            = "whenCreated"
	ATTR_changed            = "whenChanged"
	ATTR_minPwdLen          = "minPwdLength"
	ATTR_historyLen         = "pwdHistoryLength"
	ATTR_lockoutThreshold   = "lockoutThreshold"
	ATTR_lockoutDuration    = "lockoutDuration"
	ATTR_logonCount         = "logonCount"
	ATTR_maxPwdAge          = "maxPwdAge"
	ATTR_minPwdAge          = "minPwdAge"
	ATTR_pwdLastSet         = "pwdLastSet"
	ATTR_lastLogon          = "lastLogon"
	ATTR_lastLogonTimestamp = "lastLogonTimestamp"
	ATTR_trustDirection     = "trustDirection"
	ATTR_dnsRecord          = "dnsRecord"
	ATTR_sAMAccountName     = "sAMAccountName"
	ATTR_userPrincipalName  = "userPrincipalName"
	ATTR_spn                = "servicePrincipalName"
	ATTR_dnsHostname        = "dNSHostName"
	ATTR_os                 = "operatingSystem"
	ATTR_osVersion          = "operatingSystemVersion"
	ATTR_domaincomponent    = "dc"
	ATTR_memberOf           = "memberOf"
	ATTR_member             = "member"
	ATTR_mail               = "mail"
	ATTR_title              = "title"
	ATTR_adminCount         = "adminCount"
	ATTR_uac                = "userAccountControl"
	ATTR_groupType          = "groupType"
)

var (
	AD_pageMax uint32 = 1000
)
