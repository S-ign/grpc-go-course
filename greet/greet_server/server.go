package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/S-ign/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()
	result := splitBy(" ", "Hello", firstname, lastname)
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest,
	stream greetpb.GreetService_GreetManyTimesServer) error {

	fmt.Printf("Greet function was invoked with %v\n", req)
	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()

	result := splitBy(" ", "Hello", firstname, lastname)
	for i := 0; i < 10; i++ {
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

// LongGreet Handles server streaming
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet function was invoked with a streaming request")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("error reading client stream: %v\n", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		result += splitBy(" ", "Hello", firstName, lastName, "!")
	}
}

func splitBy(s string, words ...string) string {
	return strings.Join(words, s)
}

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
