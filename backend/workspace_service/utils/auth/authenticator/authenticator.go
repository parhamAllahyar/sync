package authenticator

import "github.com/google/uuid"

func IsAdminToken(token string) (uuid.UUID, error) {
	userId := uuid.New()
	return userId, nil
}

func IsUserToken(token string) (uuid.UUID, error) {
	userId := uuid.New()
	return userId, nil
}
