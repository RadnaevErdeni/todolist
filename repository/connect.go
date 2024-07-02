package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "userz"
	todoListsTable  = "todo_lists"
	usersListsTable = "userz_lists"
	todoStrTable    = "todo_str"
	listsStrTable   = "lists_str"
)

type Conf struct {
	Host     string
	Port     string
	Username string
	BDname   string
	Password string
	SSLMode  string
}

func DBC(c Conf) (*sqlx.DB, error) {

	dbcon, err := sqlx.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", c.Username, c.Password, c.BDname, c.SSLMode))
	if err != nil {
		return nil, err
	}
	err = dbcon.Ping()
	if err != nil {
		return nil, err
	}
	return dbcon, nil
}
