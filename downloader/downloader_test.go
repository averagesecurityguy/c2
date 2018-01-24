package downloader

import (
	"testing"
)

func TestDownloadExec(t *testing.T) {
	d := NewHttpDownloader("agent")
	d.DownloadExec("http://127.0.0.1:8000/exec")
}
