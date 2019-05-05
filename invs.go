package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
