package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/fd1az/mallbots/ordering/orderingpb"
	"github.com/fd1az/mallbots/payments/internal/application"
)

type OrderRepository struct {
	client orderingpb.OrderingServiceClient
}

var _ application.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(conn *grpc.ClientConn) OrderRepository {
	return OrderRepository{
		client: orderingpb.NewOrderingServiceClient(conn),
	}
}

func (r OrderRepository) Complete(ctx context.Context, invoiceID, orderID string) error {
	_, err := r.client.CompleteOrder(ctx, &orderingpb.CompleteOrderRequest{
		Id:        orderID,
		InvoiceId: invoiceID,
	})
	return err
}
