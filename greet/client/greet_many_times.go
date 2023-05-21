package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	// rpc엔드포인트 호출.
	// rpc엔드포인트에서 stream or error를 받음.
	// 에러면 에러출력
	// 에러아니면 stream.Recv를 호출하는 무한루프
	// stream.Recv로부터 EOF가 있는지 확인, 있으면 루프 빠져나옴,
	// 에러없으면 결과를 출력

	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "J",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		// 통신끝났는지 확인
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
