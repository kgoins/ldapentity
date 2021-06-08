package ad

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/kgoins/ldapentity/entity"
)

// SID is a microsoft security identifier string in S- format
type SID string

func NewSIDFromEntity(entity entity.Entity) (sid SID, err error) {
	sidB64, _ := entity.GetSingleValuedAttribute(ATTR_sid)

	sidStr, err := SidFromBase64(sidB64)
	if err != nil {
		return
	}

	return SID(sidStr), nil
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
