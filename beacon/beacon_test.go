package beacon

import (
	"testing"
)

func TestHttpBeacon(t *testing.T) {
	h := NewHttpAuthBeacon("uuid", "http://127.0.0.1:8000/", "agent")
	url := h.Ping()

	if url == "" {
		t.Error("Did not receive a URL from the server.")
	}
}

func TestDnsBeacon(t *testing.T) {
	ns := "ns.domain.com."
	d := NewDnsNsBeacon("uuid", "domain.com")
	host := d.Ping()

	if host == "" {
		t.Error("Did not receive a hostname from the server.")
	}

	if host != ns {
		t.Error("Expected", ns, "got", host)
	}
}
