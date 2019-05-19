package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

type options struct {
	port, filterURL, proxyHost string
}

var opts = options{}

func (opts options) addr() string {
	return fmt.Sprintf(":%s", opts.port)
}

func setOptions() {
	flag.StringVar(&opts.port, "p", "8080", "port")
	flag.StringVar(&opts.proxyHost, "proxy-host", "", "port")
	flag.StringVar(&opts.filterURL, "filter-url", "", "Output only URLs that contain the string")
	flag.Parse()
}

func index(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "200")
	if !strings.Contains(request.URL.Path, opts.filterURL) {
		return
	}
	dump, _ := httputil.DumpRequest(request, true)
	fmt.Println(string(dump))
}

func handler() *http.ServeMux {
	m := http.NewServeMux()
	m.Handle("/", http.HandlerFunc(index))
	return m
}

func rpHandler() *httputil.ReverseProxy {
	director := func(request *http.Request) {
		request.URL.Scheme = "http"
		request.URL.Host = opts.proxyHost
		if strings.Contains(request.URL.Path, opts.filterURL) {
			dump, _ := httputil.DumpRequest(request, true)
			fmt.Println(string(dump))
		}
	}
	return &httputil.ReverseProxy{Director: director}
}

func main() {
	setOptions()
	var server http.Server
	if opts.proxyHost == "" {
		server = http.Server{
			Addr:    opts.addr(),
			Handler: handler(),
		}
	} else {
		server = http.Server{
			Addr:    opts.addr(),
			Handler: rpHandler(),
		}
	}

	fmt.Printf("server starting on port %s\n", opts.port)
	if opts.proxyHost != "" {
		fmt.Printf("proxy to %s\n", opts.proxyHost)
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
