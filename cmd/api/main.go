package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/bernardosecades/test/internal/controller/grpc/purchase"
	"github.com/bernardosecades/test/internal/purchase/repository/mysql"
	"github.com/bernardosecades/test/internal/purchase/service"
	"github.com/bernardosecades/test/proto/go/protobuf/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main()  {
	// Repositories NewMySQLPurchaseRepository
	purchaseRepository := mysql.NewMySQLPurchaseRepository(
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	// Services
	purchaseService := service.NewPurchaseService(&purchaseRepository)

	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &purchase.Controller{PurchaseService: purchaseService}

	// Register gRPC server
	api.RegisterPurchaseServiceServer(s, srv)

	grpcPort := os.Getenv("GRPC_PORT")
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("listen in port " + grpcPort)

	// Enable reflection (Evans: `evans -r -p 8081`)
	reflection.Register(s)

	// Start gRPC server
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
