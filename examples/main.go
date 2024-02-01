package main

import (
	"log"
	"net/http"

	. "github.com/jackjohn7/lvl_balancer/lib"
)

func main() {
	bingServer, err := NewMinimalServer("https://www.wikipedia.org")
	if err != nil {
		log.Fatalln("Could not create minimal server wikipedia")
	}
	googleServer, err := NewMinimalServer("https://www.jackbranch.dev")
	if err != nil {
		log.Fatalln("Could not create minimal server jackbranch")
	}
	servers := []Server{
		bingServer,
		googleServer,
	}

	port := ":3005"

	bal, err := NewLoadBalancer(servers, port)
	if err != nil {
		return
	}

	mux := http.NewServeMux()

	bal.Balance("/", mux)

	log.Fatalln(http.ListenAndServe(port, mux))
}
