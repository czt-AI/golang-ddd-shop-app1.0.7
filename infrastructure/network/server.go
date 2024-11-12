package network

import (
	"net/http"
	"net/http/httputil"
	"time"

	"shop-app/config/network"
)

func StartServer(config *network.Config) error {
	server := &http.Server{
		Addr:    config.HTTPServer.Addr,
		Timeout: config.HTTPServer.Timeout,
	}

	// 配置HTTPS
	if config.HTTPServer.Addr[0] == 's' { // 检查是否为HTTPS
		tlsConfig := &tls.Config{
			// 配置TLS
		}
		server.TLSConfig = tlsConfig
	}

	// 配置负载均衡器
	if config.LoadBalancer.Addr != "" {
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{Host: config.LoadBalancer.Addr})
		server.Handler = proxy
	}

	// 启动服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed to start: %v", err)
		}
	}()

	return nil
}