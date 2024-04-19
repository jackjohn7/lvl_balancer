package lib

import (
	"errors"
	"log"
	"net/http"
)

type LoadBalancer struct {
	port       string
	servers    []Server
	robinCount int
	config     *LoadBalancerConfig
}

// Creates a LoadBalancer with default configurations
func NewLoadBalancer(servers []Server, balancerPort string) (balancer *LoadBalancer, err error) {
	if len(servers) == 0 {
		return nil, errors.New("You must provide at least one server")
	}
	balancer = &LoadBalancer{
		port:       balancerPort,
		servers:    servers,
		robinCount: 0,
		config: &LoadBalancerConfig{
			Logging:    false,
			TargetFunc: StandardRRTargeting,
		},
	}
	return balancer, nil
}

func NewLoadBalancerWithConfig(
	servers []Server,
	balancerPort string,
	config *LoadBalancerConfig,
) (balancer *LoadBalancer, err error) {
	if len(servers) == 0 {
		return nil, errors.New("You must provide at least one server")
	}
	balancer = &LoadBalancer{
		port:       balancerPort,
		servers:    servers,
		robinCount: 0,
		config:     config,
	}
	return balancer, nil
}

// Targets the next server in the round-robin cycle. This allows
// for a relatively simple but effective form of load balancing.
//
// Say there are three servers, A, B, and C. The request cycle
// would look like the following:
//
// Request 1 -> A
//
// Request 2 -> B
//
// Request 3 -> C
//
// Request 4 -> A ...
func StandardRRTargeting(bal *LoadBalancer) Server {
	server := bal.servers[bal.robinCount%len(bal.servers)]
	// if the first server found isn't alive, move to the next
	for !server.IsAlive() {
		bal.robinCount++
		server = bal.servers[bal.robinCount%len(bal.servers)]
	}
	bal.robinCount++

	return server
}

func (bal *LoadBalancer) getTargetServer() Server {
	server := bal.config.TargetFunc(bal)
	return server
}

func (balancer *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	srvr := balancer.getTargetServer()
	if balancer.config.Logging {
		log.Printf("Proxying request to %s\n", srvr.GetAddress())
	}

	srvr.Serve(rw, req)
}

func (balancer *LoadBalancer) Balance(path string, mux *http.ServeMux) {
	destination := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// define context (planned?)
		balancer.serveProxy(rw, req)
	})
	mux.HandleFunc(path, destination)
}
