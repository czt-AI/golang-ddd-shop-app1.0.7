package network

import (
	"net/http"
	"net/http/httputil"
)

type CDN struct {
	Backend *url.URL
}

func NewCDN(backend string) (*CDN, error) {
	parsedURL, err := url.Parse(backend)
	if err != nil {
		return nil, err
	}

	return &CDN{
		Backend: parsedURL,
	}, nil
}

func (c *CDN) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(c.Backend)
	proxy.ServeHTTP(w, r)
}