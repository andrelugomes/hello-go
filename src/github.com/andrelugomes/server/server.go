package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "GO Server: You requested is %s", r.URL.Path)
}

func get(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "GET request ")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/get", get)

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello WWW for GO : %q", r.URL.Path)
	})

    http.ListenAndServe(":8080", nil)
}