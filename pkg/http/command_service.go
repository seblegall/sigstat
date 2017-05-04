package http

import (
	"fmt"

	"github.com/seblegall/sigstat/pkg/sigstat"
)

var _ sigstat.CommandService = &CommandService{}

type CommandService struct {
	client *Client
}

func (s *CommandService) UpdateStatus(cmd *sigstat.Command) {
	if cmd.Status == "running" {
		fmt.Println("Process is runing: ", cmd.Status)
	} else {
		fmt.Println("Process is dead: ", cmd.Status)
	}
}
