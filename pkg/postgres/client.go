package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/seblegall/sigstat/pkg/sigstat"
)

// CAUTION
// All this code is made to be temp.

type Client struct {
	usr            string
	psw            string
	dbName         string
	db             *sql.DB
	commandService CommandService
}

func NewClient(usr, psw, dbName string) *Client {
	c := &Client{
		usr:    usr,
		psw:    psw,
		dbName: dbName,
	}

	c.commandService.client = c

	return c
}

func (c *Client) Open() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		c.usr, c.psw, c.dbName)

	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		log.Fatal(err)
	}
	c.db = db
}

func (c *Client) Close() {
	err := c.db.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func (c *Client) CommandService() sigstat.CommandService { return &c.commandService }
