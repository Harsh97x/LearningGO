package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.Handle("/", handler())
	log.Fatal(http.ListenAndServe(":3000", nil))
}
