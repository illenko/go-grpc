package main

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/illenko/go-grpc-common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func createGRPCClient(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func makePayment(ctx context.Context, client pb.PaymentServiceClient) (*pb.PaymentResponse, error) {
	payReq := &pb.PaymentRequest{
		OrderId: uuid.New().String(),
		UserId:  uuid.New().String(),
		Amount:  100,
	}
	payRes, err := client.Pay(ctx, payReq)
	if err != nil {
		return nil, err
	}
	return payRes, nil
}

func getPayment(ctx context.Context, client pb.PaymentServiceClient, paymentRequestId string) (*pb.PaymentResponse, error) {
	getPayReq := &pb.GetPaymentRequest{
		PaymentId: paymentRequestId,
	}
	getPayRes, err := client.GetPayment(ctx, getPayReq)
	if err != nil {
		return nil, err
	}
	return getPayRes, nil
}

func main() {
	conn, err := createGRPCClient("localhost:50051")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPaymentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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
