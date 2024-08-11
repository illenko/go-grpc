package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/illenko/go-grpc-common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func runGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		return err
	}

	log.Println("Serving gRPC-Gateway on http://localhost:8080")
	return http.ListenAndServe(":8080", mux)
}

func main() {
	if err := runGateway(); err != nil {
		log.Fatalf("failed to run gRPC-Gateway: %v", err)
	}
}
