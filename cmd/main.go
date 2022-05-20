package main

import (
	"fmt"
	"net"

	"github.com/truepay/go-boilerplate/database"
	"github.com/truepay/go-boilerplate/modules/banking/application"
	"github.com/truepay/go-boilerplate/modules/banking/infra"
	"github.com/truepay/go-boilerplate/modules/banking/interface/grpc/handlers"
	pb "github.com/truepay/go-boilerplate/modules/banking/interface/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	// Create database instance
	db, err := database.New()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	pb.RegisterAccountServicesServer(server, &handlers.AccountHandler{
		CreateAccount: application.MakeCreateAccount(infra.NewPostgresAccountRepository(db)),
	})
	fmt.Println("Server started on port 8080")
	if err := server.Serve(listener); err != nil {
		fmt.Errorf("failed to serve: %v", err)
	}
}
