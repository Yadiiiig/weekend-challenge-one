package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(details string) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", details)
}
