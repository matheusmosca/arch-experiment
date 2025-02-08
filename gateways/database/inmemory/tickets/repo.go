package tickets

import (
	"context"
	"time"

	"github.com/matheusmosca/arch-experiment/domain/entities"
)

type repo struct {
}

func (r repo) GetTicketByID(ctx context.Context, id string) (entities.Ticket, error) {
	now := time.Now()

	return entities.Ticket{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func NewRepository() repo {
	return repo{}
}
