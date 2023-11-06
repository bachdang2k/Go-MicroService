package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/nasaki/micro/client"
	"github.com/nasaki/micro/proto"
)

func main() {

	var (
		jsonAddr = flag.String("json", ":3000", "Listen adress of the json transport")
		grpcAddr = flag.String("grpc", ":4000", "Listen adress of the grpc transport")
		svc = NewLogginService(&priceService{})
		ctx = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			time.Sleep(2 * time.Second)
			resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", resp)
		}
	}()

	go MakeGRPCServerAndRun(*grpcAddr, svc)

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()

}


