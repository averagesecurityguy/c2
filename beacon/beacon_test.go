package beacon

import (
    "testing"
)

func TestBeacon(t *testing.T) {
    h := NewHttpAuthBeacon("uuid", "http://127.0.0.1:8000/", "agent")
    url := h.Ping()

    if url == "" {
        t.Error("Did not receive a URL from the server.")
    }
}
