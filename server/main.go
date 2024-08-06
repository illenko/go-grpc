package main

import (
	"context"
	"log"
	"net"

	pb "github.com/illenko/go-grpc-common"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) Pay(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Received Pay request: OrderId=%s, UserId=%s, Amount=%d", req.OrderId, req.UserId, req.Amount)
	response := &pb.PaymentResponse{
		PaymentId: "1",
		OrderId:   req.OrderId,
		UserId:    req.UserId,
		Amount:    req.Amount,
		Timestamp: timestamppb.Now(),
	}
	log.Printf("Sending Pay response: %v", response)
	return response, nil
}

func (s *server) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Received GetPayment request: PaymentId=%s", req.PaymentId)
	response := &pb.PaymentResponse{
		PaymentId: req.PaymentId,
		OrderId:   "1",
		UserId:    "1",
		Amount:    100,
		Timestamp: timestamppb.Now(),
	}
	log.Printf("Sending GetPayment response: %v", response)
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
