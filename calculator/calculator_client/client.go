package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//doUnary(c)
	//doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	// generate request messages
	requests := []*calculatorpb.AverageRequest{}
	for i := 1; i <= 4; i++ {
		requests = append(requests, &calculatorpb.AverageRequest{
			Averaging: &calculatorpb.Averaging{
				Number: int64(i),
			},
		})
	}

	// create stream to send requests to RPC
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error while calling average: %v\n", err)
	}

	// iterate through slice of requests
	// and send them one at a time
	for _, req := range requests {
		fmt.Printf("Sending... %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	// get response by telling the server
	// this is the end of the stream
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error reading msg from server: %v\n", err)
	}
	fmt.Printf("Response from AverageService: %v\n", res)
}
