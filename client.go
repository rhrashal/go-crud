package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/rhrashal/go-crud/proto" // Adjust path
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	// Example: Create Product
	createReq := &pb.CreateProductRequest{
		Product: &pb.Product{Name: "Laptop", Description: "High-end laptop", Price: 999.99},
	}
	createRes, err := client.CreateProduct(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("Created Product ID: %d\n", createRes.Id)

	// Test other methods similarly...
}
