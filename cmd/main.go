package main

import (
    "log"
    "net/http"
    "scalable-api/pkg"
)

func main() {
    http.HandleFunc("/echo", pkg.EchoHandler)
    log.Println("Starting server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
