package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// docs: NewMySQLStorage creates a new MySQL connection
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN());
	if err != nil {
		log.Fatal(err);
		return nil, err;
	}

	return db, nil;
}