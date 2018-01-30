package beacon

import (
	"fmt"
	"net"
)

type DnsNsBeacon struct {
	id     string
	domain string
}

func (d *DnsNsBeacon) Ping() string {
	rec, err := net.LookupNS(fmt.Sprintf("%s.%s", d.id, d.domain))
	if err != nil {
		return ""
	}

	return rec[0].Host
}
