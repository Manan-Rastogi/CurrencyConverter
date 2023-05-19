package main

import (
	"log"
	"net"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"google.golang.org/grpc"
	"github.com/fatih/color"
)

const (
	port = ":8080"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to start server: %v", err.Error())
	}

	grpcServer := grpc.NewServer()                             // grpc server
	
	// Register the gRPC server with the service implementation
	currencyProto.RegisterConverterServer(grpcServer, NewServer())
	color.Blue("GRPC Server started at %v", listener.Addr())

	if err = grpcServer.Serve(listener); err != nil {              
		log.Fatalf("Failed to listen the server: %v", err.Error())
	}
}
