package httpClient

import (
	"golang.org/x/net/http/httpproxy"
	"net"
	"net/http"
	"net/url"
	"time"
)

var httpClient *http.Client

func Setup(userAgent string, timeout int) {
	httpClient = newHTTPClient(userAgent, timeout)
}

func MustGetHttpClient() *http.Client {
	if httpClient == nil {
		panic("httpClient: Must call `Setup` before get http client.")
	}
	return httpClient
}

// newHTTPClient initializes and returns a http client that
// will use the specified userAgent to send http requests
// and times out after timeout seconds.
func newHTTPClient(userAgent string, timeout int) *http.Client {
	proxyFunc := httpproxy.FromEnvironment().ProxyFunc()
	return &http.Client{
		Transport: &http.Transport{
			Proxy: func(r *http.Request) (uri *url.URL, err error) {
				r.Header.Set("User-Agent", userAgent)
				return proxyFunc(r.URL)
			},
			DialContext: (&net.Dialer{
				Timeout: time.Duration(timeout) * time.Second,
			}).DialContext,
		},
	}
}
