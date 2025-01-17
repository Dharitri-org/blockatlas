package theta

import (
	"fmt"
	"net/url"

	"github.com/Dharitri-org/tw-go-libs/client"
)

type Client struct {
	client.Request
}

func (c *Client) FetchAddressTransactions(address string) ([]Tx, error) {
	query := url.Values{
		"type":        {"2"},
		"pageNumber":  {"1"},
		"limitNumber": {"100"},
		"isEqualType": {"true"},
	}
	uri := fmt.Sprintf("accounttx/%s", url.PathEscape(address))
	var transfers AccountTxList
	err := c.Get(&transfers, uri, query)
	if err != nil {
		return nil, err
	}
	return transfers.Body, nil
}
