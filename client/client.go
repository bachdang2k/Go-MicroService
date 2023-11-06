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

func New(endPoint string) *Client {
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

	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("service responsed with non OK msg error: %s", httpErr["error"])
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}

	return priceResp, nil
}



