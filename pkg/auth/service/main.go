package service

import (
	"exampleddd/pkg/auth/repository"
	"exampleddd/pkg/entity"
	profile "exampleddd/pkg/profile/repository"
)

type service struct {
	rp      repository.Repository
	profile profile.Repository
}

// NewService init
func NewService(rp repository.Repository, profile profile.Repository) *service {
	return &service{
		rp,
		profile,
	}
}

// Service interface
type Service interface {
	Login(string, string) (entity.User, error)
}
