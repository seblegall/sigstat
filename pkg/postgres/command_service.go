package postgres

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/seblegall/sigstat/pkg/sigstat"
)

var _ sigstat.CommandService = &CommandService{}

type CommandService struct {
	client *Client
}

func (s *CommandService) UpdateStatus(cmd sigstat.Command) {
	// s.client.db.Close()
	// defer s.client.db.Close()
	//
	// err := db.QueryRow("INSERT INTO monitor(datetime, app, route, verb, calls, response200, response400, response500) VALUES($1,$2,$3,$4,$5,$6,$7,$8) returning route;",
	// 	pq.FormatTimestamp(from), value[0], value[2], value[1], value[3], value[5], value[7], value[9]).Scan(&route)
	//
	// if err != nil {
	// 	return err
	// }

}

func (s *CommandService) CreateCommand(cmd sigstat.Command) {
	s.client.Open()
	defer s.client.Close()

	var id int

	err := s.client.db.QueryRow("INSERT INTO command_group(name) VALUES($1) returning id;",
		cmd.Group).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("group id is : ", id)
}
