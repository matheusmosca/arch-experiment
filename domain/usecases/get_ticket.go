package usecases

import (
	"context"
	"fmt"

	"github.com/matheusmosca/arch-experiment/domain"
	"github.com/matheusmosca/arch-experiment/domain/entities"
)

//go:generate moq -out mocks.go . ticketRepoGetTicketUC
type ticketRepoGetTicketUC interface {
	GetTicketByID(ctx context.Context, id string) (entities.Ticket, error)
}

type getTicketUC struct {
	ticketRepo ticketRepoGetTicketUC
}

type GetTicketDTO struct {
	ID string
}

func (uc getTicketUC) GetTicket(ctx context.Context, input GetTicketDTO) (entities.Ticket, error) {
	if input.ID == "" {
		return entities.Ticket{}, fmt.Errorf("%w: empty id: %s", domain.ErrInvalidParams, input.ID)
	}

	return uc.ticketRepo.GetTicketByID(ctx, input.ID)
}

func NewGetTicketUC(ticketRepo ticketRepoGetTicketUC) getTicketUC {
	return getTicketUC{
		ticketRepo: ticketRepo,
	}
}
