package handlers

import (
	"github.com/fd1az/mallbots/baskets/internal/application"
	"github.com/fd1az/mallbots/baskets/internal/domain"
	"github.com/fd1az/mallbots/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.BasketCheckedOut{}, orderHandlers.OnBasketCheckedOut)
}
