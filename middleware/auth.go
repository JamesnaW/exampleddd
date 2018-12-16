package middleware

import (
	"exampleddd/pkg/auth/repository"

	"github.com/labstack/echo"
)

func AccessToken(rp repository.Repository) func(string, echo.Context) (bool, error) {
	return func(key string, c echo.Context) (bool, error) {
		if key == "" {
			c.Set("islogin", false)
			return true, nil
		}
		ok, id, err := rp.Get(key)
		if err != nil {
			return true, err
		}
		if !ok {
			c.Set("islogin", false)
			return true, nil
		}
		c.Set("isLogin", true)
		c.Set("userLoginId", id)
		return true, nil
	}
}
