package main

import (
	"flag"
)

func main() {
	// svc := NewLogginService(&priceFetcher{})
	svc := NewLogginService(NewMetricService(&priceFetcher{}))

	listenAddr := flag.String("listenAddr", ":3000", "listen address the service is running")
	flag.Parse()

	server := NewJSONAPIServer(*listenAddr, svc)

	server.Run()
}


