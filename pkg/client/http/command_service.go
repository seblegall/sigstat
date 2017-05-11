package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (s *CommandService) CreateCommand(cmd sigstat.Command) {
	var u url.URL
	u.Scheme = "http"
	u.Host = "localhost:9000"
	u.Path = "/cmd/"

	log.Println(u.String())

	d, err := json.Marshal(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println(d)

	// Create request.
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(d))

	// Execute request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
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
