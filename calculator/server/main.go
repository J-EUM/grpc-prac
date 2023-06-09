package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-go-course/calculator/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CaculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCaculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
