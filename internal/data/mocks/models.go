package mocks

import (
	"greenlight.karronqiu.github.com/internal/data"
)

// TODO: use https://github.com/stretchr/testify to simplify the mock code

func NewModels() data.Models {
	return data.Models{
		Movies:      MockMovieModel{},
		Users:       MockUserModel{},
		Tokens:      MockTokenModel{},
		Permissions: MockPermissionModel{},
	}
}
