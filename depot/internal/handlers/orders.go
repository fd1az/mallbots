package handlers

import (
	"github.com/fd1az/mallbots/depot/internal/domain"
	"github.com/fd1az/mallbots/internal/ddd"
)

func RegisterOrderHandlers(
	orderHandlers ddd.EventHandler[ddd.AggregateEvent],
	domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent],
) {
	domainSubscriber.Subscribe(domain.ShoppingListCompletedEvent, orderHandlers)
}
