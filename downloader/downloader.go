package downloader

// Define the Downloader interface.
type Downloader interface {
	DownloadExec(location string)
}

// Create a new HTTP Downloader.
func NewHttpDownloader(agent string) *HttpDownload {
	h := new(HttpDownload)

	h.agent = agent

	return h
}

// Create a new DNS TXT record downloader.
func NewDnsTxtDownloader() *DnsTxtDownload {
	return new(DnsTxtDownload)
}
