# C2
The C2 repository seeks to provide a practical implementation of the ideas contained in the Red Team Infrastructure Wike at https://github.com/bluscreenofjeff/Red-Team-Infrastructure-Wiki. In particular, this repository provides two Go packages one for beaconers and another for downloaders. In addition, the repository provides sample implants that use these beaconers and downloaders. Finally, the repository contains backend DNS and HTTP servers for C2 and configuration information for building front-end redirector servers.

## Creating Implants
The implants folder has two implants that show how to use the beaconer and downloader packages. These implants can be used as a reference for building your own implants. The general idea is to select a beaconer and a downloader and wrap them in a loop with some type of timing mechanism to control how often the implant will beacon out.

## Creating Beacons and Downloaders
To create new beaconers and downloaders you must build a struct that satisfies the appropriate interface. The beaconer and downloader directories contain both HTTP and DNS examples.

## Backend Servers
The servers directory contains both a DNS server and an HTTP server that can be used as the backend server for the various beacons and downloaders. These servers were not designed with production use in mind but can be used in production if needed. In addition, the servers folder contains a payload file, which can be built using `go build -o payload.bin payload.go`. The servers expect to find a payload.bin file in the current directory, which will be served to the downloaders and executed on the client.

## Frontend Servers
The docs folder contains configuration information needed to build frontend redirectors for the C2 system. The configuration allows C2 traffic to be redirected to the C2 server based on specific criteria and sends all other traffic to a benign server.

## Contributions
I am not a C2 expert. I have been reading the Red Team Infrastructure Wiki and decided to build a practical implementation. With that said, I would love pull requests for new beaconers and downloaders. If you are not comfortable writing the code yourself, let me know what beacon or downloader techniques you would like implemented by submitting an issue.
