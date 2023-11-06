package main

import (
	"flag"

)

func main() {

	// client := client.New("http://localhost:3000")

	// price, err := client.FetchPrice(context.Background(), "ET")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", price)
	// svc := NewLogginService(&priceFetcher{})
	//return
	svc := NewLogginService(NewMetricService(&priceFetcher{}))

	listenAddr := flag.String("listenAddr", ":3000", "listen address the service is running")
	flag.Parse()

	server := NewJSONAPIServer(*listenAddr, svc)

	server.Run()
}


