package main

import (
	"context"
	"log"
	"time"

	pb "github.com/illenko/go-grpc-common"
)

func main() {

	conn, err := createGRPCClient("localhost:50051")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	paymentRes, err := makePayment(ctx, c)
	if err != nil {
		log.Fatalf("could not pay: %v", err)
	}
	log.Printf("Payment Response: %v", paymentRes)

	getPaymentRes, err := getPayment(ctx, c, paymentRes.PaymentId)
	if err != nil {
		log.Fatalf("could not get payment: %v", err)
	}
	log.Printf("Get Payment Response: %v", getPaymentRes)
}
