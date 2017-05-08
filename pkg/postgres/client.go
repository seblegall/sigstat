package postgres

import (
	"fmt"

	"github.com/seblegall/sigstat/pkg/sigstat"
)

// CAUTION
// All this code is made to be temp.

type Client struct {
	commandService CommandService
}

func NewClient() *Client {
	c := &Client{}
	return c
}

var _ sigstat.CommandService = &CommandService{}

type CommandService struct{}

func (s *CommandService) UpdateStatus(cmd sigstat.Command) {
	if cmd.Status == "running" {
		fmt.Println("Process is runing: ", cmd.Status)
	} else {
		fmt.Println("Process is dead: ", cmd.Status)
	}
}

func (c *Client) CommandService() sigstat.CommandService { return &c.commandService }
