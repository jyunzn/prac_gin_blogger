package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)


var (
	DB *sqlx.DB
)

func Init(dsn string) (err error) {
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)


	return
}
