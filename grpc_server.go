package main

import (
	"context"
	"math/rand"
	"net"

	"github.com/nasaki/micro/proto"
	"google.golang.org/grpc"
)

func MakeGRPCServerAndRun(listenAddr string, svc PriceService) error {
	grpcPriceFetcher := NewGRPCPriceFetcherServer(svc)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)

	return server.Serve(ln)
}

type GRPCPriceFetcherServer struct {
	svc PriceService
	proto.UnimplementedPriceFetcherServer 
}

func NewGRPCPriceFetcherServer(svc PriceService) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{
		svc: svc,
	}
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	reqID := rand.Intn(100000) 
	ctx = context.WithValue(ctx, "requestID", reqID)

	price, err := s.svc.FetchPrice(ctx, req.Ticker)

	if err != nil {
		return nil, err
	}

	resp := &proto.PriceResponse{
		Ticker: req.Ticker,
		Price: float32(price),
	}

	return resp, nil
}