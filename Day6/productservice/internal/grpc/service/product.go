package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"product/internal/service"
	"product/pb/productservice/proto/productpb"
)

type ProductService struct {
	productpb.UnimplementedProductServiceServer
	s service.ProductService
}

func NewProductService(svc service.ProductService) *ProductService {
	return &ProductService{
		s: svc,
	}
}

func (g *ProductService) GetProductByCategoryId(ctx context.Context, req *productpb.CategoryRequest) (*productpb.ProductListResponse, error) {
	// Validate request
	if req.GetCategoryId() == "" {
		return nil, status.Error(codes.InvalidArgument, "category_id is required")
	}

	// Parse UUID
	categoryUUID, err := uuid.Parse(req.GetCategoryId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid category_id format")
	}

	// Get products from service
	products, err := g.s.GetProductByCategoryId(ctx, categoryUUID)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to get products: %v", err))
	}

	// Convert to protobuf response
	pbProducts := make([]*productpb.ProductResponse, 0, len(products))
	for _, product := range products {
		pbProducts = append(pbProducts, &productpb.ProductResponse{
			Id:    product.Id.String(),
			Name:  product.Name,
			Price: float64(product.Price),
		})
	}

	return &productpb.ProductListResponse{
		Products: pbProducts,
	}, nil
}
