package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/gophertrain/material/usefulinterfaces/exercises/inventory/acme"
	"github.com/gophertrain/material/usefulinterfaces/exercises/inventory/postgres"
	"github.com/gophertrain/material/usefulinterfaces/exercises/inventory/transport/http"
	"github.com/gophertrain/material/usefulinterfaces/exercises/inventory/transport/rpc"
)

func main() {

	db, err := sql.Open("postgres",
		"docker:docker@tcp(127.0.0.1:5432)/inventory")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create data storage
	orderStore := postgres.NewOrderService(db)
	productStore := postgres.NewProductService(db)
	supplierStore := postgres.NewSupplierService(db)

	// supplier service
	acmeClient := acme.NewClient("https://acme.com/api")

	// rest http network listener
	hl, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	// rpc network listener
	rl, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	restServer := http.NewRESTService(orderStore, supplierStore, acmeClient, productStore)
	rpcServer := rpc.NewRPCService(orderStore, supplierStore, acmeClient, productStore)

	go func() {
		err := restServer.Serve(hl)
		if err != nil {
			log.Fatalf("Starting rest service: %v", err)
		}
	}()

	go func() {
		err := rpcServer.Serve(rl)
		if err != nil {
			log.Fatalf("Starting rpc service: %v", err)
		}
	}()

	<-make(chan struct{}) // block forever
}
