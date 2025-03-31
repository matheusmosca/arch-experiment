package vos

import "github.com/google/uuid"

type SecretID struct {
	value uuid.UUID
}

func NewSecretID() SecretID {
	return SecretID{
		value: uuid.New(),
	}
}

func ParseSecretID(id string) (SecretID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return SecretID{}, err
	}

	return SecretID{
		value: parsedID,
	}, nil
}

func (s SecretID) String() string {
	return s.value.String()
}
