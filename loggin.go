package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceFetcher
}

func NewLogginService(next PriceFetcher) PriceFetcher {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, tikcker string) (price float64, err error) {
	defer func(begin time.Time) {
		log.WithFields(log.Fields{
			"requestID": ctx.Value("requestID"),
			"took": time.Since(begin),
			"err": err,
			"price": price,
		}).Info("fetchPrice")
	} (time.Now()) 

	return s.next.FetchPrice(ctx, tikcker)
}


