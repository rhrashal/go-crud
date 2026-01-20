package controllers

import (
	"context"

	"github.com/rhrashal/go-crud/initializers" // Adjust to your module name
	"github.com/rhrashal/go-crud/models"
	pb "github.com/rhrashal/go-crud/proto" // Matches go_package in proto
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type productServer struct {
	pb.UnimplementedProductServiceServer
	db *gorm.DB // Use the global DB from initializers
}

func NewProductServer() *productServer {
	return &productServer{db: initializers.DB}
}

// CreateProduct
func (s *productServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := req.GetProduct()
	if product.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}

	dbProduct := models.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	result := s.db.Create(&dbProduct)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, "Failed to create product")
	}

	return &pb.CreateProductResponse{Id: uint32(dbProduct.ID)}, nil
}

// GetProduct
func (s *productServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var dbProduct models.Product
	result := s.db.First(&dbProduct, req.GetId())
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Product not found")
		}
		return nil, status.Error(codes.Internal, "Failed to get product")
	}

	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:          uint32(dbProduct.ID),
			Name:        dbProduct.Name,
			Description: dbProduct.Description,
			Price:       dbProduct.Price,
		},
	}, nil
}

// UpdateProduct
func (s *productServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	product := req.GetProduct()
	var dbProduct models.Product
	result := s.db.First(&dbProduct, product.GetId())
	if result.Error != nil {
		return nil, status.Error(codes.NotFound, "Product not found")
	}

	dbProduct.Name = product.Name
	dbProduct.Description = product.Description
	dbProduct.Price = product.Price

	s.db.Save(&dbProduct)

	return &pb.UpdateProductResponse{
		Product: &pb.Product{
			Id:          uint32(dbProduct.ID),
			Name:        dbProduct.Name,
			Description: dbProduct.Description,
			Price:       dbProduct.Price,
		},
	}, nil
}

// DeleteProduct
func (s *productServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	result := s.db.Delete(&models.Product{}, req.GetId())
	if result.Error != nil {
		return nil, status.Error(codes.Internal, "Failed to delete product")
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "Product not found")
	}
	return &pb.DeleteProductResponse{Success: true}, nil
}

// ListProducts
func (s *productServer) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	var dbProducts []models.Product
	s.db.Find(&dbProducts)

	products := make([]*pb.Product, len(dbProducts))
	for i, dbProduct := range dbProducts {
		products[i] = &pb.Product{
			Id:          uint32(dbProduct.ID),
			Name:        dbProduct.Name,
			Description: dbProduct.Description,
			Price:       dbProduct.Price,
		}
	}

	return &pb.ListProductsResponse{Products: products}, nil
}
