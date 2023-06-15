package handlers

import (
	"github.com/fd1az/mallbots/depot/internal/application"
	"github.com/fd1az/mallbots/depot/internal/domain"
	"github.com/fd1az/mallbots/internal/ddd"
)

func RegisterOrderHandlers(orderHandlers application.DomainEventHandlers, domainSubscriber ddd.EventSubscriber) {
	domainSubscriber.Subscribe(domain.ShoppingListCompleted{}, orderHandlers.OnShoppingListCompleted)
}
