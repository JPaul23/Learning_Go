package main

import (
	"context"
	pb "helloWorld/hello"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address        = "localhost:50051"
	status         = "Success"
	order_id       = "aX123"
	amount         = 12.5
	payment_method = "stripe"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Amount: amount, OrderId: order_id, PaymentMethod: payment_method})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Status: %v", r.Success)
	log.Printf("Message: %s", r.Message)
}
