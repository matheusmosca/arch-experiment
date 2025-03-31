package usecases

import (
	"context"

	"github.com/matheusmosca/arch-experiment/domain/entities"
)

type saveSecret struct {
	secretsRepo secretsRepoSaveSecretUC
}

type secretsRepoSaveSecretUC interface {
	Save(context.Context, entities.Secret) error
}

type SaveSecretInput struct {
	Content string
}

type SaveSecretOutput struct {
	ID string
}

func (s saveSecret) SaveSecret(ctx context.Context, input SaveSecretInput) (SaveSecretOutput, error) {
	secret, err := entities.NewSecret(input.Content)
	if err != nil {
		return SaveSecretOutput{}, err
	}

	if err := s.secretsRepo.Save(ctx, secret); err != nil {
		return SaveSecretOutput{}, err
	}

	return SaveSecretOutput{
		ID: secret.ID().String(),
	}, nil
}

func NewSaveSecretUC(secrestsRepo secretsRepoSaveSecretUC) saveSecret {
	return saveSecret{
		secretsRepo: secrestsRepo,
	}
}
