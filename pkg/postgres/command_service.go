package postgres

import (
	//Here we only use the database/sql interface
	_ "github.com/lib/pq"
	"github.com/seblegall/sigstat/pkg/sigstat"
)

var _ sigstat.CommandService = &CommandService{}

//CommandService represents the sigstat.CommandService
type CommandService struct {
	client *Client
}

//UpdateStatus update the status of a givent command
func (s *CommandService) UpdateStatus(cmd sigstat.Command) {}

//CreateCommand insert a command in the database.
// This function test if the command group already exist.
//If yes, then it creates the command linked to the matching group
//If no, then it first create a new command group and then, actually create the command.
func (s *CommandService) CreateCommand(cmd sigstat.Command) (int64, error) {
	s.client.Open()
	defer s.client.Close()

	var groupID int64
	var cmdID int64
	//Check if the group exit and create it if not.
	err := s.client.db.QueryRow("INSERT INTO command_group(name) VALUES($1) ON CONFLICT(name) DO UPDATE SET name=EXCLUDED.name RETURNING id;",
		cmd.Group).Scan(&groupID)
	if err != nil {
		return 0, err
	}
	cmd.GroupID = groupID

	err = s.client.db.QueryRow("INSERT INTO command(group_id, cmd, cmd_path) VALUES($1, $2, $3) RETURNING id",
		cmd.GroupID, cmd.Command, cmd.Path).Scan(&cmdID)
	if err != nil {
		return 0, err
	}
	cmd.ID = cmdID

	return cmd.ID, nil
}
