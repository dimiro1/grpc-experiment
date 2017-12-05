package main

import (
	"context"
	"log"

	"github.com/dimiro1/grpc-experiment"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := service.NewCalculatorClient(conn)

	r, err := c.Sum(context.Background(), &service.SumRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	log.Printf("Sum: %d", r.Sum)
}
