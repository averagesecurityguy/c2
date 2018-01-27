package beacon

import (
	"net"
)

type DnsCnameBeacon struct {
	id     string
	domain string
}

func (d *DnsCnameBeacon) Ping() string {
	recs, err := net.LookupCNAME(d.id + "." + d.domain)
	if err != nil {
		return ""
	}

	return recs[0]
}
