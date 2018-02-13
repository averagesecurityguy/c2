package beacon

import (
	"testing"
)

func TestHttpBeacon(t *testing.T) {
	h := NewHttpAuthBeacon("uuid", "http://127.0.0.1:8000/", "agent")
	url := h.Beacon()

	if url == "" {
		t.Error("Did not receive a URL from the server.")
	}
}

func TestDnsNsBeacon(t *testing.T) {
	loc := "ns.domain.com."
	d := NewDnsNsBeacon("uuid", "domain.com")
	host := d.Beacon()

	if host == "" {
		t.Error("Did not receive a hostname from the server.")
	}

	if host != loc {
		t.Error("Expected", loc, "got", host)
	}
}

func TestDnsMxBeacon(t *testing.T) {
	loc := "ns.domain.com."
	d := NewDnsMxBeacon("uuid", "domain.com", 10)
	host := d.Beacon()

	if host == "" {
		t.Error("Did not receive a hostname from the server.")
	}

	if host != loc {
		t.Error("Expected", loc, "got", host)
	}
}
