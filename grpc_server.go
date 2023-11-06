package main

import (
	"context"

	"github.com/nasaki/micro/proto"
)

type GRPCPriceFetcherServer struct {
	svc PriceService
	//proto.UnimplementedPriceFetcherServer 
}

func NewGRPCPriceFetcherServer(svc PriceService) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		svc: svc,
	}
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) {

}