package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/S-ign/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	//doClientStreaming(c)
	//doBiDiStreaming(c)
	doErrorUnary(c)
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

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	numbers := []int64{1, 5, 3, 6, 2, 20}

	requests := []*calculatorpb.FindMaximumRequest{}

	for _, v := range numbers {
		requests = append(requests, &calculatorpb.FindMaximumRequest{
			NextNumber: &calculatorpb.NextNumber{
				Number: v,
			},
		})
	}

	// send all reqests to server
	go func() {
		for _, req := range requests {
			fmt.Printf("sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		err = stream.CloseSend()
		if err != nil {
			log.Fatalf("error while closing send stream: %v\n", err)
		}
	}()

	// recv all responses from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error recieving msg from server: %v\n", err)
			}
			fmt.Printf("Server Response: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")

	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -2)
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user)
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
			}
		} else {
			log.Fatalf("big error calling SquareRoot: %v", err)
		}
	}
	fmt.Printf("Result of square root of %v: %v\n", n, res.GetNumberRoot())
}
