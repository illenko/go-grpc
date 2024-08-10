package main

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/illenko/go-grpc-common"
	"github.com/segmentio/kafka-go"
)

func makePayment(ctx context.Context, client pb.PaymentServiceClient, writer *kafka.Writer) (*pb.PaymentResponse, error) {
	payReq := &pb.PaymentRequest{
		OrderId: uuid.New().String(),
		UserId:  uuid.New().String(),
		Amount:  100,
	}
	payRes, err := client.Pay(ctx, payReq)
	if err != nil {
		return nil, err
	}

	if err := sendPaymentStatusUpdate(writer, PaymentEvent{
		EventType: "status_update",
		PaymentId: payRes.PaymentId,
		Status:    payRes.Status,
	}); err != nil {
		return nil, err
	}
	return payRes, nil

}

func getPayment(ctx context.Context, client pb.PaymentServiceClient, writer *kafka.Writer, paymentRequestId string) (*pb.PaymentResponse, error) {
	getPayReq := &pb.GetPaymentRequest{
		PaymentId: paymentRequestId,
	}
	res, err := client.GetPayment(ctx, getPayReq)
	if err != nil {
		return nil, err
	}

	if err := sendPaymentStatusUpdate(writer, PaymentEvent{
		EventType: "status_update",
		PaymentId: res.PaymentId,
		Status:    res.Status,
	}); err != nil {
		return nil, err
	}
	return res, nil
}
