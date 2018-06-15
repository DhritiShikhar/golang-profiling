package main

import (
	"net/http"

	_ "runtime/pprof"

	"github.com/pkg/profile"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!"))
}

func main() {
	defer profile.Start().Stop()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
