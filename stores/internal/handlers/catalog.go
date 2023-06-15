package handlers

import (
	"github.com/fd1az/mallbots/internal/ddd"
	"github.com/fd1az/mallbots/stores/internal/domain"
)

func RegisterCatalogHandlers(catalogHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(domain.ProductAddedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductRebrandedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductPriceIncreasedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductPriceDecreasedEvent, catalogHandlers)
	domainSubscriber.Subscribe(domain.ProductRemovedEvent, catalogHandlers)
}
