package main

import (
	"fmt"
	"github.com/shinji62/route-service-cf/proxy"
	"github.com/shinji62/route-service-cf/roundTripper"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
)

var (
	port = kingpin.Flag("port", "Port to listen").Envar("PORT").Short('p').Required().Int()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	roundTripper := roundTripper.NewLoggingRoundTripper()
	proxy := proxy.NewReverseProxy(roundTripper)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%v", *port), proxy))
}
