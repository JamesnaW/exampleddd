package auth

import (
	"exampleddd/pkg/profile/service"
)

type handler struct {
	s service.Service
}

// NewHandler init
func NewHandler(s service.Service) *handler {
	return &handler{
		s: s,
	}
}
