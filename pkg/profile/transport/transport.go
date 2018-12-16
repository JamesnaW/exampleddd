package auth

import "github.com/labstack/echo"

// Route init
func Route(e *echo.Group, h *handler, mm ...echo.MiddlewareFunc) {
	for _, m := range mm {
		e.Use(m)
	}

}
