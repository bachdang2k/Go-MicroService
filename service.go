package main

import (
	"context"
	"fmt"
)

type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceService struct{}

// FetchPrice implements PriceFetcher.
func (*priceService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

func (s *priceService) PriceFetcher(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 200_000.0,
	"ETH": 200.0,
	"GG":  100_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given tiker (%s) is not supported", ticker)
	}

	return price, nil
}
