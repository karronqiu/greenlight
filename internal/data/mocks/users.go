package mocks

import "greenlight.karronqiu.github.com/internal/data"

type MockUserModel struct {
}

var mockUser = &data.User{
	Name:      "John Smith",
	Email:     "john@example.com",
	Activated: true,
}

func (m MockUserModel) Insert(user *data.User) error {
	return nil
}

func (m MockUserModel) GetByEmail(email string) (*data.User, error) {
	return mockUser, nil
}

func (m MockUserModel) Update(user *data.User) error {
	return nil
}

func (m MockUserModel) GetForToken(tokenScope, tokenPlaintext string) (*data.User, error) {
	return mockUser, nil
}
