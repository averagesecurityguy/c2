// Implement a simple HTTP server that can be used to test beacons and downloaders.
// This should not be used in production.
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Log a request to /executed. This shows that our test.go file was executed.
func executed(w http.ResponseWriter, r *http.Request) {
	log.Print(r.RequestURI)
}

// Load our payload and return it to the requestor.
func exec(w http.ResponseWriter, r *http.Request) {
	log.Print(r.RequestURI)

	data, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Print("Failed to read file.")
	} else {
		w.Write(data)
	}
}

// Redirect the client to the location of the payload.
func redirect(w http.ResponseWriter, r *http.Request) {
	log.Print(r.RequestURI)

	w.Header().Add("Location", "http://127.0.0.1:8000/exec")

	http.Error(w, "Not Authorized", 401)
}

// Setup our HTTP server and route handlers.
func main() {
	http.HandleFunc("/", redirect)
	http.HandleFunc("/exec", exec)
	http.HandleFunc("/executed", executed)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
