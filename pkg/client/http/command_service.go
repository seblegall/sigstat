package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/seblegall/sigstat/pkg/sigstat"
)

var _ sigstat.CommandService = &CommandService{}

//CommandService represents the sigstat.CommandService interface
type CommandService struct {
	client *Client
}

//createCommandResponse is the body of the http response
type createCommandResponse struct {
	GroupID int64  `json:"group-id,omitempty"`
	Err     string `json:"error,omitsempty"`
}

//CreateCommand calls the sigstat server in order to create a new command
func (s *CommandService) CreateCommand(cmd sigstat.Command) (int64, error) {
	var u url.URL
	u = s.client.Url
	u.Path = "/cmd/"

	d, err := json.Marshal(cmd)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// Create request.
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(d))

	// Execute request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var respBody createCommandResponse
	//http error Handling
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return 0, err
	} else if respBody.Err != "" {
		return 0, errors.New(respBody.Err)
	}

	return respBody.GroupID, nil
}

//UpdateStatus send current cmd status to the server
func (s *CommandService) UpdateStatus(cmd sigstat.Command) {
	url := s.client.Url
	url.Path = "/status/"

	// Create request.
	req, err := http.NewRequest("PATCH", url.String(), nil)

	// Execute request.
	resp, err := http.DefaultClient.Do(req)

	log.Println(resp.Status)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
