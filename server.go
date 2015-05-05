package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

var backendURL *url.URL

func handler(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Implement me!", http.StatusNotImplemented)
}

func main() {
	var port uint
	var backend string
	var maxConcurrency uint
	var timeout time.Duration

	flag.UintVar(&port, "port", 80, "listening port")
	flag.StringVar(&backend, "backend", "http://localhost/", "backend http service url")
	flag.DurationVar(&timeout, "timeout", 300*time.Millisecond, "request timeout deadline")
	flag.UintVar(&maxConcurrency, "concurrency", 4, "max concurrency for backend requests")

	flag.Parse()

	var err error
	backendURL, err = url.Parse(backend)
	if err != nil {
		log.Fatal("url Parse: ", err)
	}

	http.HandleFunc("/", handler)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
