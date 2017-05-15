package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/seblegall/sigstat/pkg/sigstat"
)

//Client is a sigstat.Client. Here, it create a postgres connection.
type Client struct {
	usr            string
	psw            string
	dbName         string
	db             *sql.DB
	commandService CommandService
}

//NewClient create a new postgres client.
func NewClient(usr, psw, dbName string) *Client {
	c := &Client{
		usr:    usr,
		psw:    psw,
		dbName: dbName,
	}

	c.commandService.client = c

	return c
}

//Open opens a postgres connection.
func (c *Client) Open() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		c.usr, c.psw, c.dbName)

	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		log.Fatal(err)
	}
	c.db = db
}

//Close close a postgres connection
func (c *Client) Close() {
	err := c.db.Close()

	if err != nil {
		log.Fatal(err)
	}
}

//CommandService represents the CommandService interface implemented by the http client
func (c *Client) CommandService() sigstat.CommandService { return &c.commandService }
