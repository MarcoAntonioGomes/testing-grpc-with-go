package main

import (
	"log"
	"net"
	"testing-grpc-with-go/pb"
	"testing-grpc-with-go/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Could not serve: %v", err)
	}
}
