package icon

import (
	"net/url"
	"strconv"

	"github.com/Dharitri-org/tw-go-libs/client"
	"github.com/Dharitri-org/tw-go-libs/types"
)

type Client struct {
	client.Request
}

func (c *Client) GetAddressTransactions(address string) ([]Tx, error) {
	query := url.Values{
		"address": {address},
		"count":   {strconv.Itoa(types.TxPerPage)},
	}
	var res Response
	err := c.Get(&res, "address/txList", query)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}
