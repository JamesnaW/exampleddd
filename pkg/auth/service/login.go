package service

import (
	"errors"
	"exampleddd/pkg/entity"
)

// Login business logic
func (s service) Login(user, pass string) (entity.User, error) {
	profile := entity.User{}
	ok, id, err := s.rp.Check(user, pass)
	if err != nil {
		return profile, err
	}
	if !ok {
		return profile, errors.New("Username or password incorrect")
	}
	profile, err = s.profile.Get(id)
	return profile, nil
}
