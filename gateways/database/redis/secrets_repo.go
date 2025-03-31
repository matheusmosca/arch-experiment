package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/matheusmosca/arch-experiment/domain"
	"github.com/matheusmosca/arch-experiment/domain/entities"
	"github.com/matheusmosca/arch-experiment/domain/vos"
	"github.com/redis/go-redis/v9"
)

type SecretsRepo struct {
	client *redis.Client
}

func NewSecretsRepo(client *redis.Client) SecretsRepo {
	return SecretsRepo{
		client: client,
	}
}

func (s SecretsRepo) Save(ctx context.Context, secret entities.Secret) error {
	return s.client.Set(ctx, secret.ID().String(), secret.Content(), time.Hour).Err()
}

func (s SecretsRepo) GetAndDelete(ctx context.Context, secretID vos.SecretID) (entities.Secret, error) {
	content, err := s.client.GetDel(ctx, secretID.String()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return entities.Secret{}, fmt.Errorf("secret not found: %w", domain.ErrInvalidNotFound)
		}
		return entities.Secret{}, err
	}

	secret, err := entities.ParseSecret(secretID.String(), content)
	if err != nil {
		return entities.Secret{}, err
	}

	return secret, nil
}
