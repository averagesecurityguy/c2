package beacon

import (
	"fmt"
	"net/http"
)

// Define an HTTP Authentication beacon.
type HttpAuthBeacon struct {
	id    string
	url   string
	agent string
}

// Send an HTTP request with the system id in the authorization bearer token.
// Read the Location header and return it if present.
func (h *HttpAuthBeacon) Ping() string {
	client := httpClient()
	req, err := http.NewRequest("GET", h.url, nil)
	if err != nil {
		return ""
	}

	req.Header.Set("User-Agent", h.agent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.id))

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	if resp.StatusCode == 401 {
		return resp.Header.Get("Location")
	}

	return ""
}
