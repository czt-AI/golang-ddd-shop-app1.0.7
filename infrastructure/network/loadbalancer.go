package network

import (
	"net/http"
	"net/http/httputil"
	"sync"
)

type LoadBalancer struct {
	Backends []*url.URL
	sync.Mutex
}

func NewLoadBalancer(backends ...*url.URL) *LoadBalancer {
	return &LoadBalancer{
		Backends: backends,
	}
}

func (lb *LoadBalancer) ChooseBackend() *url.URL {
	lb.Lock()
	defer lb.Unlock()

	if len(lb.Backends) == 0 {
		return nil
	}

	return lb.Backends[0] // Simple round-robin for example
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.ChooseBackend()
	if backend == nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(backend)
	proxy.ServeHTTP(w, r)
}