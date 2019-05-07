package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

type options struct {
	port      string
	filterUrl string
}

var opts = options{}

func (opts options) addr() string {
	return fmt.Sprintf(":%s", opts.port)
}

func setOptions() {
	flag.StringVar(&opts.port, "p", "8080", "port")
	flag.StringVar(&opts.filterUrl, "filter-url", "", "Output only URLs that contain the string")
	flag.Parse()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
	if !strings.Contains(r.URL.Path, opts.filterUrl) {
		return
	}
	dump, _ := httputil.DumpRequest(r, true)
	fmt.Println(string(dump))
}

func main() {
	setOptions()
	http.HandleFunc("/", handler)
	fmt.Printf("server started on port %s\n", opts.port)
	http.ListenAndServe(opts.addr(), nil)
}
