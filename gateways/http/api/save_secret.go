package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/matheusmosca/arch-experiment/domain/usecases"
)

type saveSecretUC interface {
	SaveSecret(ctx context.Context, input usecases.SaveSecretInput) (usecases.SaveSecretOutput, error)
}

type saveSecretHandler struct {
	usecase saveSecretUC
}

type saveSecretResponse struct {
	ID string `json:"id"`
}

type saveSecretRequestBody struct {
	Content string `json:"content"`
}

func (h saveSecretHandler) SaveSecret(w http.ResponseWriter, r *http.Request) {
	var body saveSecretRequestBody

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "malformed paremeters"}`))
		return
	}

	secret, err := h.usecase.SaveSecret(r.Context(), usecases.SaveSecretInput{
		Content: body.Content,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "internal server error"}`))
		slog.Error(err.Error())
		return
	}

	response := saveSecretResponse{
		ID: secret.ID,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
}

func NewSaveSecret(uc saveSecretUC) saveSecretHandler {
	return saveSecretHandler{
		usecase: uc,
	}
}
