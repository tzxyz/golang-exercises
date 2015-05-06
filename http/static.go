package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("c://")))
	http.ListenAndServe(":80", nil)
}
