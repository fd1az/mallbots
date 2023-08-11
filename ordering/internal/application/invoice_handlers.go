package application

import (
	"context"

	"github.com/fd1az/mallbots/internal/ddd"
	"github.com/fd1az/mallbots/ordering/internal/domain"
)

type InvoiceHandlers[T ddd.AggregateEvent] struct {
	invoices domain.InvoiceRepository
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*InvoiceHandlers[ddd.AggregateEvent])(nil)

func NewInvoiceHandlers(invoices domain.InvoiceRepository) *InvoiceHandlers[ddd.AggregateEvent] {
	return &InvoiceHandlers[ddd.AggregateEvent]{
		invoices: invoices,
	}
}

func (h InvoiceHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	}
	return nil
}

func (h InvoiceHandlers[T]) onOrderReadied(ctx context.Context, event ddd.AggregateEvent) error {
	orderReadied := event.Payload().(*domain.OrderReadied)
	return h.invoices.Save(ctx, event.AggregateID(), orderReadied.PaymentID, orderReadied.Total)
}
