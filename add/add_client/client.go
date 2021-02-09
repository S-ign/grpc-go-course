package main

import (
	"context"
	"fmt"
	"log"

	"github.com/S-ign/grpc-go-course/add/addpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("connecting to RPC server...")
	cc, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	}
	defer cc.Close()

	c := addpb.NewAddServiceClient(cc)

	doUnary(c)
}

func doUnary(c addpb.AddServiceClient) {
	fmt.Println("Starting Unary RPC...")
	req := &addpb.AddRequest{
		Adding: &addpb.Adding{
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
