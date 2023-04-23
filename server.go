package main

import (
    "fmt"
    "log"
    "net/http"
)

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/"+r.URL.Path[1:])
}
func main() {
    fmt.Printf("Starting server at port 8080\n")
    http.HandleFunc("/", fileServerHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

