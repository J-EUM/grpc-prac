package main

import (
	"context"
	"log"

	pb "grpc-go-course/calculator/proto"
)

func doSum(c pb.CaculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Result: %d\n", res.Result)
}
