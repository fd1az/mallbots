package handlers

import (
	"github.com/fd1az/mallbots/internal/ddd"
	"github.com/fd1az/mallbots/ordering/internal/domain"
)

func RegisterInvoiceHandlers(
	invoiceHandlers ddd.EventHandler[ddd.AggregateEvent],
	domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent],
) {
	domainSubscriber.Subscribe(domain.OrderReadiedEvent, invoiceHandlers)
}
