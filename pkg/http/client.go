package http

import (
	"net/url"

	"github.com/seblegall/sigstat/pkg/sigstat"
)

var _ sigstat.Client = &Client{}

type Client struct {
	Url            url.URL
	commandService CommandService
}

func NewClient() *Client {
	c := &Client{}
	return c
}

func (c *Client) CommandService() sigstat.CommandService { return &c.commandService }
