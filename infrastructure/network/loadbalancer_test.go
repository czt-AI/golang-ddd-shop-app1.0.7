package network

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoadBalancerServeHTTP(t *testing.T) {
	// 创建模拟的后端服务
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Backend Response"))
	}))

	// 解析后端服务地址
	testServerURL, _ := url.Parse(testServer.URL)

	// 创建负载均衡器，添加后端服务
	lb := NewLoadBalancer(testServerURL)

	// 创建请求并转发到负载均衡器
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	lb.ServeHTTP(response, request)

	// 验证响应
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}

	// 验证响应内容
	expectedBody := "Backend Response"
	if response.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, response.Body.String())
	}
}