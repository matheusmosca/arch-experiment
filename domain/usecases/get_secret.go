package usecases

import (
	"context"

	"github.com/matheusmosca/arch-experiment/domain/entities"
	"github.com/matheusmosca/arch-experiment/domain/vos"
)

//go:generate moq -out mocks.go . getSecretRepo
type getSecretRepo interface {
	GetAndDelete(ctx context.Context, secretID vos.SecretID) (entities.Secret, error)
}

type getSecretUC struct {
	secretsRepo getSecretRepo
}

type GetSecretInput struct {
	ID string
}

func (uc getSecretUC) GetSecret(ctx context.Context, input GetSecretInput) (GetSecretOutput, error) {
	secretID, err := vos.ParseSecretID(input.ID)
	if err != nil {
		return GetSecretOutput{}, err
	}

	secret, err := uc.secretsRepo.GetAndDelete(ctx, secretID)
	if err != nil {
		return GetSecretOutput{}, err
	}

	return GetSecretOutput{
		ID:      secret.ID().String(),
		Content: secret.Content(),
	}, nil
}

func NewGetSecretUC(secretsRepo getSecretRepo) getSecretUC {
	return getSecretUC{
		secretsRepo: secretsRepo,
	}
}

type GetSecretOutput struct {
	ID      string
	Content string
}
