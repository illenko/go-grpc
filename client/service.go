package main

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/illenko/go-grpc-common"
)

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
	res, err := client.GetPayment(ctx, getPayReq)
	if err != nil {
		return nil, err
	}

	return res, nil
}
