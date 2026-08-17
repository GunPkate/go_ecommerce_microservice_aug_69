package main

import (
	"go_ecommerce_aug69/core"
	"go_ecommerce_aug69/ecomm-api/server"
	"go_ecommerce_aug69/ecomm-api/storer"
	"log"
)

func main() {
	db, err := core.NewDatabase()
	if err != nil {
		log.Fatalf("error database: %v", err)
	}
	defer db.Close()
	log.Println("successfully connected")

	st := storer.NewSQLStorer(db.GetDB())
	_ = server.NewServer(st)
hdl:
}
