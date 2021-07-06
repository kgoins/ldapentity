package ad

import (
	"fmt"
	"os"
	"regexp"

	ldap "gopkg.in/ldap.v2"
)

const AdGeneralizedTimeFmt string = "20060102150405-0700"

var guidRegex *regexp.Regexp

func init() {
	var err error
	guidRegex, err = regexp.Compile(`[\da-zA-Z]{8}-([\da-zA-Z]{4}-){3}[\da-zA-Z]{12}`)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
}

/*
TODO: this is a better solution, but requires translation to RE2 or PCRE lib
var ldapRegex, _ = regexp.Compile(`/^(\s*\((?:[&|](?1)+|(?:!(?1))|[a-zA-Z][a-zA-Z0-9-]*[<>~]?=[^()]*)\s*\)\s*)$/`)

func IsValidLdapQuery(ldapFilter string) bool {
	return ldapRegex.MatchString(ldapFilter)
}
*/

func IsValidLdapQuery(ldapFilter string) bool {
	_, err := ldap.CompileFilter(ldapFilter)
	return (err == nil)
}

func IsGuid(guidStr string) bool {
	return guidRegex.MatchString(guidStr)
}
