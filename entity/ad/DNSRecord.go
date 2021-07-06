package ad

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"net"

	"github.com/kgoins/ldapentity/entity"
)

const (
	DNS_A_RECORD     uint16 = 1
	DNS_NS_RECORD    uint16 = 2
	DNS_CNAME_RECORD uint16 = 5
	DNS_SOA_RECORD   uint16 = 6
	DNS_SRV_RECORD   uint16 = 33
)

// DNSRecord represents a minimal DNS Record according to AD
type DNSRecord struct {
	Hostname string
	Domain   string
	Addr     net.IP
}

// DNSZone maps the DN of the DNS Zone object in AD to it's domain
type DNSZone struct {
	DN     string
	Domain string
}

func decodeARecordBytesFromB64(recordB64 string) (ip net.IP, err error) {
	record, err := base64.StdEncoding.DecodeString(recordB64)
	if err != nil {
		return
	}

	return decodeARecordBytes(record)
}

func decodeARecordBytes(record []byte) (net.IP, error) {
	rDataTypeBytes := []byte{record[2], record[3]}

	rDataType := binary.LittleEndian.Uint16(rDataTypeBytes)

	if rDataType != DNS_A_RECORD {
		return net.IP{}, errors.New("unknown record format")
	}

	rData := record[24:28]
	addr := net.IP(rData)

	return addr, nil
}

func NewDNSZoneFromEntry(entry entity.Entity) (zone DNSZone, err error) {
	found := true
	zone.DN, found = entry.GetDN()
	if !found {
		err = errors.New("unable to get DN")
		return
	}

	zone.Domain, _ = entry.GetSingleValuedAttribute(ATTR_name)
	return
}

func NewDNSRecordFromEntity(entry entity.Entity) (r DNSRecord, err error) {
	recordStr, found := entry.GetSingleValuedAttribute(ATTR_dnsRecord)
	if !found {
		err = errors.New("unable to get record value")
		return
	}

	r.Addr, err = decodeARecordBytesFromB64(recordStr)
	if err != nil {
		return
	}

	r.Hostname, _ = entry.GetSingleValuedAttribute(ATTR_name)
	r.Domain, _ = entry.GetSingleValuedAttribute(ATTR_domaincomponent)
	return
}
