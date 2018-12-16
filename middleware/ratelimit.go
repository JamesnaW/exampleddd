package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/time/rate"
)

func RateLimit(limiter *rate.Limiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if !limiter.Allow() {
				c.JSON(http.StatusTooManyRequests, nil)
				return next(c)
			}
			return next(c)
		}
	}
}
