package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/S-ign/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("connecting to RPC server...")
	cc, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	}
	defer cc.Close()

	//c := calculatorpb.NewAddServiceClient(cc)
	//doUnary(c)

	c := calculatorpb.NewPrimeNumberDecompositionServiceClient(cc)
	doServerStreaming(c)
}

func doUnary(c calculatorpb.AddServiceClient) {
	fmt.Println("Starting Unary RPC...")
	req := &calculatorpb.AddRequest{
		Adding: &calculatorpb.Adding{
			FirstNumber:  3379846737546378550,
			SecondNumber: 9007489396297220334,
		},
	}

	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Add RPC %v\n", err)
	}
	log.Printf("Response from Add: %v\n", res.Sum)
}

func doServerStreaming(c calculatorpb.PrimeNumberDecompositionServiceClient) {
	fmt.Println("Starting Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		PrimeNumberDecompositionData: &calculatorpb.PrimeNumberDecompositionData{
			CData: 2983460236533357648,
		},
	}
	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition %v\n", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v\n", err)
		}
		log.Printf("Response from PrimeNumberDecomposition: %v\n", msg.GetSData())
	}
}
