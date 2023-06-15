package commands

import (
	"context"

	"github.com/fd1az/mallbots/stores/internal/domain"
)

type EnableParticipation struct {
	ID string
}

type EnableParticipationHandler struct {
	stores domain.StoreRepository
}

func NewEnableParticipationHandler(stores domain.StoreRepository) EnableParticipationHandler {
	return EnableParticipationHandler{
		stores: stores,
	}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
	store, err := h.stores.Load(ctx, cmd.ID)

	if err != nil {
		return err
	}
	err = store.EnableParticipation()
	if err != nil {
		return err
	}
	return h.stores.Save(ctx, store)

}
