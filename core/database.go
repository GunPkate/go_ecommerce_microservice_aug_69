package core

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase() (*Database, error) {
	db, err := sqlx.Open(
		"mysql",
		"GP:GP@tcp(localhost:3406)/go_ecommerce_aug69?parseTime=true",
	)
	if err != nil {
		return nil, fmt.Errorf("error database: %v", err)
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetDB() *sqlx.DB {
	return d.db
}
