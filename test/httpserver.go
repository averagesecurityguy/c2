package main

import (
    "log"
    "net/http"
    "io/ioutil"
)

func executed(w http.ResponseWriter, r *http.Request) {
    log.Print(r.RequestURI)
}

func exec(w http.ResponseWriter, r *http.Request) {
    log.Print(r.RequestURI)

    data, err := ioutil.ReadFile("/Users/shaywood/exec")
    if err != nil {
        log.Print("Failed to read file.")
    } else {
        w.Write(data)
    }
}

func redirect(w http.ResponseWriter, r *http.Request) {
    log.Print(r.RequestURI)

    http.Redirect(w, r, "http://127.0.0.1:8000/exec", 302)
}

func main() {
    http.HandleFunc("/", redirect)
    http.HandleFunc("/exec", exec)
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
