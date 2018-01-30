// Simple HTTP server that can be used to test the HTTP beacons and downloaders.
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func executed(w http.ResponseWriter, r *http.Request) {
	log.Print(r.RequestURI)
}

func exec(w http.ResponseWriter, r *http.Request) {
	log.Print(r.RequestURI)

	data, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Print("Failed to read file.")
	} else {
		w.Write(data)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	log.Print(r.RequestURI)

	w.Header().Add("Location", "http://127.0.0.1:8000/exec")

	http.Error(w, "Not Authorized", 401)
}

func main() {
	http.HandleFunc("/", redirect)
	http.HandleFunc("/exec", exec)
	http.HandleFunc("/executed", executed)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
