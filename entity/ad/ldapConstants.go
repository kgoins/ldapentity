package ad

const (
	ATTR_DN                 string = "distinguishedName"
	ATTR_CN                 string = "cn"
	ATTR_name               string = "name"
	ATTR_displayName        string = "displayName"
	ATTR_description        string = "description"
	ATTR_sid                string = "objectSid"
	ATTR_created            string = "whenCreated"
	ATTR_changed            string = "whenChanged"
	ATTR_minPwdLen          string = "minPwdLength"
	ATTR_historyLen         string = "pwdHistoryLength"
	ATTR_lockoutThreshold   string = "lockoutThreshold"
	ATTR_lockoutDuration    string = "lockoutDuration"
	ATTR_logonCount         string = "logonCount"
	ATTR_maxPwdAge          string = "maxPwdAge"
	ATTR_minPwdAge          string = "minPwdAge"
	ATTR_pwdLastSet         string = "pwdLastSet"
	ATTR_lastLogon          string = "lastLogon"
	ATTR_lastLogonTimestamp string = "lastLogonTimestamp"
	ATTR_trustDirection     string = "trustDirection"
	ATTR_dnsRecord          string = "dnsRecord"
	ATTR_sAMAccountName     string = "sAMAccountName"
	ATTR_userPrincipalName  string = "userPrincipalName"
	ATTR_spn                string = "servicePrincipalName"
	ATTR_dnsHostname        string = "dNSHostName"
	ATTR_os                 string = "operatingSystem"
	ATTR_osVersion          string = "operatingSystemVersion"
	ATTR_domaincomponent    string = "dc"
	ATTR_memberOf           string = "memberOf"
	ATTR_member             string = "member"
	ATTR_mail               string = "mail"
	ATTR_title              string = "title"
	ATTR_adminCount         string = "adminCount"
	ATTR_uac                string = "userAccountControl"
	ATTR_groupType          string = "groupType"
)

var (
	AD_pageMax uint32 = 1000
)
