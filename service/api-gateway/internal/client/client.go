package client

import (
	"context"

	"api-gateway/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type InventoryClient struct {
	productClient  proto.ProductServiceClient
	categoryClient proto.CategoryServiceClient
	conn          *grpc.ClientConn
}

func NewInventoryClient(conn *grpc.ClientConn) *InventoryClient {
	return &InventoryClient{
		productClient:  proto.NewProductServiceClient(conn),
		categoryClient: proto.NewCategoryServiceClient(conn),
		conn:          conn,
	}
}

func (c *InventoryClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *InventoryClient) GetProducts(ctx context.Context, page, limit int32) ([]*proto.Product, error) {
	resp, err := c.productClient.ListProducts(ctx, &proto.ListProductsRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}
	return resp.Products, nil
}

func (c *InventoryClient) GetCategories(ctx context.Context) ([]*proto.Category, error) {
	resp, err := c.categoryClient.ListCategories(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.Categories, nil
} 