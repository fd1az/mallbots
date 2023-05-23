package stores

import (
	"context"

	"github.com/fd1az/mallbots/internal/ddd"
	"github.com/fd1az/mallbots/internal/monolith"
	"github.com/fd1az/mallbots/stores/internal/application"
	"github.com/fd1az/mallbots/stores/internal/grpc"
	"github.com/fd1az/mallbots/stores/internal/logging"
	"github.com/fd1az/mallbots/stores/internal/postgres"
	"github.com/fd1az/mallbots/stores/internal/rest"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	// setup Driven adapters
	domainDispatcher := ddd.NewEventDispatcher()
	stores := postgres.NewStoreRepository("stores.stores", mono.DB())
	participatingStores := postgres.NewParticipatingStoreRepository("stores.stores", mono.DB())
	products := postgres.NewProductRepository("stores.products", mono.DB())

	// setup application
	var app application.App
	app = application.New(stores, participatingStores, products, domainDispatcher)
	app = logging.LogApplicationAccess(app, mono.Logger())

	// setup Driver adapters
	if err := grpc.RegisterServer(ctx, app, mono.RPC()); err != nil {
		return err
	}
	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
