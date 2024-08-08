package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net"
	"sync"

	pb "github.com/illenko/go-grpc-common"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
	mu             sync.Mutex
	paymentStorage map[string]*pb.PaymentResponse
}

func (s *server) Pay(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Received Pay request: OrderId=%s, UserId=%s, Amount=%d", req.OrderId, req.UserId, req.Amount)
	response := &pb.PaymentResponse{
		PaymentId: uuid.New().String(),
		OrderId:   req.OrderId,
		UserId:    req.UserId,
		Amount:    req.Amount,
		Status:    "init",
		Timestamp: timestamppb.Now(),
	}
	s.mu.Lock()
	s.paymentStorage[response.PaymentId] = response
	s.mu.Unlock()
	log.Printf("Sending Pay response: %v", response)
	return response, nil
}

func (s *server) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.PaymentResponse, error) {
	log.Printf("Received GetPayment request: PaymentId=%s", req.PaymentId)
	s.mu.Lock()
	response, exists := s.paymentStorage[req.PaymentId]
	s.mu.Unlock()
	if !exists {
		return nil, fmt.Errorf("payment not found")
	}
	log.Printf("Sending GetPayment response: %v", response)
	response.Status = "success"
	return response, nil
}

func createGRPCServer() (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	paymentServer := &server{
		paymentStorage: make(map[string]*pb.PaymentResponse),
	}
	pb.RegisterPaymentServiceServer(s, paymentServer)
	return s, lis, nil
}

func main() {
	s, lis, err := createGRPCServer()
	if err != nil {
		log.Fatalf("failed to create gRPC server: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
