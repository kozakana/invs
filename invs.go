package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type options struct {
	port string
}

func (opts options) addr() string {
	return fmt.Sprintf(":%s", opts.port)
}

func getOptions() options {
	opts := options{}
	flag.StringVar(&opts.port, "p", "8080", "port")
	flag.Parse()
	return opts
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
}

func main() {
	opts := getOptions()
	http.HandleFunc("/", handler)
	fmt.Printf("server started on port %s\n", opts.port)
	http.ListenAndServe(opts.addr(), nil)
}
