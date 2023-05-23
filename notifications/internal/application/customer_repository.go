package application

import (
	"context"

	"github.com/fd1az/mallbots/notifications/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
