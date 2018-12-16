package auth

import (
	"exampleddd/pkg/auth/service"
	"exampleddd/pkg/entity"

	"github.com/labstack/echo"
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

// Login business logic
func (h handler) Login(c echo.Context) error {
	user := c.FormValue("username")
	pass := c.FormValue("password")
	results := entity.Results{}
	if user == "" {
		results.Error = "Username incorrect"
		return c.JSON(200, results)
	}
	if pass == "" {
		results.Error = "Password incorrect"
		return c.JSON(200, results)
	}

	profile, err := h.s.Login(user, pass)
	if profile != (entity.User{}) {
		results.Message = profile
	}
	if err != nil {
		results.Error = err.Error()
	}
	return c.JSON(200, results)
}
