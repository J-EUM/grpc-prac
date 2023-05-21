package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	// 이 함수가 불러졌다
	log.Println("doGreet was invoked")

	// rpc엔드포인트를 콜하고 이 rpc엔드포인트로부터 그릿리스폰스를 받거나 에러를 받거나.
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "J",
	})

	//에러를 받으면 뭔문젠지 체크하고
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	// 에러가 아니면 결과를 프린트할거다
	log.Printf("Greeting: %s\n", res.Result)
}
