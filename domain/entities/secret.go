package entities

import (
	"fmt"

	"github.com/matheusmosca/arch-experiment/domain/vos"
)

type Secret struct {
	id      vos.SecretID
	content string
}

func (s Secret) ID() vos.SecretID {
	return s.id
}

func (s Secret) Content() string {
	return s.content
}

func NewSecret(content string) (Secret, error) {
	if content == "" {
		return Secret{}, fmt.Errorf("empty content")
	}

	if len(content) > 500 {
		return Secret{}, fmt.Errorf("max length exceeded")
	}

	return Secret{
		id:      vos.NewSecretID(),
		content: content,
	}, nil
}

func ParseSecret(id string, content string) (Secret, error) {
	parsedID, err := vos.ParseSecretID(id)
	if err != nil {
		return Secret{}, err
	}

	return Secret{
		id:      parsedID,
		content: content,
	}, nil
}
