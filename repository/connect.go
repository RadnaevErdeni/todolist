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

//var Dbcon *sql.DB

func DBC(c Conf) (*sqlx.DB, error) {
	/*
		Connect.DBC("FantasyQuests", "postgres", "PostgreSQL")
		err := http.ListenAndServe("26.224.38.49:51944", nil)
		log.Fatal(err)
	*/
	//conStr := "user=" + user + " password=" + password + " dbname=" + db + " sslmode=disable"
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

/*
func Close() {
	Dbcon.Close()
}
*/
