package service

import (
	"exampleddd/pkg/profile/repository"
)

type service struct {
	rp repository.Repository
}

// NewService init
func NewService(rp repository.Repository) *service {
	return &service{
		rp,
	}
}

// Service interface
type Service interface {
}
