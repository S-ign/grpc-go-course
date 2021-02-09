package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/S-ign/grpc-go-course/add/addpb"
	"google.golang.org/grpc"
)

type server struct {
	addpb.UnimplementedAddServiceServer
}

func (*server) Add(ctx context.Context, req *addpb.AddRequest) (*addpb.AddResponse, error) {
	fmt.Printf("Add function was invoked with %v\n", req)
	firstnumber := req.GetAdding().GetFirstNumber()
	secondnumber := req.GetAdding().GetSecondNumber()
	sum := firstnumber + secondnumber
	res := &addpb.AddResponse{
		Sum: sum,
	}
	return res, nil
}

func main() {
	fmt.Println("I Add {first_number} and {second_number} together!")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	addpb.RegisterAddServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
