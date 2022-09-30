package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type Balance struct {
	Id   int `json:"id"`
	Cash int `json:"cash"`
}

var db *sql.DB

func Initialize() (err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	return err
}

func Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			return
		}
	}
}

func GetBalance(id int) (cash int, err error) {
	err = Initialize()
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	defer Close()
	p := Balance{Id: id}

	rows := db.QueryRow("SELECT cash FROM quote WHERE id=$1", id)
	err = rows.Scan(&p.Cash)
	return p.Cash, err

}

func Check(id int) (check bool, err error) {
	Initialize()
	defer Close()
	rows := db.QueryRow("SELECT EXISTS(SELECT 1 FROM quote WHERE id=$1)", id)
	err = rows.Scan(&check)
	if err != nil {
		return
	}
	return
}

func InsertBalance(id int, cash int) (err error) {
	Initialize()
	defer Close()
	_, err = db.Exec("INSERT INTO quote (id, cash) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET cash = excluded.cash;", id, cash)
	return
}
