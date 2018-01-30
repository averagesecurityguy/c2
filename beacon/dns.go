package beacon

import (
	"fmt"
	"net"
)

// Define an DNS Name Server Beacon
type DnsNsBeacon struct {
	id     string
	domain string
}

// Send a lookup request for a Name server at id.domain. Return the record if
// any.
func (d *DnsNsBeacon) Ping() string {
	rec, err := net.LookupNS(fmt.Sprintf("%s.%s", d.id, d.domain))
	if err != nil {
		return ""
	}

	return rec[0].Host
}
