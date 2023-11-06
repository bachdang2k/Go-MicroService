package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nasaki/micro/types"
)

type Client struct {
	endPoint string
}

func NewClient(endPoint string) *Client {
	return &Client{
		endPoint: endPoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endPoint, ticker)
	
	req, err := http.NewRequest("get", endpoint, nil)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}

	return priceResp, nil
}



