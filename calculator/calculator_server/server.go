package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/S-ign/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	fmt.Printf("Add function was invoked with %v\n", req)
	firstnumber := req.GetAdding().GetFirstNumber()
	secondnumber := req.GetAdding().GetSecondNumber()
	sum := firstnumber + secondnumber
	res := &calculatorpb.AddResponse{
		Sum: sum,
	}
	return res, nil
}

// PrimeNumberDecomposition takes in int64 and streams its lowest divisible prime until no longer divisible.
func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest,
	stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {

	fmt.Printf("PrimeNumberDecomposition function was invoked with %v\n", req)
	cData := req.GetPrimeNumberDecompositionData().GetCData()

	var lowestPrime int64 = 2
	for {
		if cData%lowestPrime == 0 {
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				SData: lowestPrime,
			}
			stream.Send(res)
			cData /= lowestPrime
			time.Sleep(1000 * time.Millisecond)
		} else {
			lowestPrime++
			if cData <= 1 {
				break
			}
		}
	}

	return nil
}

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	fmt.Printf("Average function was invoked with a streaming request...")
	var data float64
	var counter float64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Data: data / counter,
			})
		}
		if err != nil {
			log.Fatalf("error reading client stream: %v\n", err)
		}
		number := float64(req.GetAveraging().GetNumber())
		data += number
		counter++
		fmt.Printf("data: %f\ncounter: %f\n", data, counter)
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Printf("FindMaximum function was invoked with a streaming request...")

	var maxNumber int64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream %v\n", err)
			return err
		}
		nextNumber := req.GetNextNumber().GetNumber()
		if nextNumber > maxNumber {
			maxNumber = nextNumber
			err = stream.Send(&calculatorpb.FindMaximumResponse{
				Result: maxNumber,
			})
			if err != nil {
				log.Fatalf("error sending response to client: %v\n", err)
				return err
			}
		}
	}
}

func main() {
	fmt.Println("CalculatorService initialized...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
