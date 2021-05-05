package ad

import (
	"encoding/binary"
	"errors"
	"net"

	"gopkg.in/ldap.v2"
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

func decodeARecordBytes(record []byte) (net.IP, error) {
	rDataTypeBytes := []byte{record[2], record[3]}

	var rDataType uint16
	rDataType = binary.LittleEndian.Uint16(rDataTypeBytes)

	if rDataType != DNS_A_RECORD {
		return net.IP{}, errors.New("Unknown record format")
	}

	rData := record[24:28]
	addr := net.IP(rData)

	return addr, nil
}

func NewDNSZoneFromEntry(entry *ldap.Entry) DNSZone {
	return DNSZone{
		DN:     entry.DN,
		Domain: entry.GetAttributeValue(ATTR_name),
	}
}

func NewDNSRecordFromEntry(entry *ldap.Entry) (DNSRecord, error) {
	ip, err := decodeARecordBytes(entry.GetRawAttributeValue(ATTR_dnsRecord))
	if err != nil {
		return DNSRecord{}, err
	}

	return DNSRecord{
		Hostname: entry.GetAttributeValue(ATTR_name),
		Domain:   entry.GetAttributeValue(ATTR_domaincomponent),
		Addr:     ip,
	}, nil
}
