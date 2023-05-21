package main

import (
	"context"
	"io"
	"log"

	pb "grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CaculatorServiceClient) {
	log.Println("doPrime was invoked")

	stream, err := c.Primes(context.Background(), &pb.PrimesRequest{
		Number: 120,
	})
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", msg.Prime)
	}

}
