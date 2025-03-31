package api

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matheusmosca/arch-experiment/domain"
	"github.com/matheusmosca/arch-experiment/domain/usecases"
)

type getSecretUC interface {
	GetSecret(ctx context.Context, input usecases.GetSecretInput) (usecases.GetSecretOutput, error)
}

type getSecretHandler struct {
	usecase getSecretUC
}

type getSecretResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func (h getSecretHandler) GetSecret(w http.ResponseWriter, r *http.Request) {
	secretID := chi.URLParam(r, "id")

	secret, err := h.usecase.GetSecret(r.Context(), usecases.GetSecretInput{
		ID: secretID,
	})
	if err != nil {
		if errors.Is(err, domain.ErrInvalidNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error": "secret not found"}`))
			slog.Error(err.Error())
			return

		}

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "internal server error"}`))
		slog.Error(err.Error())
		return
	}

	response := getSecretResponse{
		ID:      secret.ID,
		Content: secret.Content,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error(err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

func NewGetSecret(uc getSecretUC) getSecretHandler {
	return getSecretHandler{
		usecase: uc,
	}
}
