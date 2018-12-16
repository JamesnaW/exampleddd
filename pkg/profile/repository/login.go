package repository

import (
	"exampleddd/pkg/entity"
)

// Get user profile
func (r repo) Get(id int) (entity.User, error) {
	if id != 1 {
		return entity.User{}, nil
	}
	profile := entity.User{}
	profile.ID = 1
	profile.Username = "admin"
	return profile, nil
}
