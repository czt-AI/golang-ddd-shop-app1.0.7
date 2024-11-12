package network

import (
	"net/http"
	"net/http/httputil"
)

type Proxy struct {
	Target *url.URL
}

func NewProxy(target string) (*Proxy, error) {
	parsedURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	return &Proxy{
		Target: parsedURL,
	}, nil
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(p.Target)
	proxy.ServeHTTP(w, r)
}