package mocks

import (
	"time"

	"greenlight.karronqiu.github.com/internal/data"
)

type MockTokenModel struct {
}

func (m MockTokenModel) New(userID int64, ttl time.Duration, scope string) (*data.Token, error) {
	return nil, nil
}

func (m MockTokenModel) Insert(token *data.Token) error {
	return nil
}

func (m MockTokenModel) DeleteAllForUser(scope string, userID int64) error {
	return nil
}
