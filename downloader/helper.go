package downloader

import (
	"io/ioutil"
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

func save(data []byte) string {
	tmpFile, err := ioutil.TempFile("", "")

	f, err := os.OpenFile(tmpFile.Name(), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}

	f.Write(data)
	f.Close()

	return tmpFile.Name()
}

func run(filename string) {
	cmd := exec.Command(filename)
	cmd.Start()
}
