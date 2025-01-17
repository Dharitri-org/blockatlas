package stellar

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/Dharitri-org/tw-go-libs/client"
)

type Client struct {
	client.Request
}

func (c *Client) GetTxsOfAddress(address string) ([]Payment, error) {
	query := url.Values{
		"order": {"desc"},
		"limit": {"25"},
		"join":  {"transactions"},
	}
	path := fmt.Sprintf("accounts/%s/payments", url.PathEscape(address))

	var payments PaymentsPage
	err := c.Get(&payments, path, query)
	if err != nil {
		return nil, err
	}
	return payments.Embedded.Records, nil
}

func (c *Client) CurrentBlockNumber() (int64, error) {
	query := url.Values{
		"order": {"desc"},
		"limit": {"1"},
	}
	var ledgers LedgersPage
	err := c.Get(&ledgers, "ledgers", query)
	if err != nil {
		return 0, nil
	}

	if len(ledgers.Embedded.Records) == 0 {
		return 0, errors.New("CurrentBlockNumber: Records is empty")
	}
	return ledgers.Embedded.Records[0].Sequence, nil
}

func (c *Client) GetBlockByNumber(num int64) (*Block, error) {
	ledger, err := c.getLedger(num)
	if err != nil {
		return nil, err
	}

	query := url.Values{
		"order": {"desc"},
		"limit": {"100"},
		"join":  {"transactions"},
	}
	path := fmt.Sprintf("ledgers/%d/payments", num)

	var payments PaymentsPage
	err = c.Get(&payments, path, query)
	if err != nil {
		return nil, err
	}
	return &Block{Ledger: *ledger, Payments: payments.Embedded.Records}, nil
}

func (c *Client) getLedger(num int64) (ledger *Ledger, err error) {
	path := fmt.Sprintf("ledgers/%d", num)
	err = c.Get(&ledger, path, nil)
	return
}
