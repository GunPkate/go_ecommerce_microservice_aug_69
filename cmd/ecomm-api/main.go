package main

import (
	"go_ecommerce_aug69/core"
	"go_ecommerce_aug69/ecomm-api/server"
	"go_ecommerce_aug69/ecomm-api/storer"
	"go_ecommerce_aug69/handler" // Import your handler package
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := core.NewDatabase()
	if err != nil {
		log.Fatalf("error database: %v", err)
	}
	defer db.Close()
	log.Println("successfully connected to database")

	st := storer.NewSQLStorer(db.GetDB())

	srvInstance := server.NewServer(st)

	hdl := handler.NewHandler(srvInstance)
	router := handler.RegisterRoutes(hdl)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server is running on port 8080...")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen error: %s\n", err)
	}
}
