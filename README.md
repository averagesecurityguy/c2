# C2
The C2 package provides beacons and downloaders that can be used in implants. The beacons and downloaders are designed to work together but can be used separately.

## Usage
To use the C2 package, first, import the beacon and/or downloader package. Next, create a new beacon, which will ping the given C2 server. When the server is ready to activate the beacon it should respond with the URL of the file that should be downloaded and executed. If desired, create a new downloader, which takes a URL and download and executes file at the URL.

## Examples
Check the examples folder for detailed examples, including an HTTP server that can be used for testing the examples.

## Custom Beacons and Downloaders
You can create your own beacons and downloaders by creating a struct that satisfies the appropriate interface.
