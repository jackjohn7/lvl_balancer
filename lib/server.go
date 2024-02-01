package lib

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	GetAddress() string

	IsAlive() bool

	Serve(rw http.ResponseWriter, req *http.Request)
}

// MinimalServer implements the `Server` interface such that
// it can be used within the LoadBalancer. You can create one
// of these with the `NewMinimalServer` function.
type MinimalServer struct {
	url   *url.URL
	addr  string
	proxy *httputil.ReverseProxy
}

func (m *MinimalServer) GetAddress() string {
	return m.addr
}

func (m *MinimalServer) GetUrl() *url.URL {
	return m.url
}

func (m *MinimalServer) IsAlive() bool {
	return true
}

func (m *MinimalServer) Serve(rw http.ResponseWriter, req *http.Request) {
	req.Host = m.GetUrl().Host
	m.proxy.ServeHTTP(rw, req)
}

func NewMinimalServer(addr string) (ms *MinimalServer, err error) {
	url, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	log.Println("Registered: ", url)

	ms = &MinimalServer{
		url:   url,
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(url),
	}

	return ms, err
}
