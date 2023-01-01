package mocks

import "greenlight.karronqiu.github.com/internal/data"

type MockPermissionModel struct {
}

var mockPermissions = data.Permissions{
	"movies:read",
	"movies:write",
}

func (m MockPermissionModel) GetAllForUser(userID int64) (data.Permissions, error) {
	return mockPermissions, nil
}

func (m MockPermissionModel) AddForUser(userID int64, codes ...string) error {
	return nil
}
