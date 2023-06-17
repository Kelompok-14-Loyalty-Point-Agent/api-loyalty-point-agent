package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CORSConfig struct {
	AllowOrigins []string
	AllowMethods []string
}

func (c *CORSConfig) Init() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: c.AllowOrigins,
		AllowMethods: c.AllowMethods,
	})
}
