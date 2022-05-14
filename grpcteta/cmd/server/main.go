package main

import (
	"fmt"
	"net"
	"stepik_golang/grpcteta/internal/handlers"
	"stepik_golang/grpcteta/pkg/orderservice"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	server := &handlers.Server{}
	orderservice.RegisterOrdersServer(s, server)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
