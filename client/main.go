package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	pb "github.com/illenko/go-grpc-common"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call Pay method
	payReq := &pb.PaymentRequest{
		OrderId: "1",
		UserId:  "1",
		Amount:  100,
	}
	payRes, err := c.Pay(ctx, payReq)
	if err != nil {
		log.Fatalf("could not pay: %v", err)
	}
	log.Printf("Payment Response: %v", payRes)

	// Call GetPayment method
	getPayReq := &pb.GetPaymentRequest{
		PaymentId: "1",
	}
	getPayRes, err := c.GetPayment(ctx, getPayReq)
	if err != nil {
		log.Fatalf("could not get payment: %v", err)
	}
	log.Printf("Get Payment Response: %v", getPayRes)
}
