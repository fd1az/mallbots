package ordering

import (
	"context"

	"github.com/fd1az/mallbots/internal/ddd"
	"github.com/fd1az/mallbots/internal/es"
	"github.com/fd1az/mallbots/internal/monolith"
	pg "github.com/fd1az/mallbots/internal/postgres"
	"github.com/fd1az/mallbots/internal/registry"
	"github.com/fd1az/mallbots/internal/registry/serdes"
	"github.com/fd1az/mallbots/ordering/internal/application"
	"github.com/fd1az/mallbots/ordering/internal/domain"
	"github.com/fd1az/mallbots/ordering/internal/grpc"
	"github.com/fd1az/mallbots/ordering/internal/handlers"
	"github.com/fd1az/mallbots/ordering/internal/logging"
	"github.com/fd1az/mallbots/ordering/internal/rest"
)

type Module struct{}

func (Module) Startup(ctx context.Context, mono monolith.Monolith) error {

	reg := registry.New()
	err := registrations(reg)
	if err != nil {
		return err
	}
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	aggregateStore := es.AggregateStoreWithMiddleware(
		pg.NewEventStore("stores.events", mono.DB(), reg),
		es.NewEventPublisher(domainDispatcher),
		pg.NewSnapshotStore("stores.snapshots", mono.DB(), reg),
	)
	orders := es.NewAggregateRepository[*domain.Order](
		domain.OrderAggregate,
		reg,
		aggregateStore,
	)
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}
	customers := grpc.NewCustomerRepository(conn)
	payments := grpc.NewPaymentRepository(conn)
	invoices := grpc.NewInvoiceRepository(conn)
	shopping := grpc.NewShoppingListRepository(conn)
	notifications := grpc.NewNotificationRepository(conn)

	// setup application
	var app application.App
	app = application.New(orders, customers, payments, shopping)
	app = logging.LogApplicationAccess(app, mono.Logger())
	// setup application handlers
	notificationHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewNotificationHandlers(notifications),
		"Notification", mono.Logger(),
	)
	invoiceHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
		application.NewInvoiceHandlers(invoices),
		"Invoice", mono.Logger(),
	)

	// setup Driver adapters
	if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)
	handlers.RegisterInvoiceHandlers(invoiceHandlers, domainDispatcher)

	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Store
	if err = serde.Register(domain.Order{}, func(v any) error {
		order := v.(*domain.Order)
		order.Aggregate = es.NewAggregate("", domain.OrderAggregate)
		return nil
	}); err != nil {
		return
	}
	// store events
	if err = serde.Register(domain.OrderCreated{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.OrderCanceledEvent, domain.OrderCanceled{}); err != nil {
		return
	}
	if err = serde.RegisterKey(domain.OrderCompletedEvent, domain.OrderCompleted{}); err != nil {
		return
	}
	if err = serde.Register(domain.OrderReadied{}); err != nil {
		return
	}
	// store snapshots
	if err = serde.RegisterKey(domain.OrderV1{}.SnapshotName(), domain.OrderV1{}); err != nil {
		return
	}

	return
}
