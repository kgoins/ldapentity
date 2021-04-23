package ad

const (
	ATTR_DN               = "distinguishedName"
	ATTR_name             = "name"
	ATTR_sid              = "objectSid"
	ATTR_created          = "whenCreated"
	ATTR_changed          = "whenChanged"
	ATTR_minPwdLen        = "minPwdLength"
	ATTR_historyLen       = "pwdHistoryLength"
	ATTR_lockoutThreshold = "lockoutThreshold"
	ATTR_lockoutDuration  = "lockoutDuration"
	ATTR_maxPwdAge        = "maxPwdAge"
	ATTR_minPwdAge        = "minPwdAge"
	ATTR_trustDirection   = "trustDirection"
	ATTR_dnsRecord        = "dnsRecord"
	ATTR_sAMAccountName   = "sAMAccountName"
	ATTR_spn              = "servicePrincipalName"
	ATTR_dnsHostname      = "dNSHostName"
	ATTR_os               = "operatingSystem"
	ATTR_osVersion        = "operatingSystemVersion"
	ATTR_domaincomponent  = "dc"
)

var (
	AD_pageMax uint32 = 1000
)
