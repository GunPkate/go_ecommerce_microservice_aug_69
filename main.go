package main

import (
	"go_ecommerce_aug69/core"
	"log"
)

func main() {
	db, err := core.NewDatabase()
	if err != nil {
		log.Fatalf("error database: %v", err)
	}
	defer db.Close()
	log.Println("successfully connected")
}
