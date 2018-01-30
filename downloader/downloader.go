package downloader

// Define the Downloader interface.
type Downloader interface {
	DownloadExec(location string)
}

// Create a new HTTP Downloader.
func NewHttpDownloader(agent string) Downloader {
	h := new(HttpDownload)

	h.agent = agent

	return Downloader(h)
}

// Create a new DNS TXT record downloader.
func NewDnsTxtDownloader() Downloader {
	return Downloader(new(DnsTxtDownload))
}
