package beacon

type Beacon interface {
	Ping() string
}

func NewHttpAuthBeacon(sysid, url, agent string) Beacon {
	h := new(HttpAuthBeacon)

	h.id = sysid
	h.agent = agent
	h.url = url

	return Beacon(h)
}

func NewDnsBeacon(sysid, domain string) Beacon {
	d := new(DnsBeacon)

	d.id = sysid
	d.domain = domain

	return Beacon(d)
}
