package main

import (
	"net/http"

	"github.com/robertding/The-Go-Programming-Language/ch3/surface"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(surface.Draw()))
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}
