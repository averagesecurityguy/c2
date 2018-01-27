package downloader

import (
	"testing"
)

func TestDownloadExec(t *testing.T) {
	d := NewDnsTxtDownloader()
	d.DownloadExec("domain.com")
}
