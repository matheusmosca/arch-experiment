package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/matheusmosca/arch-experiment/domain"
	"github.com/matheusmosca/arch-experiment/domain/entities"
)

func TestGetTicket(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx   context.Context
		input GetTicketDTO
	}

	dumbTime := time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		args    args
		usecase getTicketUC
		wantErr error
		want    entities.Ticket
	}{
		{
			name: "should get a ticket by id successfully",
			args: args{
				ctx: context.Background(),
				input: GetTicketDTO{
					ID: "731411a7-96ce-4629-ad05-c408bdb530b9",
				},
			},
			usecase: getTicketUC{
				ticketRepo: &ticketRepoGetTicketUCMock{
					GetTicketByIDFunc: func(ctx context.Context, id string) (entities.Ticket, error) {
						assert.Equal(t, "731411a7-96ce-4629-ad05-c408bdb530b9", id)

						return entities.Ticket{
							ID:        "731411a7-96ce-4629-ad05-c408bdb530b9",
							CreatedAt: dumbTime,
							UpdatedAt: dumbTime,
						}, nil
					},
				},
			},
			wantErr: nil,
			want: entities.Ticket{
				ID:        "731411a7-96ce-4629-ad05-c408bdb530b9",
				CreatedAt: dumbTime,
				UpdatedAt: dumbTime,
			},
		},
		{
			name: "should return an error when ticket is not found by repo",
			args: args{
				ctx: context.Background(),
				input: GetTicketDTO{
					ID: "731411a7-96ce-4629-ad05-c408bdb530b9",
				},
			},
			usecase: getTicketUC{
				ticketRepo: &ticketRepoGetTicketUCMock{
					GetTicketByIDFunc: func(ctx context.Context, id string) (entities.Ticket, error) {
						assert.Equal(t, "731411a7-96ce-4629-ad05-c408bdb530b9", id)

						return entities.Ticket{}, domain.ErrInvalidNotFound
					},
				},
			},
			wantErr: domain.ErrInvalidNotFound,
			want:    entities.Ticket{},
		},
		{
			name: "should return an error when input is invalid due to empty id",
			args: args{
				ctx: context.Background(),
				input: GetTicketDTO{
					ID: "",
				},
			},
			usecase: getTicketUC{},
			wantErr: domain.ErrInvalidParams,
			want:    entities.Ticket{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.usecase.GetTicket(tt.args.ctx, tt.args.input)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
