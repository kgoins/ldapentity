package sadAD

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
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

var dnsRecordAttrs = []string{
	ATTR_name,
	ATTR_dnsRecord,
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

// GetAllZones returns all DNSZones in the current domain
func (ls LdapSession) GetAllZones() ([]DNSZone, error) {
	zonesDN := fmt.Sprintf(DN_domainDnsZones, ls.RootDSE.CurrentDomain)
	params, err := BuildSearchParams(QUERY_DNSZones, []string{ATTR_name}, zonesDN, ScopeSubtree, AD_pageMax)
	if err != nil {
		return nil, err
	}

	zoneEntries, err := ls.Search(params)
	if err != nil {
		return []DNSZone{}, err
	}

	zones := make([]DNSZone, 0, len(zoneEntries))
	for _, zoneEntry := range zoneEntries {
		zone := DNSZone{
			DN:     zoneEntry.DN,
			Domain: zoneEntry.GetAttributeValue(ATTR_name),
		}
		zones = append(zones, zone)
	}

	return zones, nil
}

func (ls LdapSession) getRecordsInZone(zone DNSZone) ([]DNSRecord, error) {
	params, err := BuildSearchParams(QUERY_DNSRecords, dnsRecordAttrs, zone.DN, ScopeSubtree, AD_pageMax)
	if err != nil {
		return nil, err
	}

	dnsRecordEntries, err := ls.Search(params)
	if err != nil {
		return nil, err
	}

	var dnsRecord DNSRecord
	dnsRecords := make([]DNSRecord, 0, len(dnsRecordEntries))

	for _, entry := range dnsRecordEntries {
		ip, err := decodeARecordBytes(entry.GetRawAttributeValue(ATTR_dnsRecord))
		if err != nil {
			continue
		}

		dnsRecord = DNSRecord{
			Hostname: entry.GetAttributeValue(ATTR_name),
			Domain:   zone.Domain,
			Addr:     ip,
		}

		dnsRecords = append(dnsRecords, dnsRecord)
	}

	return dnsRecords, nil
}

// GetAllARecords extracts every DNS A Record in every zone the current domain knows about.
func (ls LdapSession) GetAllARecords() ([]DNSRecord, error) {
	zones, err := ls.GetAllZones()
	if err != nil {
		return []DNSRecord{}, err
	}

	var allRecords []DNSRecord
	for _, zone := range zones {
		zoneRecords, err := ls.getRecordsInZone(zone)
		if err != nil {
			continue
		}

		allRecords = append(allRecords, zoneRecords...)
	}

	return allRecords, nil
}
