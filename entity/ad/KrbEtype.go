package ad

import (
	"strings"
)

type KrbEtype int

const (
	DES_CBC_CRC             KrbEtype = 0x01
	DES_CBC_MD5             KrbEtype = 0x02
	RC4_HMAC                KrbEtype = 0x04
	AES128_CTS_HMAC_SHA1_96 KrbEtype = 0x08
	AES256_CTS_HMAC_SHA1_96 KrbEtype = 0x10
)

var etypeMap = map[KrbEtype]string{
	DES_CBC_CRC:             "DES-CBC-CRC",
	DES_CBC_MD5:             "DES-CBC-MD5",
	RC4_HMAC:                "ARCFOUR-HMAC-MD5",
	AES128_CTS_HMAC_SHA1_96: "AES128-CTS-HMAC-SHA1-96",
	AES256_CTS_HMAC_SHA1_96: "AES256-CTS-HMAC-SHA1-96",
}

func (kt KrbEtype) String() string {
	return etypeMap[kt]
}

func DecodeKrbEtypes(intEtypes int) KrbEtypeList {
	eTypes := []KrbEtype{}
	for eType := range etypeMap {
		if int(eType)&intEtypes != 0 {
			eTypes = append(eTypes, eType)
		}
	}

	return eTypes
}

type KrbEtypeList []KrbEtype

func (eTypes KrbEtypeList) String() string {
	eTypeStr := make([]string, 0, len(eTypes))
	for _, eType := range eTypes {
		eTypeStr = append(eTypeStr, eType.String())
	}

	return strings.Join(eTypeStr, " ")
}

func (eTypes KrbEtypeList) Weakest() KrbEtype {
	minVal := AES256_CTS_HMAC_SHA1_96
	for _, eType := range eTypes {
		if eType < minVal {
			minVal = eType
		}
	}

	return minVal
}
