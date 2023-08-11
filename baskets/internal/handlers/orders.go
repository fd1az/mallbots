package handlers

import (
	"github.com/fd1az/mallbots/baskets/internal/domain"
	"github.com/fd1az/mallbots/internal/ddd"
)

func RegisterOrderHandlers(
	orderHandlers ddd.EventHandler[ddd.AggregateEvent],
	domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent],
) {
	domainSubscriber.Subscribe(domain.BasketCheckedOutEvent, orderHandlers)
}
