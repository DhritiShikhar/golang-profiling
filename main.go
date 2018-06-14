package main

import (
	"net/http"
	_ "net/http/pprof"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
