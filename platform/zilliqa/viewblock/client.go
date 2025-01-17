package viewblock

import (
	"fmt"

	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/network/middleware"
)

type Client struct {
	client.Request
}

func InitClient(api, apiKey string) Client {
	c := Client{client.InitClient(api, middleware.SentryErrorHandler)}
	c.Headers["X-APIKEY"] = apiKey
	return c
}

func (c *Client) GetTxsOfAddress(address string) (tx []Tx, err error) {
	path := fmt.Sprintf("addresses/%s/txs", address)
	err = c.Get(&tx, path, nil)
	return
}
