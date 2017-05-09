package http

import (
	"log"
	"net/http"
	"net/url"

	"github.com/seblegall/sigstat/pkg/sigstat"
)

var _ sigstat.CommandService = &CommandService{}

type CommandService struct {
	client *Client
	url    *url.URL
}

func (s *CommandService) UpdateStatus(cmd sigstat.Command) {
	var u url.URL
	u.Scheme = "http"
	u.Host = "localhost:9000"
	u.Path = "/status/"

	log.Println(u.String())

	// Create request.
	req, err := http.NewRequest("PATCH", u.String(), nil)

	// Execute request.
	resp, err := http.DefaultClient.Do(req)

	log.Println(resp.Status)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
