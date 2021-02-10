package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/S-ign/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedAddServiceServer
	calculatorpb.UnimplementedPrimeNumberDecompositionServiceServer
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
	stream calculatorpb.PrimeNumberDecompositionService_PrimeNumberDecompositionServer) error {

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

func main() {
	fmt.Println("CalculatorService initialized...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterAddServiceServer(s, &server{})
	calculatorpb.RegisterPrimeNumberDecompositionServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
