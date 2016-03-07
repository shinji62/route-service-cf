package proxy

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/shinji62/route-service-cf/headers"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

func NewReverseProxy(transport http.RoundTripper, httpClient *http.Client, debug bool) *httputil.ReverseProxy {

	reverseProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			start := time.Now()
			RouterServiceheader := headers.NewRouteServiceHeaders()

			err := RouterServiceheader.ParseHeadersAndClean(&req.Header)
			if RouterServiceheader.IsValidRequest() && err == nil {
				req.URL = RouterServiceheader.ParsedUrl
				req.Host = RouterServiceheader.ParsedUrl.Host
			} else {
				req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte{}))
				req.Host = "No Host"
				log.Print("Header are not Valid")
			}

			if debug {
				dump, err := httputil.DumpRequest(req, true)
				if err != nil {
					log.Fatalln(err.Error())
				}
				log.Printf("%q", dump)
				log.Printf("Time Elapsed header %v ", time.Since(start))

			}

		},
		Transport: transport,
	}
	return reverseProxy
}
