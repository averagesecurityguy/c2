// The testing infrastructure expects a test.bin file in the examples directory.
// You can compile this file using go build -o test.bin test.go or you can create
// your own payload called test.bin in the examples directory.
//
// This file will send a GET request to an HTTP server to prove the payload
// executed.
package main

import (
	"net/http"
)

func main() {

	http.Get("http://127.0.0.1:8000/executed")
}
