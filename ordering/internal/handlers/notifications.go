package handlers

import (
	"github.com/fd1az/mallbots/internal/ddd"
	"github.com/fd1az/mallbots/ordering/internal/domain"
)

func RegisterNotificationHandlers(
	notificationHandlers ddd.EventHandler[ddd.AggregateEvent],
	domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent],
) {
	domainSubscriber.Subscribe(domain.OrderCreatedEvent, notificationHandlers)
	domainSubscriber.Subscribe(domain.OrderReadiedEvent, notificationHandlers)
	domainSubscriber.Subscribe(domain.OrderCanceledEvent, notificationHandlers)
}
