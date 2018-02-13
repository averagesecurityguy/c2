package beacon

import (
	"fmt"
	"net"
)

// Define an DNS Name Server Beacon
type DnsNsBeacon struct {
	id       string
	domain   string
}

// Send a lookup request for a Name server at id.domain. Return the hostname of
// the first record. If there are no records return an empty string.
func (d *DnsNsBeacon) Beacon() string {
	rec, err := net.LookupNS(fmt.Sprintf("%s.%s", d.id, d.domain))
	if err != nil {
		return ""
	}

	return rec[0].Host
}

// Define an DNS MX Server Beacon
type DnsMxBeacon struct {
	id     string
	domain string
	pref   uint16
}

// Send a lookup request for an MX server at id.domain. Return the hostname of
// the record with the matching preference. If there are no records return an
// empty string.
func (d *DnsMxBeacon) Beacon() string {
	rec, err := net.LookupMX(fmt.Sprintf("%s.%s", d.id, d.domain))
	if err != nil {
		return ""
	}

	for i := range rec {
		if rec[i].Pref == d.pref {
			return rec[i].Host
		}
	}

	return ""
}
