package main

import (
	"github.com/huseyinbabal/microservices/order/config"
	"github.com/huseyinbabal/microservices/order/internal/adapters/db"
	"github.com/huseyinbabal/microservices/order/internal/adapters/grpc"
	"github.com/huseyinbabal/microservices/order/internal/adapters/payment"
	"github.com/huseyinbabal/microservices/order/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
