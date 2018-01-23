package beacon

import (
	"net"
)

type DnsBeacon struct {
	id     string
	domain string
}

func (d *DnsBeacon) Ping() string {
	recs, err := net.LookupTXT(d.id + "." + d.domain)
	if err != nil {
		return ""
	}

	return recs[0]
}
