package ad

import (
	"encoding/base64"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

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
	if err != nil {
		return false
	}

	return true
}

func IsGuid(guidStr string) bool {
	return guidRegex.MatchString(guidStr)
}

func SidFromBase64(b64Sid string) (string, error) {
	sidBytes, err := base64.StdEncoding.DecodeString(b64Sid)
	if err != nil {
		return "", err
	}

	sid := SidFromBytes(sidBytes)
	return sid, nil
}

func SidFromBytes(sidBytes []byte) string {
	var sidBuilder strings.Builder
	sidBuilder.WriteString("S-")

	revision := int(sidBytes[0])
	sidBuilder.WriteString(strconv.Itoa(revision))

	numSubAuths := int(sidBytes[1] & 0xFF)
	authority := int64(0)

	for i := 2; i <= 7; i++ {
		authority |= int64(uint64(sidBytes[i]) << uint64(8*(5-(i-2))))
	}

	sidBuilder.WriteString("-")
	sidBuilder.WriteString(strconv.FormatInt(authority, 16))

	offset := 8
	subAuthSize := 4

	for j := 0; j < numSubAuths; j++ {
		subAuthority := int64(0)
		for k := 0; k < subAuthSize; k++ {
			subAuthority |= int64(uint64(sidBytes[offset+k]&0xFF) << uint64(8*k))
		}

		sidBuilder.WriteString("-")
		sidBuilder.WriteString(strconv.FormatInt(subAuthority, 10))

		offset += subAuthSize
	}

	sid := sidBuilder.String()
	return sid
}
