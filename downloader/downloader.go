package downloader

type Downloader interface {
	DownloadExec(location string)
}

func NewHttpDownloader(agent string) Downloader {
	h := new(HttpDownload)

	h.agent = agent

	return Downloader(h)
}

func NewDnsTxtDownloader() Downloader {
	return Downloader(new(DnsTxtDownload))
}
