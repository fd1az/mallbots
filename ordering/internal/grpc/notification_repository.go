package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/fd1az/mallbots/notifications/notificationspb"
	"github.com/fd1az/mallbots/ordering/internal/domain"
)

type NotificationRepository struct {
	client notificationspb.NotificationsServiceClient
}

var _ domain.NotificationRepository = (*NotificationRepository)(nil)

func NewNotificationRepository(conn *grpc.ClientConn) NotificationRepository {
	return NotificationRepository{client: notificationspb.NewNotificationsServiceClient(conn)}
}

func (r NotificationRepository) NotifyOrderCreated(ctx context.Context, orderID, customerID string) error {
	_, err := r.client.NotifyOrderCreated(ctx, &notificationspb.NotifyOrderCreatedRequest{
		OrderId:    orderID,
		CustomerId: customerID,
	})
	return err
}

func (r NotificationRepository) NotifyOrderCanceled(ctx context.Context, orderID, customerID string) error {
	_, err := r.client.NotifyOrderCanceled(ctx, &notificationspb.NotifyOrderCanceledRequest{
		OrderId:    orderID,
		CustomerId: customerID,
	})
	return err
}

func (r NotificationRepository) NotifyOrderReady(ctx context.Context, orderID, customerID string) error {
	_, err := r.client.NotifyOrderReady(ctx, &notificationspb.NotifyOrderReadyRequest{
		OrderId:    orderID,
		CustomerId: customerID,
	})
	return err
}
