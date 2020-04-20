package factories

import "github.com/google/uuid"

func GenerateNewSessionToken() (string, error) {
	_uuid, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	return _uuid.String(), nil
}
