package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/rhrashal/go-crud/controllers" // Adjust to your module name
	"github.com/rhrashal/go-crud/initializers"
	pb "github.com/rhrashal/go-crud/proto"
	"github.com/rhrashal/go-crud/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	// Start Gin HTTP server
	r := gin.Default()
	routes.TodoRoutes(r) // Existing TODO routes

	go func() { // Run Gin in a goroutine
		if err := r.Run(); err != nil {
			log.Fatalf("Failed to run Gin server: %v", err)
		}
	}()

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051") // gRPC on port 50051
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, controllers.NewProductServer())
	reflection.Register(grpcServer)
	fmt.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
