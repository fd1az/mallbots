package handlers

import (
	"github.com/fd1az/mallbots/internal/ddd"
	"github.com/fd1az/mallbots/ordering/internal/application"
	"github.com/fd1az/mallbots/ordering/internal/domain"
)

func RegisterInvoiceHandlers(invoiceHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.OrderReadied{}, invoiceHandlers.OnOrderReadied)
}
