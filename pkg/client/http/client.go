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

//NewClient create a new http client for the sigstat server
func NewClient() *Client {
	var u url.URL
	u.Scheme = "http"
	u.Host = "localhost:9000"

	c := &Client{
		Url: u,
	}
	c.commandService.client = c
	return c
}

//CommandService represents the CommandService interface implemented by the http client
func (c *Client) CommandService() sigstat.CommandService { return &c.commandService }
