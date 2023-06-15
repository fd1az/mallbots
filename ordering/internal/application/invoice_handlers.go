package application

import (
	"context"

	"github.com/fd1az/mallbots/internal/ddd"
	"github.com/fd1az/mallbots/ordering/internal/domain"
)

type InvoiceHandlers struct {
	invoices domain.InvoiceRepository
	ignoreUnimplementedDomainEvents
}

func NewInvoiceHandlers(invoices domain.InvoiceRepository) *InvoiceHandlers {
	return &InvoiceHandlers{
		invoices: invoices,
	}
}

func (h InvoiceHandlers) OnOrderReadied(ctx context.Context, event ddd.Event) error {
	orderReadied := event.(*domain.OrderReadied)
	return h.invoices.Save(ctx, orderReadied.Order.ID, orderReadied.Order.PaymentID, orderReadied.Order.GetTotal())
}
