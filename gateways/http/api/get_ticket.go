package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/matheusmosca/arch-experiment/domain/entities"
	"github.com/matheusmosca/arch-experiment/domain/usecases"
)

type getTicketUC interface {
	GetTicket(ctx context.Context, input usecases.GetTicketDTO) (entities.Ticket, error)
}

type getTicketHandler struct {
	usecase getTicketUC
}

type getTicketResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h getTicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	ticketID := chi.URLParam(r, "id")

	ticket, err := h.usecase.GetTicket(r.Context(), usecases.GetTicketDTO{
		ID: ticketID,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "internal server error"}`))
		slog.Error(err.Error(), slog.String("id", ticketID))
		return
	}

	response := getTicketResponse{
		ID:        ticket.ID,
		CreatedAt: ticket.CreatedAt,
		UpdatedAt: ticket.UpdatedAt,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error(err.Error())
	}
}

func NewGetTicket(uc getTicketUC) getTicketHandler {
	return getTicketHandler{
		usecase: uc,
	}
}
